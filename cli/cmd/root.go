/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"arsenal/pkg/logs"
)

var rootCmd = &cobra.Command{
	Use:   "arsenal",
	Short: "An easy to use and powerful RAS toolkit for Linux",
	Long:  "An easy to use and powerful RAS toolkit for Linux",
}

// MakeRootCmd 构造cli root命令，为root命令添加子命令。
func MakeRootCmd() *cobra.Command {
	cobraInit()

	rootBaseCmd := &baseCommand{
		cmd: rootCmd,
	}

	// add version command
	rootBaseCmd.AddCommand(&versionCommand{})

	// add inject reference command
	rootBaseCmd.AddCommand(&injectCommand{})

	// add remove reference command
	rootBaseCmd.AddCommand(&removeCommand{})

	// add query reference command
	rootBaseCmd.AddCommand(&queryCommand{})

	// add server reference command
	rootBaseCmd.AddCommand(&serverCommand{})

	return rootBaseCmd.CobraCmd()
}

// cobraInit viper配置文件读取初始化操作。
func cobraInit() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	initConfig()
	logs.LogrusInit("Debug")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	configName := fmt.Sprintf("arsenal-spec-%s", ArsenalVersion)
	viper.SetConfigName(configName)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	// read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Reading config file: %s failed, Err: %s", viper.ConfigFileUsed(), err)
	}
}
