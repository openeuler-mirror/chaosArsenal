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
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"arsenal/pkg/data"
	"arsenal/pkg/logs"
	"arsenal/pkg/util"
)

func init() {
	var newHandlerOps = cliHandlerOperations{
		interactiveMode: Cli,
	}
	Add(newHandlerOps.interactiveMode, &newHandlerOps)
}

type cliHandlerOperations struct {
	interactiveMode string
}

// OperationsCheck 对cli输入命令进行合法性检查。
func (c *cliHandlerOperations) OperationsCheck(inputInfo *data.InputInfo, opsInfo *data.OpsInfo) error {
	return inputInfo.CheckOps(opsInfo)
}

// SetRecordedFlag 根据cli输入参数初始化的opsInfo在数据库中查找相应表项，如果表项存在，
// 将inputInfo成员IsRecorded置true。
func (c *cliHandlerOperations) SetRecordedFlag(inputInfo *data.InputInfo, opsInfo *data.OpsInfo) error {
	// cli方式注入和清理均需要查找数据库中对应表项。
	if err := inputInfo.SetRecordedFlag(opsInfo); err != nil {
		return fmt.Errorf("get recorded flag failed: %v", err)
	}

	// http交互模式注入的故障，需要用http的方式清理故障。
	if opsInfo.InteractiveMode == HTTP && inputInfo.OpsType == data.Remove {
		return fmt.Errorf("uuid: %s injected interactive mode is http, "+
			"use the http interaction to remove the fault", opsInfo.UUID)
	}
	return nil
}

// SetTimeoutValue 从输入的cli参数中获取timeout需要延迟的时间，单位为秒。
func (c *cliHandlerOperations) SetTimeoutValue(inputInfo *data.InputInfo) {
	flag := inputInfo.Cmd.PersistentFlags().Lookup("timeout")
	if flag == nil {
		return
	}
	flagValue := flag.Value.String()
	if flagValue == "" {
		return
	}

	inputInfo.TimeoutStr = flagValue
	timeout, err := inputInfo.ParseTimeout(flagValue)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	inputInfo.Timeout = timeout
}

// SetShellCmd 根据cli输入的参数拼装底层原子故障注入的shell命令。
func (c *cliHandlerOperations) SetShellCmd(inputInfo *data.InputInfo, _ *data.OpsInfo) error {
	if err := inputInfo.SetFullShellCommandByInputInfo(); err != nil {
		return fmt.Errorf("get full shell command failed: %v", err)
	}
	return nil
}

// KillFaultCleanBackgroundProcess 将cli后台运行延时清理的进程kill掉，私有数据的存放格式为$pid,1h:1h:10s。
func (c *cliHandlerOperations) KillFaultCleanBackgroundProcess(_ *data.InputInfo, opsInfo *data.OpsInfo) error {
	pid, err := strconv.Atoi(strings.Split(opsInfo.Private, ",")[0])
	if err != nil {
		return fmt.Errorf("get background fault cleaning pid failed: %v", err)
	}

	// 当前进程pid号与记录在数据库中后台执行清理进程pid号一致，
	// 说明是后台清理进程执行到该逻辑，直接返回nil，不做任何操作。
	if pid == os.Getpid() {
		return nil
	}
	logs.Debug("kill timeout fault clean process:, current process:", pid, os.Getpid())

	if !util.FileIsExist(fmt.Sprintf("/proc/%d", pid)) {
		return fmt.Errorf("background fault cleaning pid not exist: %v", err)
	}
	// TODO: 需要实现进程的优雅退出，当正在做清理操作时被强制退出可能引发问题。
	if err = syscall.Kill(pid, syscall.SIGKILL); err != nil {
		return fmt.Errorf("kill backup fault cleaning process failed: %v", err)
	}
	return nil
}

// RunFaultCleanBackgroundProcess 后台运行cli延迟清理进程。
func (c *cliHandlerOperations) RunFaultCleanBackgroundProcess(
	inputInfo *data.InputInfo,
	opsInfo *data.OpsInfo,
	cmd string,
) error {
	pid, err := util.ExecCommandUnblock(fmt.Sprintf("sleep %d; %s", inputInfo.Timeout, cmd))
	if err != nil {
		return fmt.Errorf("execute shell command failed: %v", err)
	}

	// 如果注入命令带timeout参数，将后台执行的故障清理进程的pid号写入数据库表项的private字段。
	if pid != "" {
		opsInfo.Private = fmt.Sprintf("%s,%s", pid, inputInfo.TimeoutStr)
	}
	return nil
}

// cliFaultOpsHandler cli故障注入与清理入口。
func cliFaultOpsHandler(cmd *cobra.Command, _ []string) error {
	// 初始化InputInfo结构体。
	inputInfo, err := data.NewInputInfoByCli(cmd)
	if err != nil {
		return fmt.Errorf("input info init failed: %v", err)
	}

	// 在故障注入和清理场景，opsInfo状态字段均会被初始化为Injected。
	opsInfo, err := data.NewOpsInfo(inputInfo)
	if err != nil {
		return fmt.Errorf("ops info init failed: %v", err)
	}

	if err := OpsHandler(inputInfo, opsInfo); err != nil {
		return fmt.Errorf("handle ops failed(%v)", err)
	}
	return nil
}

