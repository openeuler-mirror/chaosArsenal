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

package parse

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-openapi/runtime/middleware"

	"arsenal/models"
	"arsenal/pkg/data"
	"arsenal/pkg/logs"
	"arsenal/pkg/util"
)

func init() {
	var newHandlerOps = httpHandlerOperations{
		interactiveMode: HTTP,
	}
	Add(newHandlerOps.interactiveMode, &newHandlerOps)
	newHandlerOps.idChanel = make(map[string]chan bool)
	newHandlerOps.mutex = sync.Mutex{}
}

type httpHandlerOperations struct {
	interactiveMode string
	idChanel        map[string]chan bool
	mutex           sync.Mutex
}

func (c *httpHandlerOperations) addFaultCleanBackgroundProcessChan(id string, quit chan bool) {
	c.mutex.Lock()
	c.idChanel[id] = quit
	c.mutex.Unlock()
}

func (c *httpHandlerOperations) deleteFaultCleanBackgroundProcessChan(id string) {
	c.mutex.Lock()
	close(c.idChanel[id])
	delete(c.idChanel, id)
	c.mutex.Unlock()
}

// OperationsCheck 对http输入命令进行合法性检查。
func (c *httpHandlerOperations) OperationsCheck(inputInfo *data.InputInfo, opsInfo *data.OpsInfo) error {
	return inputInfo.CheckOps(opsInfo)
}

// SetRecordedFlag 根据http输入参数初始化的opsInfo在数据库中查找相应表项，如果表项存在，
// 将inputInfo成员IsRecorded置true。
func (c *httpHandlerOperations) SetRecordedFlag(inputInfo *data.InputInfo, opsInfo *data.OpsInfo) error {
	// 只有注入操作才需要查找输入故障信息在数据库中是否存在，
	// 清理操作在根据uuid获取opsInfo时就可以确定故障信息可以在数据库中查到。
	if inputInfo.OpsType != data.Inject {
		return nil
	}

	if err := inputInfo.SetRecordedFlag(opsInfo); err != nil {
		return fmt.Errorf("get recorded flag failed: %v", err)
	}
	return nil
}

// SetTimeoutValue 从输入的http参数中获取timeout需要延迟的时间，单位为秒。
func (c *httpHandlerOperations) SetTimeoutValue(inputInfo *data.InputInfo) {
	if inputInfo.TimeoutStr == "" {
		return
	}
	timeout, err := inputInfo.ParseTimeout(inputInfo.TimeoutStr)
	if err != nil {
		logs.Fatal("%v", err)
	}
	inputInfo.Timeout = timeout
}

// SetShellCmd 根据http输入的参数拼装底层原子故障注入的shell命令。
func (c *httpHandlerOperations) SetShellCmd(inputInfo *data.InputInfo, opsInfo *data.OpsInfo) error {
	// 命令示例：arsenal-os prepare process caton --pid 10 --interval 10
	if inputInfo.OpsType == data.Remove {
		inputInfo.FlagsString = opsInfo.Flags
	}

	if err := inputInfo.SetFullShellCommandByInputInfo(); err != nil {
		return fmt.Errorf("get full shell command failed: %v", err)
	}
	return nil
}

// KillFaultCleanBackgroundProcess 向http故障延迟清理协程发送退出信号。
func (c *httpHandlerOperations) KillFaultCleanBackgroundProcess(
	_ *data.InputInfo,
	opsInfo *data.OpsInfo,
) error {
	quit := c.idChanel[opsInfo.UUID]
	if quit != nil {
		quit <- true
		c.deleteFaultCleanBackgroundProcessChan(opsInfo.UUID)
	}
	return nil
}

// RunFaultCleanBackgroundProcess 运行故障延时清理协程。
func (c *httpHandlerOperations) RunFaultCleanBackgroundProcess(
	inputInfo *data.InputInfo,
	opsInfo *data.OpsInfo,
	cmd string,
) error {
	quit := make(chan bool)
	c.addFaultCleanBackgroundProcessChan(opsInfo.UUID, quit)
	go func() {
		opsInfo.Private = inputInfo.TimeoutStr
		var i uint64
		for i = 0; i < inputInfo.Timeout; i++ {
			time.Sleep(1 * time.Second)

			// 每隔1s钟检查退出信号。
			select {
			case <-quit:
				logs.Info("receive exit signal in advance, goroutine exit")
				return
			default:
				// 继续做延时。
			}
		}
		c.deleteFaultCleanBackgroundProcessChan(opsInfo.UUID)
		if _, err := util.ExecCommandUnblock(cmd); err != nil {
			logs.Error("execute shell command(%s) failed(%v) ", cmd, err)
			return
		}
	}()
	return nil
}

func make500ErrorResponse(errInfo string) models.Error500Response {
	var code models.Code = 500
	message := models.Message(errInfo)

	return models.Error500Response{
		Code:    &code,
		Message: &message,
	}
}

func makeNr200Response(errInfo string) models.Nr200Response {
	var code models.Code = 200
	message := models.Message(errInfo)

	return models.Nr200Response{
		Code:    &code,
		Message: &message,
	}
}

