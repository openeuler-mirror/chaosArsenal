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

import "github.com/spf13/cobra"

// command cobra cli命令相关接口。
type command interface {
	// Init cobra命令初始化函数。
	Init()

	CobraCmd() *cobra.Command
}

type baseCommand struct {
	cmd *cobra.Command
}

// CobraCmd 返回cobra command指针。
func (b *baseCommand) CobraCmd() *cobra.Command {
	return b.cmd
}

// AddCommand 为baseCommand添加子命令。
func (b *baseCommand) AddCommand(child command) {
	child.Init()
	childCmd := child.CobraCmd()
	childCmd.SilenceUsage = true
	childCmd.DisableFlagsInUseLine = true
	b.CobraCmd().AddCommand(childCmd)
}
