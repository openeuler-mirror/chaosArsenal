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

	"github.com/spf13/viper"

	"arsenal/models"
	"arsenal/pkg/util"
)

func isInSlice(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

// inputFlagsCheck 读取配置arsenal-spec文件中flags与输入的params做比较，来判断输入参数的正确性。
func inputFlagsCheck(info *InputInfo) error {
	flagsMap := viper.GetStringMap(info.getFlagsKey())
	requiredFlags := make([]string, 0, len(flagsMap))
	optionalFlags := make([]string, 0, len(flagsMap))
	for flag := range flagsMap {
		if viper.GetBool(info.getFlagRequiredKey(flag)) {
			requiredFlags = append(requiredFlags, flag)
		} else {
			optionalFlags = append(optionalFlags, flag)
		}
	}

	// 输入参数大于配置文件中flags的个数，则报错。
	inputFlagsLength := len(info.Flags)
	if inputFlagsLength < len(requiredFlags) || inputFlagsLength > len(flagsMap) {
		return fmt.Errorf("params is invalid, required params list: %s, "+
			"optional params list: %s", requiredFlags, optionalFlags)
	}

	tmpFlags := make(map[string]string)
	for key, value := range info.Flags {
		tmpFlags[key] = value
	}

	// 必选参参数校验，如果params没有包必选参，则报错。
	for _, value := range requiredFlags {
		if _, ok := tmpFlags[value]; !ok {
			return fmt.Errorf("required params is invalid, params list : %s", requiredFlags)
		}
		delete(tmpFlags, value)
	}

	// 可选参参数校验，如果剩余参数不在optionalFlags切片内，则报错。
	for key := range tmpFlags {
		if !isInSlice(optionalFlags, key) {
			return fmt.Errorf("optional params is invalid, params is: %s", optionalFlags)
		}
	}
	return nil
}

// inputParamsCheck 输入参数合法性校验。
func inputParamsCheck(info *InputInfo) error {
	if !viper.IsSet(info.getEnvKey()) {
		return fmt.Errorf("not supported env:%s", info.Env)
	}
	if !viper.IsSet(info.getDomainKey()) {
		return fmt.Errorf("not supported domain:%s in env:%s", info.Domain, info.Env)
	}
	if !viper.IsSet(info.getFaultTypeKey()) {
		return fmt.Errorf("not supported fault type: %s in env:%s domain:%s",
			info.FaultType, info.Env, info.Domain)
	}
	return inputFlagsCheck(info)
}

// NewInputInfoByHTTPRequest 初始化HTTPShellCmdParser结构体中故障注入相关信息。
func NewInputInfoByHTTPRequest(info *models.FaultCreate) (*InputInfo, error) {
	if info.Env == nil {
		return nil, errors.New("param env is nil")
	}
	if info.Domain == nil {
		return nil, errors.New("param domain is nil")
	}
	if info.FaultType == nil {
		return nil, errors.New("param fault mode is nil")
	}

	flags := make(map[string]string)
	for key, value := range info.Params {
		if value != nil {
			flags[key] = *value
		}
	}

	inputInfo := InputInfo{
		InteractiveMode: "http",
		Flags:           flags,
		FlagsString:     util.GetFlagsString(flags),
		FaultType:       *info.FaultType,
		Domain:          *info.Domain,
		Env:             *info.Env,
		OpsType:         Inject,
	}

	if err := inputParamsCheck(&inputInfo); err != nil {
		return nil, err
	}
	// timeout是可选参。
	inputInfo.TimeoutStr = info.Timeout

	if inputInfo.OpsType == Inject {
		if err := inputInfo.SetFaultCheckObject(); err != nil {
			return nil, fmt.Errorf("set fault check object failed: %v", err)
		}
	}
	return &inputInfo, nil
}
