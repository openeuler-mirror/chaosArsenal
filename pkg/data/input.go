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

package data

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"arsenal/pkg/util"
)

const (
	Inject        = "inject"   // 操作类型为故障注入。
	InjectStatus  = "injected" // 故障处于生效状态。
	Remove        = "remove"   // 操作类型为故障清理。
	RemovedStatus = "removed"  // 故障已经被清理状态。
	SucceedStatus = "succeed"  // 用来标识不需要手动清理的故障状态。
)

// InputInfo 用于描述输入的故障相关信息。
type InputInfo struct {
	// Flags 输入参数的map信息。
	Flags map[string]string

	// FlagsString 输入参数组成的字符串。
	FlagsString string

	// FaultType 故障模式。
	FaultType string

	// Env 运行环境。
	Env string

	// Domain 故障模式所属域。
	Domain string

	// ID 故障注入时产生的ID。
	UUID string

	// OpsType 故障操作的类型。
	OpsType string

	// IsRecorded 在数据库中是否存在的标志位。
	IsRecorded bool

	// Exe 具体故障模式对应可执行文件的路径信息。
	Exe exeInfo

	// BlockExecution shell命令阻塞执行标志位。
	BlockExecution bool

	// ShouldRemove 故障是否需要主动清理标志位。
	ShouldRemove bool

	// Timeout 故障延迟清理时间。
	Timeout uint64

	// TimeoutStr 故障延迟清理时间原始字符串。
	TimeoutStr string

	// ShellCmd 调用故障注入原子能力shell命令。
	ShellCmd string

	// Object 故障注入对象信息。
	Object string

	// Cmd 故障模式对应的cobra cmd。
	Cmd *cobra.Command

	// InteractiveMode 请求下发路径。
	InteractiveMode string
}

type exeInfo struct {
	// ExePath 原子故障注入能力可执行文件路径。
	ExePath string

	// CustomCommand 配置文件是否有完整命令标志位。
	CustomCommand bool
}

func (i *InputInfo) getEnvKey() string {
	return fmt.Sprintf("env.%s", i.Env)
}

func (i *InputInfo) getDomainKey() string {
	return fmt.Sprintf("%s.domain.%s", i.getEnvKey(), i.Domain)
}

func (i *InputInfo) getExecutorKey() string {
	return fmt.Sprintf("%s.domain.%s.executor", i.getEnvKey(), i.Domain)
}

func (i *InputInfo) getFaultTypesKey() string {
	return fmt.Sprintf("%s.domain.%s.faultTypes", i.getEnvKey(), i.Domain)
}

func (i *InputInfo) getFaultTypeKey() string {
	return fmt.Sprintf("%s.%s", i.getFaultTypesKey(), i.FaultType)
}

func (i *InputInfo) getObjectKey() string {
	return fmt.Sprintf("%s.%s.object", i.getFaultTypesKey(), i.FaultType)
}

func (i *InputInfo) getCommandsKey() string {
	return fmt.Sprintf("%s.%s.commands", i.getFaultTypesKey(), i.FaultType)
}

func (i *InputInfo) getFlagsKey() string {
	return fmt.Sprintf("%s.%s.flags", i.getFaultTypesKey(), i.FaultType)
}

func (i *InputInfo) getCommandKey() string {
	return fmt.Sprintf("%s.%s.command", i.getCommandsKey(), i.OpsType)
}

func (i *InputInfo) getBlockExecutionKey() string {
	return fmt.Sprintf("%s.%s.blockExecution", i.getCommandsKey(), i.OpsType)
}

func (i *InputInfo) getRemoveCommandKey() string {
	return fmt.Sprintf("%s.remove", i.getCommandsKey())
}

func (i *InputInfo) getFlagRequiredKey(flag string) string {
	return fmt.Sprintf("%s.%s.required", i.getFlagsKey(), flag)
}

func (i *InputInfo) getShellCmd() string {
	return fmt.Sprintf("%s %s %s %s %s", i.Exe.ExePath, i.OpsType, i.Domain,
		i.FaultType, i.FlagsString)
}

