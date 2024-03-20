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
	"github.com/spf13/pflag"

	"arsenal/pkg/data"
)

// queryCommand 故障信息查询command。
type queryCommand struct {
	baseCommand
}

// Init 初始化故障查询命令并添加子命令。
func (q *queryCommand) Init() {
	q.cmd = &cobra.Command{
		Use:     "query",
		Short:   "Query fault's information",
		Long:    "Query for information about faults that have been injected",
		Example: queryExample(),
		RunE: func(cmd *cobra.Command, args []string) error {
			infos, err := data.QueryOpsInfo(data.GetCmdFlags(cmd))
			if err != nil {
				return fmt.Errorf("query ops info failed: %v", err)
			}
			if err = data.PrintOpsInfoByJSONFormat(infos); err != nil {
				return fmt.Errorf("print ops info by JSON format failed: %v", err)
			}
			return nil
		},
	}
	q.addFlags()
}

func (q *queryCommand) addFlags() {
	q.cmd.Flags().StringP("uuid", "u", "", "Injected fault's uuid")
	q.cmd.Flags().StringP("domain", "", "", "Injected fault's domain")
	q.cmd.Flags().StringP("fault-type", "f", "", "Injected fault's type")
	q.cmd.Flags().StringP("object", "o", "", "Injected fault's object")
	q.cmd.Flags().StringP("inject-time", "", "", "Injected fault's time")
	q.cmd.Flags().StringP("update-time", "", "", "Update fault's info time")
	q.cmd.Flags().StringP("status", "s", "", "Injected fault's status")
}

func queryExample() string {
	return "# Query all removed faults information\n" +
		"arsenal query --status removed\n" +
		"# Query all faults information by domain\n" +
		"arsenal query --domain file\n" +
		"# Query all faults information by fault type\n" +
		"arsenal query --domain file --fault-type readonly\n" +
		"# Query all faults information by object\n" +
		"arsenal query --object '/tmp/test.txt'\n" +
		"# Query fault information by fault uuid\n" +
		"arsenal query --uuid 5861d089e8b4de29\n" +
		"# Query fault information by injected time\n" +
		"arsenal query --inject-time '2023-06-25T02:10:20'\n" +
		"# Query fault information by update time\n" +
		"arsenal query --update-time '2023-06-25T02:10:20'"
}

// GetCmdFlagsName 获取cobra输入命令中的所有flag名组成的数组。
func (q *queryCommand) GetCmdFlagsName() []string {
	var flags []string
	getInputFlags := func(f *pflag.Flag) {
		flags = append(flags, f.Name)
	}
	q.cmd.Flags().Visit(getInputFlags)
	return flags
}
