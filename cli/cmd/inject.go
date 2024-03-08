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
	"log"

	"github.com/spf13/cobra"

	"arsenal/internal/parse"
)

// injectCommand 故障注入command。
type injectCommand struct {
	baseCommand
}

// Init 根据配置文件初始化故障注入命令并添加子命令。
func (i *injectCommand) Init() {
	i.cmd = &cobra.Command{
		Use:     "inject",
		Short:   "Inject fault",
		Long:    "Inject fault",
		Example: injectExample(),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			return nil
		},
	}

	if err := parse.AddCobraCliSubCommand(i.cmd); err != nil {
		log.Fatalf("Init inject cli command failed: %v\n", err)
	}
}

func injectExample() string {
	return `arsenal inject os file lost --path /tmp/test.txt`
}