// addCobraCliCmdFlags 解析配置文件中命令的flags。
func addCobraCliCmdFlags(faultTypeCmd *cobra.Command, configFlagsMap map[string]interface{}) error {
	for flag, value := range configFlagsMap {
		if value == nil {
			continue
		}

		flagsInfo, ok := value.(map[string]interface{})
		if !ok {
			return errors.New("trans config flags to map[sting]interface{} failed")
		}

		var shortHand string
		if configShortHand := flagsInfo["shorthand"]; configShortHand != nil {
			shortHand = configShortHand.(string)
		}

		var usage string
		if configUsage := flagsInfo["usage"]; configUsage != "" {
			usage = configUsage.(string)
		}
		required := flagsInfo["required"].(bool)

		faultTypeCmd.Flags().StringP(flag, shortHand, "", usage)
		if required {
			if err := faultTypeCmd.MarkFlagRequired(flag); err != nil {
				return fmt.Errorf("mark flag: %s as required failed", flag)
			}
		}
	}
	return nil
}

// addFaultTypeLevelCobraCliCmd 解析故障模式层级到cli。
func addFaultTypeLevelCobraCliCmd(moduleCmd *cobra.Command, fautTypeMap map[string]interface{}, opsType string) error {
	for faultType, value := range fautTypeMap {
		if value == nil {
			continue
		}

		faultInfoMap, ok := value.(map[string]interface{})
		if !ok {
			return errors.New("trans fault info to map[sting]interface{} failed")
		}

		// 故障清理操作不需要添加timeout可选参数。
		var faultTypeLevelCmd *cobra.Command
		if opsType == data.Remove {
			faultTypeLevelCmd = addSubCobraCmd(moduleCmd, faultType, faultInfoMap)
		} else {
			faultTypeLevelCmd = addSubCobraCmd(moduleCmd, faultType, faultInfoMap, true)
		}
		faultTypeLevelCmd.RunE = cliFaultOpsHandler

		flags := faultInfoMap["flags"]
		if flags == nil {
			continue
		}
		configFlagsMap, ok := flags.(map[string]interface{})
		if !ok {
			return errors.New("trans flags to map[sting]interface{} failed")
		}
		err := addCobraCliCmdFlags(faultTypeLevelCmd, configFlagsMap)
		if err != nil {
			return fmt.Errorf("add cobra command :%s flags failed: %v", faultTypeLevelCmd.Use, err)
		}
	}
	return nil
}

// addDomainLevelCobraCliCmd 将cli作用域层级向上添加到env层级。
func addDomainLevelCobraCliCmd(runEnvCmd *cobra.Command, domainsMap map[string]interface{}, opsType string) error {
	for domain, value := range domainsMap {
		if value == nil {
			continue
		}
		domainMap, ok := value.(map[string]interface{})
		if !ok {
			return errors.New("trans domain to map[sting]interface{} failed")
		}
		domainLevelCmd := addSubCobraCmd(runEnvCmd, domain, domainMap)

		faultTypes := domainMap["faulttypes"]
		if faultTypes == nil {
			continue
		}
		faultTypeMap, ok := faultTypes.(map[string]interface{})
		if !ok {
			return errors.New("trans faultTypes info to map[sting]interface{} failed")
		}

		if err := addFaultTypeLevelCobraCliCmd(domainLevelCmd, faultTypeMap, opsType); err != nil {
			return fmt.Errorf("add cobra command :%s sub command failed: %v", domainLevelCmd.Use, err)
		}
	}
	return nil
}

func addSubCobraCmd(
	parentCmd *cobra.Command, use string,
	cmdInfo map[string]interface{},
	timeoutFlag ...bool,
) *cobra.Command {
	newCmd := &cobra.Command{
		Use: use,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	newCmd.SilenceUsage = true

	if shortDesc, ok := cmdInfo["shortdesc"].(string); ok && shortDesc != "" {
		newCmd.Short = shortDesc
	}
	if longDesc, ok := cmdInfo["longdesc"].(string); ok && longDesc != "" {
		newCmd.Long = longDesc
	}

	for _, flag := range timeoutFlag {
		if flag {
			newCmd.PersistentFlags().StringP("timeout", "t", "", "duration of fault")
		}
	}
	parentCmd.AddCommand(newCmd)
	return newCmd
}

// AddCobraCliSubCommand 通过解析配置文件的方式迭代添加子命令和flags，
// 解析的顺序为：env->domain->faultType->flags。
func AddCobraCliSubCommand(opsCmd *cobra.Command) error {
	envs := viper.GetStringMap("env")
	if envs == nil {
		return errors.New("env info is nil in config file")
	}

	for runEnv, value := range envs {
		// env下信息为空，退出本次循环。
		if value == nil {
			continue
		}
		runEnvMap, ok := value.(map[string]interface{})
		if !ok {
			return errors.New("trans env info to map[sting]interface{} failed")
		}
		runEnvLevelCmd := addSubCobraCmd(opsCmd, runEnv, runEnvMap)

		domain := runEnvMap["domain"]
		if domain == nil {
			continue
		}

		domainsMap, ok := domain.(map[string]interface{})
		if !ok {
			return errors.New("trans domain info to map[sting]interface{} failed")
		}
		if err := addDomainLevelCobraCliCmd(runEnvLevelCmd, domainsMap, opsCmd.Use); err != nil {
			return fmt.Errorf("init domain level cobra command failed: %v", err)
		}
	}
	return nil
}
