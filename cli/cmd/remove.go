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

package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"arsenal/internal/parse"
	"arsenal/pkg/data"
	"arsenal/pkg/logs"
	executor "arsenal/pkg/run"
)

// removeCommand 故障清理command。
type removeCommand struct {
	baseCommand
}

// Init 根据配置文件初始化故障清理命令并添加子命令，故障清理包含两种方式：
// 1、通过传入uuid的方式清理故障。
// 2、通过传入完整cli命令行参数的方式清理故障。
func (r *removeCommand) Init() {
	r.cmd = &cobra.Command{
		Use:     "remove",
		Short:   "Remove fault that have been injected",
		Long:    "Remove fault that have been injected",
		Example: removeExample(),
		RunE: func(cmd *cobra.Command, args []string) error {
			UUID, _ := cmd.Flags().GetString("uuid")
			if UUID == "" {
				return cmd.Help()
			}
			return removeFaultViaUUID(UUID)
		},
	}

	r.cmd.Flags().StringP("uuid", "u", "", "Injected fault's uuid")
	if err := parse.AddCobraCliSubCommand(r.cmd); err != nil {
		log.Fatalf("Init remove cli command failed: %v\n", err)
	}
}

func removeExample() string {
	return "# Remove fault by full cli shell command\n" +
		"arsenal remove os file lost --path /tmp/test.txt\n" +
		"# Remove fault by uuid\n" +
		"arsenal remove --uuid 65c7e772d3b65a9d"
}

// removeFaultViaUUID 通过传入的UUID来清除对应故障。
func removeFaultViaUUID(uuid string) error {
	info, err := data.FindOpsInfoByUUID(uuid)
	if err != nil {
		return fmt.Errorf("get ops info via uuid failed: %v", err)
	}

	// TODO: 当前cli只能清理cli模式的故障，http需要与服务交互才能清理对应故障。
	if info.InteractiveMode != parse.Cli {
		return fmt.Errorf("cli remove not support %s interactive mode's fault", info.InteractiveMode)
	}

	// 如果故障状态已经处于不需要清理的状态，报错返回。
	for _, status := range data.NoNeedRemoveFaultStatus {
		if info.Status == status {
			return errors.New("the fault has been removed or does not need to be removed")
		}
	}

	// 如果故障不需要手动清理直接返回出错。
	if !info.ProactiveCleanup {
		return fmt.Errorf("faults with UUID: %s do not need to be removed", uuid)
	}

	inputInfo := data.NewInputInfoByOpsInfo(info)
	if err := inputInfo.SetAtomicFaultInjectionExePath(fmt.Sprintf("arsenal-%s", info.Env)); err != nil {
		return fmt.Errorf("get the atomic fault injection executable failed by ops info: %v", err)
	}
	if err = inputInfo.SetFullShellCommandByInputInfo(); err != nil {
		return fmt.Errorf("set full shell command failed: %v", err)
	}
	inputInfo.SetBlockExecutionFlag()

	if err := parse.KillTimeoutBackgroundProcess(inputInfo, info); err != nil {
		return fmt.Errorf("kill fault clean background process failed: %v", err)
	}
	logs.Info("Remove fault shell command by uuid: ", inputInfo.ShellCmd)

	_, err = executor.ShellCmd(inputInfo.OpsType, inputInfo.BlockExecution, inputInfo.ShellCmd)
	if err != nil {
		return fmt.Errorf("execute shell command failed: %v", err)
	}

	if err = parse.UpdateDB(inputInfo, info); err != nil {
		return fmt.Errorf("update db failed after remove fault by uuid: %v", err)
	}
	return nil
}
