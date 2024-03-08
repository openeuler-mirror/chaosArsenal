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
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"arsenal/pkg/util"
)

// GetCmdFlags 获取输入cobra cmd的flags组成的map。
func GetCmdFlags(cmd *cobra.Command) map[string]string {
	var flags = make(map[string]string)
	getInputFlags := func(f *pflag.Flag) {
		flags[f.Name] = f.Value.String()
	}

	cmd.Flags().Visit(getInputFlags)
	return flags
}

func getCmdFlagsString(cmd *cobra.Command) string {
	return util.GetFlagsString(GetCmdFlags(cmd))
}

// NewInputInfoByCli 根据cobra cmd传入的基本信息初始化InputInfo结构体。
func NewInputInfoByCli(cmd *cobra.Command) (*InputInfo, error) {
	inputInfo := InputInfo{
		InteractiveMode: "cli",
		Flags:           GetCmdFlags(cmd),
		FlagsString:     getCmdFlagsString(cmd),
		FaultType:       cmd.Name(),
		Domain:          cmd.Parent().Name(),
		Env:             cmd.Parent().Parent().Name(),
		OpsType:         cmd.Parent().Parent().Parent().Name(),
		IsRecorded:      false,
		BlockExecution:  true,
		Cmd:             cmd,
	}
	if err := inputInfo.SetFaultCheckObject(); err != nil {
		return nil, err
	}
	return &inputInfo, nil
}

// NewInputInfoByOpsInfo 根据传入的opsInfo初始化InputInfo结构体，该场景仅在通过uuid清理故障的场景下使用。
func NewInputInfoByOpsInfo(info *OpsInfo) *InputInfo {
	return &InputInfo{
		OpsType:     Remove,
		FlagsString: info.Flags,
		FaultType:   info.FaultType,
		Env:         info.Env,
		Domain:      info.Domain,
	}
}
