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
	"bytes"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/go-openapi/strfmt"
	// sqlite数据库初始化
	_ "github.com/mattn/go-sqlite3"

	"arsenal/pkg/logs"
	"arsenal/pkg/util"
)

var sqlCreate = `create table if not exists "faults" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"UUID" VARCHAR(32) UNIQUE,
	"interactive_mode" VARCHAR NOT NULL,
	"env" VARCHAR NOT NULL,
	"domain" VARCHAR NOT NULL,
	"fault_type" VARCHAR NOT NULL,
	"object" VARCHAR NOT NULL,
    "flags" VARCHAR,
    "private" VARCHAR,
	"proactive_cleanup" BOOL,
    "status" VARCHAR NOT NULL,
    "inject_time" VARCHAR NOT NULL,
	"update_time" VARCHAR NOT NULL
);`

const (
	cleanupTableCount = 1000000
	uuidSize          = 8
)

var (
	// NoNeedRemoveFaultStatus 不需要做清理操作的数据库表项状态集合。
	NoNeedRemoveFaultStatus = [2]string{
		RemovedStatus,
		SucceedStatus,
	}
	once sync.Once
)

// OpsInfo 数据库中记录故障相关信息的表项。
type OpsInfo struct {
	id int

	// UUID 故障注入时产生的16位16进制字符组成的字符串。
	UUID string

	// InteractiveMode 交互方式。
	InteractiveMode string

	// Env 故障操作运行环境。
	Env string

	// Domain 故障模式所属域。
	Domain string

	// FaultType 故障模式。
	FaultType string

	// Object 故障注入对象。
	Object string

	// Flags 参数组成的字符串，如：--path /bar/foo。
	Flags string

	// Private 故障私有数据区，用于存放后台清理进程pid和延迟时间。
	Private string

	// ProactiveCleanup 故障手动清理标志位，Yes 需要手动清理，No 不需要手动清理。
	ProactiveCleanup bool

	// Status 故障状态信息, 可能存在的状态为有 injected,removed,succeed。
	Status string

	// InjectTime 故障注入的时间，格式为date-time。
	InjectTime strfmt.DateTime

	// UpdateTime 数据库表项更新时间，格式为date-time
	UpdateTime strfmt.DateTime
}

// OpsDatabaseInit 初始化故障信息数据库中的faults表项。
func opsDatabaseInit(db *sql.DB) error {
	var tableName string
	var faultTableExist = false
	rows, err := db.Query("select name from sqlite_master where type = 'table' and name = 'faults'")
	if err != nil {
		return fmt.Errorf("query database fault table failed: %v", err)
	}
	if rows.Next() {
		if err = rows.Scan(&tableName); err != nil {
			rows.Close()
			return err
		}

		if tableName == "faults" {
			faultTableExist = true
		}
	}
	if err = rows.Close(); err != nil {
		return err
	}

	// sqlite数据库中不存在fault表项则创建。
	if !faultTableExist {
		if _, err = db.Exec(sqlCreate); err != nil {
			return err
		}
	}
	return err
}

func opsDatabaseCleanup(db *sql.DB) error {
	// 检查数据库连接是否成功打开。
	if err := db.Ping(); err != nil {
		return fmt.Errorf("database connection failed: %v", err)
	}

	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM faults").Scan(&count); err != nil {
		return fmt.Errorf("query database fault table failed: %v", err)
	}

	if count > cleanupTableCount {
		// 删除status为removed或succeed的相关表项。
		_, err := db.Exec("DELETE FROM faults WHERE status = %s OR status = %s", InjectStatus, RemovedStatus)
		if err != nil {
			return fmt.Errorf("delete fault table failed: %v", err)
		}
	}
	return nil
}

// TODO: 需要考虑数据查找的复杂度，当数据量较大时如何让数据库数据查找的时间最佳？
// OpenOpsInfoDatabase 打开数据库，获取db句柄。
func OpenOpsInfoDatabase() (*sql.DB, error) {
	arsenalPath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("get arsenal absolute path failed: %v", err)
	}

	db, err := sql.Open("sqlite3", fmt.Sprintf("%s/arsenal.db", filepath.Dir(arsenalPath)))
	if err != nil {
		return nil, fmt.Errorf("open database failed: %v", err)
	}

	var initErr error
	// 数据库初始化在每次初始化过程中只执行一次。
	once.Do(func() {
		if initErr = opsDatabaseInit(db); initErr != nil {
			fmt.Println("Ops info data base init failed")
		}
	})
	if initErr != nil {
		return nil, initErr
	}

	// 如果数据库表项有效数据大于100000条，则删除所有status为Removed的表项。
	if err = opsDatabaseCleanup(db); err != nil {
		return nil, fmt.Errorf("cleanup database faults table failed: %v", err)
	}
	return db, nil
}