// makeInfos200Response 将从数据库中查找的信息转换成openApi定义的返回值类型。
func makeInfos200Response(infos []data.OpsInfo) models.Infos200Response {
	pointerInfos := make([]*models.OpsInfo, len(infos))
	for i, info := range infos {
		var modelsInfo models.OpsInfo
		modelsInfo.UUID = info.UUID
		modelsInfo.Domain = info.Domain
		modelsInfo.FaultType = info.FaultType
		modelsInfo.Flags = info.Flags
		modelsInfo.Private = info.Private
		modelsInfo.ProactiveCleanup = info.ProactiveCleanup
		modelsInfo.Status = info.Status
		modelsInfo.InjectTime = info.InjectTime
		modelsInfo.UpdateTime = info.UpdateTime
		pointerInfos[i] = &modelsInfo
	}

	var code models.Code = 200
	return models.Infos200Response{
		Code:  &code,
		Infos: pointerInfos,
	}
}

// GetFaultsHandler 通过url传入的Query关键字，在数据库中查找故障注入信息并返回。
func GetFaultsHandler(request *http.Request) middleware.Responder {
	queryParams := request.URL.Query()
	queryMap := make(map[string]string)
	for key, values := range queryParams {
		if len(values) > 0 {
			queryMap[key] = values[0]
		}
	}

	matchResult, err := data.QueryOpsInfo(queryMap)
	if err != nil {
		response := make500ErrorResponse(fmt.Sprintf("query ops info failed: %v", err))
		return middleware.Error(int(*(response.Code)), response)
	}

	response := makeInfos200Response(nil)
	if len(matchResult) != 0 {
		response = makeInfos200Response(matchResult)
	}
	return middleware.Error(int(*(response.Code)), response)
}

func commonHandler(inputInfo *data.InputInfo, opsInfo *data.OpsInfo) middleware.Responder {
	if err := OpsHandler(inputInfo, opsInfo); err != nil {
		response := make500ErrorResponse(err.Error())
		return middleware.Error(int(*(response.Code)), response)
	}

	var id models.ID
	if inputInfo.OpsType == data.Inject {
		id = models.ID(opsInfo.UUID)
	}
	response := makeNr200Response("success")
	response.ID = &id
	return middleware.Error(int(*(response.Code)), response)
}

// getOpsInfo 获取故障注入或清理的opsInfo对象：
// 如果是故障注入，根据传入InputInfo重新初始化一个opsInfo类型的结构体；
// 如果是故障清理，则直接从数据库中读取对应操作信息；
func getOpsInfo(inputInfo *data.InputInfo) (*data.OpsInfo, error) {
	switch inputInfo.OpsType {
	case data.Inject:
		opsInfo, err := data.NewOpsInfo(inputInfo)
		if err != nil {
			return nil, fmt.Errorf("int ops info failed: %v", err)
		}
		return opsInfo, nil
	case data.Remove:
		// 故障清理操作从数据库中通过uuid获取对应的opsInfo信息。
		info, err := data.FindOpsInfoByUUID(inputInfo.UUID)
		if err != nil {
			return nil, fmt.Errorf("get ops info via uuid failed: %v", err)
		}
		// 在数据库中根据id已经查找对应表项，直接将Recorded标志位置true。
		inputInfo.IsRecorded = true
		return info, nil
	default:
		break
	}
	return nil, fmt.Errorf("unsupported operation when init ops info: %s", inputInfo.OpsType)
}

// DeleteFaultsIDHandler 故障清理处理函数。
func DeleteFaultsIDHandler(inputUUID string) middleware.Responder {
	var inputInfo data.InputInfo

	// 初始化InputInfo结构体。
	inputInfo.UUID = inputUUID
	inputInfo.OpsType = data.Remove
	inputInfo.InteractiveMode = HTTP

	// 初始化OpsInfo结构体。
	opsInfo, err := getOpsInfo(&inputInfo)
	if err != nil {
		response := make500ErrorResponse(err.Error())
		return middleware.Error(int(*(response.Code)), response)
	}

	// cli交互模式注入的故障，需要用cli的方式清理故障。
	if opsInfo.InteractiveMode == Cli {
		errInfo := fmt.Sprintf("uuid: %s injected interactive mode is cli, "+
			"use the cli interaction to remove the fault", opsInfo.UUID)
		response := make500ErrorResponse(errInfo)
		return middleware.Error(int(*(response.Code)), response)
	}

	// inputInfo基本信息初始化。
	inputInfo.Env = opsInfo.Env
	inputInfo.Domain = opsInfo.Domain
	inputInfo.FaultType = opsInfo.FaultType

	// 设置原子故障注入能力可执行文件的全路径。
	if err := inputInfo.SetAtomicFaultInjectionExePath(fmt.Sprintf("arsenal-%s", opsInfo.Env)); err != nil {
		response := make500ErrorResponse(err.Error())
		return middleware.Error(int(*(response.Code)), response)
	}
	return commonHandler(&inputInfo, opsInfo)
}

// PostFaultsHandler 故障注入处理函数。
func PostFaultsHandler(info *models.FaultCreate) middleware.Responder {
	// 初始化InputInfo结构体。
	inputInfo, err := data.NewInputInfoByHTTPRequest(info)
	if err != nil {
		response := make500ErrorResponse(err.Error())
		return middleware.Error(int(*(response.Code)), response)
	}

	// 初始化OpsInfo结构体。
	opsInfo, err := getOpsInfo(inputInfo)
	if err != nil {
		response := make500ErrorResponse(err.Error())
		return middleware.Error(int(*(response.Code)), response)
	}
	return commonHandler(inputInfo, opsInfo)
}