// SetFaultCheckObject 获取配置文件中故障的对象信息。
func (i *InputInfo) SetFaultCheckObject() error {
	objectStr := viper.GetString(i.getObjectKey())
	if objectStr == "" {
		return fmt.Errorf("please make sure %s %s %s has object attribute",
			i.Env, i.Domain, i.FaultType)
	}
	switch objectStr {
	case "NA":
		i.Object = "NA"
	case i.FaultType:
		i.Object = i.FaultType
	default:
		objValue := i.Flags[objectStr]
		if objValue == "" {
			return fmt.Errorf("please make sure %s %s %s object attribute"+
				" in flags", i.Env, i.Domain, i.FaultType)
		}
		i.Object = objValue
	}
	return nil
}

// SetRecordedFlag 在数据库中查找opsInfo是否存在。
func (i *InputInfo) SetRecordedFlag(info *OpsInfo) error {
	isRecord, err := info.OpsInfoInDatabase(i.OpsType)
	if err != nil {
		return fmt.Errorf("find ops info in database failed: %v", err)
	}
	i.IsRecorded = isRecord
	return nil
}

// SetBlockExecutionFlag 获取配置文件中对应命令阻塞执行的标志位。
func (i *InputInfo) SetBlockExecutionFlag() {
	if !viper.GetBool(i.getBlockExecutionKey()) {
		i.BlockExecution = false
	} else {
		i.BlockExecution = true
	}
}

// ParseTimeout 将传入的timeout字符串类型转换成秒。
func (i *InputInfo) ParseTimeout(timeoutStr string) (uint64, error) {
	timeoutRegex := regexp.MustCompile(`^(?:(\d+d):)?(?:(\d+h):)?(?:(\d+m):)?(?:(\d+s))?$`)
	if !timeoutRegex.MatchString(timeoutStr) {
		return 0, fmt.Errorf("invalid timeout format %s, valid format: 1h,1m,1s,1h:1s,1h:1m:1s", timeoutStr)
	}

	var duration time.Duration
	parts := strings.Split(timeoutStr, ":")
	for _, part := range parts {
		// 获取每个时间等级的设定值。
		value, err := strconv.Atoi(part[:len(part)-1])
		if err != nil {
			return 0, err
		}

		// 获取倒数第一个字符为单位。
		unit := part[len(part)-1]
		switch unit {
		case 'd':
			const hoursPerDay = 24
			duration += time.Duration(value) * time.Hour * hoursPerDay
		case 'h':
			duration += time.Duration(value) * time.Hour
		case 'm':
			duration += time.Duration(value) * time.Minute
		case 's':
			duration += time.Duration(value) * time.Second
		default:
			return 0, errors.New("input invalid timeout format, example: 1h,1m,1s,1h:1s,1h:1m:1s")
		}
	}
	return uint64(duration / time.Second), nil
}

// SetShouldRemoveFlag 获取配置文件中是否存在故障清理命令来判断是否需要做故障清理操作。
func (i *InputInfo) SetShouldRemoveFlag() {
	removeFlag := viper.GetStringMap(i.getRemoveCommandKey())
	if len(removeFlag) != 0 {
		i.ShouldRemove = true
	} else {
		i.ShouldRemove = false
	}
}

// getShellCmdFromConfig 在自定义命令的使用场景，从配置文件中获取对应命令。
func (i *InputInfo) getShellCmdFromConfig() string {
	return viper.GetString(i.getCommandKey())
}

// getCommandArray 修改配置文件中shell命令中对应参数为cli传入参数，
// 如将$path替换为path这个参数所对应的值，将配置文件中命令$path修改为/bar/foo。
func getCommandArray(cmd *cobra.Command, shellCmd string) ([]string, error) {
	cmdStrArray := strings.Fields(shellCmd)
	const minValidElementsLength = 2
	for index, value := range cmdStrArray {
		if len(value) < minValidElementsLength {
			continue
		}

		// 如果字符串以$开头，则认为是配置文件中命令中的变量，需要替换成cli传入的参数。
		if value[0] == '$' {
			var err error
			// 获取$符之后的字符串，获取flag对应的值，并将配置文件中命令$option替换
			// flagName := value[1:]。
			cmdStrArray[index], err = cmd.Flags().GetString(value[1:])
			if err != nil {
				return nil, fmt.Errorf("get flag string failed: %v", err)
			}
		}
	}
	return cmdStrArray, nil
}