// generateUUID 产生一个16位16进制字符串作为注入信息唯一表示符。
func generateUUID() (string, error) {
	b := make([]byte, uuidSize)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// NewOpsInfo 根据输入的InputInfo初始化opsInfo结构体，故障默认初始状态为Injected。
func NewOpsInfo(inputInfo *InputInfo) (*OpsInfo, error) {
	var err error
	var newUUID string

	// 如果操作类型为inject，将新创建一个uuid。
	if inputInfo.OpsType == Inject {
		if newUUID, err = generateUUID(); err != nil {
			return nil, fmt.Errorf("generate new UUID failed: %v", err)
		}
	}

	info := OpsInfo{
		UUID:             newUUID,
		InteractiveMode:  inputInfo.InteractiveMode,
		Env:              inputInfo.Env,
		Domain:           inputInfo.Domain,
		FaultType:        inputInfo.FaultType,
		Object:           inputInfo.Object,
		Flags:            inputInfo.FlagsString,
		Private:          "",
		ProactiveCleanup: false,
		Status:           InjectStatus,
		InjectTime:       util.GetCurrentTime(),
		UpdateTime:       util.GetCurrentTime(),
	}
	return &info, nil
}

// InsertOpsInfoIntoDB 将一个opsInfo 故障注入表项插入到数据库。
func (info *OpsInfo) InsertOpsInfoIntoDB(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO faults(UUID, interactive_mode, env, domain, fault_type, object, " +
		"flags, private, proactive_cleanup, status, inject_time, update_time) values(?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return fmt.Errorf("insert ops into data base prepare failed: %v", err)
	}

	_, err = stmt.Exec(info.UUID, info.InteractiveMode, info.Env, info.Domain, info.FaultType, info.Object,
		info.Flags, info.Private, info.ProactiveCleanup, info.Status, info.InjectTime,
		info.UpdateTime)
	if err != nil {
		return fmt.Errorf("stet exec failed: %v", err)
	}
	return nil
}

// PrintOpsInfoByJSONFormat 将opsInfo数组以json格式打印到终端。
func PrintOpsInfoByJSONFormat(infos []OpsInfo) error {
	if infos == nil {
		return errors.New("not found ops info in database")
	}

	bt, err := json.Marshal(infos)
	if err != nil {
		fmt.Println("")
		return err
	}

	var out bytes.Buffer
	if err = json.Indent(&out, bt, "", "\t"); err != nil {
		return err
	}

	if _, err = out.WriteTo(os.Stdout); err != nil {
		return err
	}

	fmt.Printf("\n")
	return nil
}

func getSQLPrepareCmdInfo(keywords map[string]string) (string, []interface{}) {
	queryKey := "select * from faults where"
	queryValue := []interface{}{}
	firstElement := true
	for keyword, value := range keywords {
		// 输入参数为inject-time或update-time需要将'-'替换成'_'。
		keyword = strings.ReplaceAll(keyword, "-", "_")

		// 传入参数为id，数据库中字段key为UUID，需要做特殊处理。
		if keyword == "uuid" {
			keyword = "UUID"
		}

		// 注入时间和清理时间在数据库中做模糊匹配需要特殊处理。
		if keyword == "inject_time" || keyword == "update_time" {
			if firstElement {
				keyword = fmt.Sprintf("%s LIKE ?", keyword)
			} else {
				keyword = fmt.Sprintf("and %s LIKE ?", keyword)
			}
			value = fmt.Sprintf("%%%s%%", value)
		} else {
			if !firstElement {
				keyword = fmt.Sprintf("and %s=?", keyword)
			} else {
				keyword = fmt.Sprintf("%s=?", keyword)
			}
		}
		queryKey = fmt.Sprintf("%s %s", queryKey, keyword)
		queryValue = append(queryValue, value)
		firstElement = false
	}
	return queryKey, queryValue
}

func iterateRows(rows *sql.Rows) ([]OpsInfo, error) {
	var info OpsInfo
	var infos []OpsInfo
	for rows.Next() {
		err := rows.Scan(&info.id, &info.UUID, &info.InteractiveMode, &info.Env, &info.Domain, &info.FaultType,
			&info.Object, &info.Flags, &info.Private, &info.ProactiveCleanup, &info.Status,
			&info.InjectTime, &info.UpdateTime)
		if err != nil {
			return nil, fmt.Errorf("scan db rows failed: %v", err)
		}
		infos = append(infos, info)
	}
	return infos, nil
}

// QueryOpsInfo 根据输入的关键字查询数据库中对应的故障注入表项。
func QueryOpsInfo(keywords map[string]string) ([]OpsInfo, error) {
	if len(keywords) == 0 {
		return nil, errors.New("query keywords is empty")
	}

	db, err := OpenOpsInfoDatabase()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := db.Close(); err != nil {
			logs.Error("closing database error: %v", err)
		}
	}()

	queryKey, queryValue := getSQLPrepareCmdInfo(keywords)
	stmt, err := db.Prepare(queryKey)
	if err != nil {
		return nil, fmt.Errorf("sql prepare failed: %v", err)
	}

	rows, err := stmt.Query(queryValue...)
	if err != nil {
		return nil, fmt.Errorf("query db failed: %v", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			logs.Error("Error closing rows: %v", err)
		}
	}()
	return iterateRows(rows)
}

