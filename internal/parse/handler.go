/*
Copyright 2023 Sangfor Technologies Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package parse

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"arsenal/pkg/data"
	"arsenal/pkg/logs"
	executor "arsenal/pkg/run"
)

const (
	Cli  = "cli"
	HTTP = "http"
)

var OpsHandlerOperationsTypes = map[string]OpsHandlerOperations{}

// OpsHandlerOperations 故障相关操作的集合。
type OpsHandlerOperations interface {
	// SetRecordedFlag 设定数据库是否记录表项的标志位。
	SetRecordedFlag(inputInfo *data.InputInfo, opsInfo *data.OpsInfo) error

	// OperationCheck 判断操作的合法性。
	OperationsCheck(inputInfo *data.InputInfo, opsInfo *data.OpsInfo) error

	// SetTimeoutValue 通过输入参数设定延迟执行的时间。
	SetTimeoutValue(inputInfo *data.InputInfo)

	// SetShellCmd 设定运行的shell命令。
	SetShellCmd(inputInfo *data.InputInfo, opsInfo *data.OpsInfo) error

	// KillFaultCleanBackgroundProcess timeout时间没有到提前清理后台执行进程。
	KillFaultCleanBackgroundProcess(inputInfo *data.InputInfo, opsInfo *data.OpsInfo) error

	// RunFaultCleanBackgroundProcess 运行timeout延迟清理进程。
	RunFaultCleanBackgroundProcess(inputInfo *data.InputInfo, opsInfo *data.OpsInfo, cmd string) error
}

// Add 将对应控制类型的接口添加进入map。
func Add(interactiveMode string, newType OpsHandlerOperations) {
	OpsHandlerOperationsTypes[interactiveMode] = newType
}

// OpsHandler 故障相关操作的处理函数入口。
func OpsHandler(inputInfo *data.InputInfo, opsInfo *data.OpsInfo) error {
	handler, ok := OpsHandlerOperationsTypes[inputInfo.InteractiveMode]
	if !ok {
		return fmt.Errorf("invalid interactive mode(%s)", inputInfo.InteractiveMode)
	}

	// 获取相关的标志信息。
	if err := setRefFlags(inputInfo, opsInfo, handler); err != nil {
		return fmt.Errorf("%v", err)
	}

	// 操作合法性判断。
	if err := handler.OperationsCheck(inputInfo, opsInfo); err != nil {
		return fmt.Errorf("%v", err)
	}

	// 获取timeout时间，用于判断是否做故障延时清理操作。
	handler.SetTimeoutValue(inputInfo)

	// 拼接原子故障注入命令并执行。
	if err := handler.SetShellCmd(inputInfo, opsInfo); err != nil {
		return err
	}

	// 处理timeout时间还没有达到，故障需要做提前清理的场景下。
	if err := KillTimeoutBackgroundProcess(inputInfo, opsInfo); err != nil {
		return fmt.Errorf("kill fault clean background process failed: %v", err)
	}

	_, err := executor.ShellCmd(inputInfo.OpsType, inputInfo.BlockExecution, inputInfo.ShellCmd)
	if err != nil {
		return fmt.Errorf("ops handler: %v", err)
	}

	// 如果输入参数timeout > 0且操作类型为注入，则需要后台执行故障清理命令。
	if inputInfo.Timeout > 0 && inputInfo.OpsType == data.Inject {
		cmd, err := getTimeoutShellCmdByOpsInfo(opsInfo)
		if err != nil {
			return fmt.Errorf("get timeout shell command failed: %v", err)
		}
		logs.Debug("timeout backup shell cmd: %s", cmd)

		if err = handler.RunFaultCleanBackgroundProcess(inputInfo, opsInfo, cmd); err != nil {
			return fmt.Errorf("run background fault cleaning process failed: %v", err)
		}
	}

	logInfo(inputInfo, opsInfo)
	return UpdateDB(inputInfo, opsInfo)
}

// KillTimeoutBackgroundProcess 杀掉后台运行的timeout进程。
func KillTimeoutBackgroundProcess(inputInfo *data.InputInfo, opsInfo *data.OpsInfo) error {
	// 需要清零后台故障清理进程的条件：
	// 1. 私有数据区包含数据，当前私有数据区用于存放pid和timeout时间；
	// 2. 操作类型为清理操作，只有清理操作才可能存在后台清理协程；
	if opsInfo.Private == "" || inputInfo.OpsType != data.Remove {
		return nil
	}

	// http类型timeout清理操作也是通过bash cli下发的，inputInfo.InteractiveMode为cli，
	// 所以需要采用从db中读取表项的InteractiveMode来选择对应的提前清理的操作方法。
	handler := OpsHandlerOperationsTypes[opsInfo.InteractiveMode]
	if handler == nil {
		return fmt.Errorf("invalid ops ctl path(%s)", opsInfo.InteractiveMode)
	}

	if err := handler.KillFaultCleanBackgroundProcess(inputInfo, opsInfo); err != nil {
		return fmt.Errorf("kill fault clean background process failed(%v)", err)
	}
	return nil
}

func setRefFlags(inputInfo *data.InputInfo, opsInfo *data.OpsInfo, handler OpsHandlerOperations) error {
	if err := handler.SetRecordedFlag(inputInfo, opsInfo); err != nil {
		return fmt.Errorf("get record flag failed(%v)", err)
	}

	// 从配置文件中获取阻塞执行标志位。
	inputInfo.SetBlockExecutionFlag()

	// 从配置文件获取手动清理标志位，如果不包含remove命令，说明不需要手动清理故障。
	inputInfo.SetShouldRemoveFlag()
	if inputInfo.ShouldRemove {
		opsInfo.ProactiveCleanup = true
	} else {
		opsInfo.ProactiveCleanup = false
		opsInfo.Status = data.SucceedStatus
	}
	return nil
}

func getTimeoutShellCmdByOpsInfo(opsInfo *data.OpsInfo) (string, error) {
	arsenal, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("get arsenal path failed: %v", err)
	}

	return fmt.Sprintf("%s remove %s %s %s %s", arsenal, opsInfo.Env,
		opsInfo.Domain, opsInfo.FaultType, opsInfo.Flags), nil
}

// logInfo 将故障注入信息写入日志文件实现持久化。
func logInfo(inputInfo *data.InputInfo, opsInfo *data.OpsInfo) {
	logs.WithFields(logrus.Fields{
		"runShellCommand": inputInfo.ShellCmd,
		"opsType":         inputInfo.OpsType,
		"CustomCommand":   inputInfo.Exe.CustomCommand,
		"isRecord":        inputInfo.IsRecorded,
		"blockExecution":  inputInfo.BlockExecution,
		"shouldRemove":    inputInfo.ShouldRemove,
		"timeout":         inputInfo.Timeout,
	}).Info(opsInfo.InteractiveMode, " input info")

	logs.WithFields(logrus.Fields{
		"updateTime":       opsInfo.UpdateTime,
		"injectTime":       opsInfo.InjectTime,
		"status":           opsInfo.Status,
		"proactiveCleanup": opsInfo.ProactiveCleanup,
		"private":          opsInfo.Private,
		"flags":            opsInfo.Flags,
		"faultType":        opsInfo.FaultType,
		"domain":           opsInfo.Domain,
		"interactiveMode":  opsInfo.InteractiveMode,
		"Env":              opsInfo.Env,
		"uuid":             opsInfo.UUID,
	}).Info(opsInfo.InteractiveMode, " operation information")
}

func UpdateDB(inputInfo *data.InputInfo, opsInfo *data.OpsInfo) error {
	var err error
	const retryCount = 3
	for i := 0; i < retryCount; i++ {
		if err = inputInfo.UpdateDB(opsInfo); err == nil {
			return nil
		}
		time.Sleep(time.Second)
	}

	return fmt.Errorf("unable to update database after retrying %d times, error(%v)",
		retryCount, err)
}
