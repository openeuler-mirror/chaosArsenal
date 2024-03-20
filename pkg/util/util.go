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

package util

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
)

// ArrayToString 将一个数组转换成空格隔开的字符串。
func ArrayToString(arr []string) string {
	var result string
	for _, value := range arr {
		result += value + " "
	}
	return result
}

// GetCurrentTime 获取格式为date-time的时间戳。
func GetCurrentTime() strfmt.DateTime {
	return strfmt.DateTime(time.Now().Local())
}

// FileIsExist 判断文件是否存在。
func FileIsExist(path string) bool {
	_, ret := os.Stat(path)
	return ret == nil
}

// ExecCommandBlock 阻塞执行shell命令，默认的超时时间为5s。
func ExecCommandBlock(shellCmd string, timeout ...uint64) (string, error) {
	var ctx context.Context
	const defaultTimeout = 10
	var cancel context.CancelFunc
	if len(timeout) > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), (time.Duration(timeout[0]))*(time.Second))
	} else {
		ctx, cancel = context.WithTimeout(context.Background(), defaultTimeout*time.Second)
	}
	defer cancel()

	cmd := exec.CommandContext(ctx, "/bin/bash", "-c", shellCmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		return "", fmt.Errorf("execute command: %s timeout, default %ds", shellCmd, defaultTimeout)
	}
	return out.String(), err
}

// ExecCommandUnblock 非阻塞执行shell命令，命令执行过程中可能出现错误，例如命令不存在或者权限不足等问题。
func ExecCommandUnblock(shellCmd string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", shellCmd)
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("exec command:%v error:%v", shellCmd, err)
	}
	return strconv.Itoa(cmd.Process.Pid), nil
}

// GetArsenalAbsPath 获取arsenal二进制文件的全路径。
func GetArsenalAbsPath() (string, error) {
	path, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return abs, nil
}

// GetArsenalDirPath 获取arsenal二进制文件所在文件夹的路径。
func GetArsenalDirPath() (string, error) {
	arsenalPath, err := GetArsenalAbsPath()
	if err != nil {
		return "", err
	}
	return filepath.Dir(arsenalPath), nil
}

// GetAtomicFaultInjectionExePath 获取底层原子故障注入可执行文件的全路径。
func GetAtomicFaultInjectionExePath(executorName string) (string, error) {
	// 获取可执行文件的路径。
	arsenalDir, err := GetArsenalDirPath()
	if err != nil {
		return "", fmt.Errorf("get arsenal dir path failed: %v", err)
	}

	exePath := fmt.Sprintf("%s/bin/%s", arsenalDir, executorName)
	// 判断底层原子故障注入可执行文件文件是否存在。
	if !FileIsExist(exePath) {
		return "", fmt.Errorf("executor: %s not exist", exePath)
	}
	return exePath, nil
}

func orderFlagsMapKey(inputMap map[string]string) []string {
	// 将map的key转换成切片。
	keys := make([]string, 0, len(inputMap))
	for key := range inputMap {
		keys = append(keys, key)
	}

	// 对切片进行排序。
	sort.Strings(keys)
	return keys
}

// GetFlagsString 拼接flags组成的字符串，格式为--$key value。
func GetFlagsString(flags map[string]string) string {
	// 将输入map的key进行排序，解决flags字符串乱序问题。
	orderKeys := orderFlagsMapKey(flags)

	var flagsString string
	for _, key := range orderKeys {
		// timeout参数不传递给底层工具，需要从参数中剔除。
		if key == "timeout" {
			continue
		}
		flagsString += fmt.Sprintf("--%s %s ", key, flags[key])
	}

	const minimumFlagStringLength = 3
	if len(flagsString) > minimumFlagStringLength {
		flagsString = strings.TrimSpace(flagsString)
	}
	return flagsString
}