func (info *OpsInfo) logsOpsInfo() {
	logs.Debug(info.UUID, info.Domain, info.FaultType, info.Object, info.Flags,
		info.ProactiveCleanup, info.Status, info.InjectTime, info.UpdateTime)
}

// FindOpsInfoByUUID 通过uuid查找到对应的故障表项。
func FindOpsInfoByUUID(uuid string) (*OpsInfo, error) {
	infos, err := QueryOpsInfo(map[string]string{"uuid": uuid})
	if err != nil {
		return nil, fmt.Errorf("query ops info failed: %v", err)
	}

	// 通过ID只能唯一表项，0意味着没有找到，大于1意味着找到多个，同一ID不存在多个表项。
	switch len(infos) {
	case 0:
		return nil, fmt.Errorf("no injection fault's information found by uuid: %s", uuid)
	case 1:
		infos[0].logsOpsInfo()
		return &infos[0], err
	default:
		return nil, fmt.Errorf("two entries were found for the same uuid: %s", uuid)
	}
}

// ModifyOpsInfoStatus 修改故障注入表项的状态信息。
func (info *OpsInfo) ModifyOpsInfoStatus(db *sql.DB, newStatus string) error {
	stmt, err := db.Prepare("update faults set status=?,update_time=? where UUID=?")
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(newStatus, util.GetCurrentTime(), info.UUID); err != nil {
		return err
	}
	info.Status = newStatus
	return nil
}

func (info *OpsInfo) recordOpsInfo(opsType string, tmpRecordInfo *OpsInfo) {
	info.UUID = tmpRecordInfo.UUID
	info.Domain = tmpRecordInfo.Domain
	info.FaultType = tmpRecordInfo.FaultType
	info.Flags = tmpRecordInfo.Flags

	// 故障清理操作将ops关键信息返回回去。
	if opsType == Remove {
		info.InteractiveMode = tmpRecordInfo.InteractiveMode
		info.Private = tmpRecordInfo.Private
		info.Status = tmpRecordInfo.Status
	}
}

// OpsInfoInDatabase 根据opsInfo关键字（域名、故障注入对象、状态）在数据中查找匹配表项。
func (info *OpsInfo) OpsInfoInDatabase(opsType string) (bool, error) {
	keywords := map[string]string{"domain": info.Domain, "object": info.Object, "status": info.Status}
	infos, err := QueryOpsInfo(keywords)
	if err != nil {
		return false, fmt.Errorf("query ops info failed: %v", err)
	}

	switch len(infos) {
	case 0:
		if opsType == Remove {
			logs.Warn("No corresponding table entry was found by ops info")
		}
		return false, nil
	case 1:
		// 将查到的相关表项信息写回到info。
		info.recordOpsInfo(opsType, &infos[0])
		info.logsOpsInfo()
	default:
		return true, errors.New("two or more entries were found for the same ops info")
	}
	return true, nil
}