// SetFullShellCommandByInputInfo 根据已经初始化的inputInfo获取完整的故障注入命令,
// 如果配置文件中存在对应故障注入命令，说明是定制化命令，仅将命令中的变量替换成输入参数的值，
// 反之，从inputInfo中的字段来拼接完整的shell命令，如：env字段为os，则调用子工具二进制文件名
// 为arsenal-os，其他类推。
func (i *InputInfo) SetFullShellCommandByInputInfo() error {
	var executorName string
	var cmdStrArray []string
	configShellCmd := i.getShellCmdFromConfig()
	if configShellCmd != "" {
		var err error
		cmdStrArray, err = getCommandArray(i.Cmd, configShellCmd)
		if err != nil {
			return fmt.Errorf("replace the variable in the configuration file with a key value type failed: %v", err)
		}
		executorName = cmdStrArray[0]
		i.Exe.CustomCommand = true
	} else {
		executorName = viper.GetString(i.getExecutorKey())
	}

	if err := i.SetAtomicFaultInjectionExePath(executorName); err != nil {
		return fmt.Errorf("get executor full path failed: %v", err)
	}

	// 如果配置文件中有对应操作的command，则优先使用配置文件中的command，
	// 否则根据cli输入层级拼装shell命令。
	if i.Exe.CustomCommand {
		cmdStrArray[0] = i.Exe.ExePath
		i.ShellCmd = util.ArrayToString(cmdStrArray)
	} else {
		i.ShellCmd = i.getShellCmd()
	}
	return nil
}

// SetAtomicFaultInjectionExePath 获取具体模块可执行文件路径，这个路径相对于arsenal/bin/。
func (i *InputInfo) SetAtomicFaultInjectionExePath(executorName string) error {
	exePath, err := util.GetAtomicFaultInjectionExePath(executorName)
	if err != nil {
		return fmt.Errorf("get the atomic fault injection executable failed by input info: %v", err)
	}
	i.Exe.ExePath = exePath
	return nil
}

// CheckOps 根据输入参数判断相关操作是不是非法操作。
func (i *InputInfo) CheckOps(info *OpsInfo) error {
	// 配置文件中没有清理命令，下清理命令认为是非法操作。
	if !i.ShouldRemove && i.OpsType == Remove {
		return fmt.Errorf("the fault: %s does not require cleaning", i.FaultType)
	}
	// 当某个故障不存在清理命令时，没有故障注入对象，意味着可以连续做注入操作。
	if i.Object == "NA" && i.OpsType == Inject {
		return nil
	}

	switch i.OpsType {
	case Inject:
		// 在故障注入场景下，且数据库中可以查找到故障状态为Injected的表项，
		// 说明是重复注入，为非法操作。
		if i.IsRecorded {
			return fmt.Errorf("[%s %s-%s %s] has been injected by [%s]",
				info.UUID, info.Domain, info.FaultType, info.Flags, info.InteractiveMode)
		}
	case Remove:
		// 在故障清理场景，且数据库中没有查找到故障状态为Injected的表项，
		// 说明该故障没有注入过或者故障已经被清理，为非法操作。
		if !i.IsRecorded {
			return fmt.Errorf("the [%s-%s %s] has not been injected or has been removed",
				i.Domain, i.FaultType, i.FlagsString)
		}
		// 在通过http故障清理场景，当前只能通过传入id来清理故障，
		// 已经通过uuid在数据库中找到对应表项，如果状态为Removed，说明是重复清理，为非法操作。
		if info.Status == RemovedStatus {
			return fmt.Errorf("[%s %s-%s %s] has been removed", info.UUID, i.Domain, i.FaultType, info.Flags)
		}
	default:
		return errors.New("not support operation type")
	}
	return nil
}

// UpdateDB 故障注入、清除之后更新数据库信息，注入操作添加表项，清除操作
// 修改表项的状态信息由Injected修改为Removed。
func (i *InputInfo) UpdateDB(info *OpsInfo) error {
	db, err := OpenOpsInfoDatabase()
	if err != nil {
		return fmt.Errorf("open data base failed: %v", err)
	}
	defer db.Close()

	switch i.OpsType {
	case Inject:
		if err = info.InsertOpsInfoIntoDB(db); err != nil {
			return fmt.Errorf("insert ops info into data base failed: %v", err)
		}
	case Remove:
		if err = info.ModifyOpsInfoStatus(db, RemovedStatus); err != nil {
			return fmt.Errorf("modify fault's status as %s failed: %v", RemovedStatus, err)
		}
	default:
		return fmt.Errorf("unsupported operation type: %s when update db", i.OpsType)
	}
	return nil
}
