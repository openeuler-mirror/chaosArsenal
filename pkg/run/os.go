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

package run

import (
	"fmt"
	"strings"

	"arsenal/pkg/util"
)

// ShellCmd 执行shell命令，如果是非阻塞执行需要先执行prepare命令做前置检查。
func ShellCmd(opsType string, blockExecutionFlag bool, fullShellCmd string) (string, error) {
	if blockExecutionFlag {
		if result, err := util.ExecCommandBlock(fullShellCmd); err != nil {
			return "", fmt.Errorf("execute command [%s] failed, Error [%s], Result [%s]",
				fullShellCmd, err, strings.ReplaceAll(result, "\n", ""))
		}
	} else {
		prepareCmd := strings.ReplaceAll(fullShellCmd, opsType, "prepare")
		if result, err := util.ExecCommandBlock(prepareCmd); err != nil {
			return "", fmt.Errorf("execute command [%s] failed, Error [%s], Result [%s]",
				prepareCmd, err, strings.ReplaceAll(result, "\n", ""))
		}

		pid, err := util.ExecCommandUnblock(fullShellCmd)
		if err != nil {
			return "", err
		}
		return pid, nil
	}
	return "", nil
}
