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
	"fmt"

	"github.com/spf13/cobra"
)

const ArsenalVersion = "1.0.0"

// versionCommand arsenal版本信息命令。
type versionCommand struct {
	baseCommand
}

// Init 初始化故障注入工具版本查询命令。
func (v *versionCommand) Init() {
	v.cmd = &cobra.Command{
		Use:     "version",
		Short:   "Arsenal version information",
		Long:    "Arsenal version information",
		Example: versionExample(),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("chaos-arsenal version", ArsenalVersion)
		},
	}
}

func versionExample() string {
	return `arsenal version`
}
