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
	"flag"
	"fmt"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"

	"arsenal/pkg/logs"
	"arsenal/pkg/util"
	"arsenal/restapi"
	"arsenal/restapi/operations"
)

const (
	Umask       = 27
	PidFilePerm = 0644
	LogFilePerm = 0644
	defaultPort = 9095
)

// serverCommand http服务相关命令。
type serverCommand struct {
	baseCommand
}

// Init 初始化http服务命令并添加子命令。
func (s *serverCommand) Init() {
	s.cmd = &cobra.Command{
		Use:     "server",
		Short:   "Http server operations",
		Long:    "Http server operations",
		Example: serverExample(),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			return nil
		},
	}

	// 添加http server start命令
	httpServerStartCmd := &cobra.Command{
		Use:   "start",
		Short: "Start http server",
		Long:  "Start http server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return httpServerHandler(cmd)
		},
		SilenceUsage: true,
	}
	httpServerStartCmd.Flags().Int("port", defaultPort, "Http server port")
	httpServerStartCmd.Flags().StringP("host", "", "localhost", "Http server host ip")
	s.cmd.AddCommand(httpServerStartCmd)

	// 添加http server stop命令
	httpServerStopCmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop http server",
		Long:  "Stop http server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return httpServerHandler(cmd)
		},
		SilenceUsage: true,
	}
	signal := "signal"
	httpServerStopCmd.Flags().String(signal, "", "Send signal to the daemon: quit, stop, reload")
	if err := httpServerStopCmd.MarkFlagRequired(signal); err != nil {
		logs.Error("% mark flag %s required failed: %v", httpServerStopCmd.Name(), signal, err)
	}
	s.cmd.AddCommand(httpServerStopCmd)
}

func serverExample() string {
	return "# Start arsenal http server\n" +
		"arsenal server start --host 192.168.0.131 --port 9095\n" +
		"# Stop arsenal http server\n" +
		"arsenal server stop --signal stop"
}

func initDaemonContext() (*daemon.Context, error) {
	arsenalAbsDir, err := util.GetArsenalDirPath()
	if err != nil {
		return nil, fmt.Errorf("get arsenal work directory path failed: %v", err)
	}

	operationsLogPath, err := logs.GetOperationsLogPath()
	if err != nil {
		return nil, fmt.Errorf("get arsenal operations log path failed: %v", err)
	}

	serverContext := &daemon.Context{
		PidFileName: fmt.Sprintf("%s/server.pid", arsenalAbsDir),
		PidFilePerm: PidFilePerm,
		LogFileName: operationsLogPath,
		LogFilePerm: LogFilePerm,
		WorkDir:     arsenalAbsDir,
		Umask:       Umask,
		Args:        []string{},
	}
	return serverContext, nil
}

func httpServerHandler(cmd *cobra.Command) error {
	flag.Parse()
	signal, _ := cmd.Flags().GetString("signal")
	daemon.AddCommand(daemon.StringFlag(&signal, "quit"), syscall.SIGQUIT, termHandler)
	daemon.AddCommand(daemon.StringFlag(&signal, "stop"), syscall.SIGTERM, termHandler)
	daemon.AddCommand(daemon.StringFlag(&signal, "reload"), syscall.SIGHUP, reloadHandler)

	serverContext, err := initDaemonContext()
	if err != nil {
		return fmt.Errorf("init server daemon context failed: %v", err)
	}

	if len(daemon.ActiveFlags()) > 0 {
		d, err := serverContext.Search()
		if err != nil {
			return fmt.Errorf("unable send signal to the daemon: %v", err)
		}
		return daemon.SendCommands(d)
	}

	d, err := serverContext.Reborn()
	if err != nil {
		return fmt.Errorf("reborn failed: %v", err)
	}
	if d != nil {
		return nil
	}
	defer func() {
		if err := serverContext.Release(); err != nil {
			logs.Error("release daemon failed: %v", err)
		}
	}()

	logs.Info("- - - http server daemon started - - -")
	go func() {
		if err := serverHTTP(cmd); err != nil {
			logs.Fatal("server http failed: %v", err)
		}
	}()

	if err := daemon.ServeSignals(); err != nil {
		return fmt.Errorf("serve signals error: %v", err)
	}
	return nil
}

var (
	stop = make(chan struct{})
	done = make(chan struct{})
)

func serverHTTP(cmd *cobra.Command) error {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return fmt.Errorf("loads swagger spec failed: %v", err)
	}

	api := operations.NewChaosArsenalFaultInjectionAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer func() {
		if err := server.Shutdown(); err != nil {
			logs.Error("shut down arsenal server failed")
		}
	}()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "chaosArsenal fault injection"
	parser.LongDescription = "chaosArsenal fault injection module API specification"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			return fmt.Errorf("parser add group failed: %v", err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		if code == 1 {
			return fmt.Errorf("parser error failed: %v", err)
		}
	}
	// cli没有输入自定义端口号，设定默认端口为9095。
	if !cmd.Flags().Changed("port") {
		server.Port = 9095
	}

	server.ConfigureAPI()

	go func() {
		if err := server.Serve(); err != nil {
			logs.Fatal("server failed: %v", err)
		}
	}()

	waitSignal()
	return nil
}

func waitSignal() {
LOOP:
	for {
		time.Sleep(time.Second) // this is work to be done by worker.
		select {
		case <-stop:
			break LOOP
		default:
		}
	}
	done <- struct{}{}
}

func termHandler(sig os.Signal) error {
	stop <- struct{}{}
	if sig == syscall.SIGQUIT {
		<-done
	}
	return daemon.ErrStop
}

func reloadHandler(_ os.Signal) error {
	log.Println("server reloaded has not yet been implemented")
	return nil
}
