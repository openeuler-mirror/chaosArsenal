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

package logs

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"arsenal/pkg/util"
)

var Logger *logrus.Logger

// LogrusInit 初始化logrus日志记录器。
func LogrusInit(logLevel string) {
	Logger = logrus.New()

	// 默认的日志等级为“trace”。
	level, err := logrus.ParseLevel("trace")
	if err != nil {
		level = logrus.DebugLevel
		Logger.Warnf("Invalid log level '%s', using default level '%s'", logLevel, level)
	}
	Logger.SetLevel(level)

	// 设定输入日志的时间戳格式。
	Logger.Formatter = &logrus.TextFormatter{
		TimestampFormat:  "2006/01/02 15:04:05",
		DisableSorting:   false,
		QuoteEmptyFields: true,
	}

	// 获取operations.log日志文件的全路径。
	operationsLogPath, err := GetOperationsLogPath()
	if err != nil {
		Logger.Errorf("get arsenal operations log path failed: %v", err)
		return
	}

	// 配置输出日志文件的相关信息。
	Logger.SetOutput(&lumberjack.Logger{
		Filename:   operationsLogPath,
		MaxSize:    1024, // megabytes
		MaxBackups: 10,
		MaxAge:     28,   // days
		Compress:   true, // 是否压缩日志文件
		LocalTime:  true, // 是否使用本地时间
	})
}

// GetOperationsLogPath 获取操作日志文件的全路径。
func GetOperationsLogPath() (string, error) {
	arsenalDir, err := util.GetArsenalDirPath()
	if err != nil {
		return "", fmt.Errorf("get arsenal dir path failed: %v", err)
	}
	return fmt.Sprintf("%s/logs/operations.log", arsenalDir), nil
}

// WithFields 将字段结构添加到日志条目。
func WithFields(fields logrus.Fields) *logrus.Entry {
	return Logger.WithFields(fields)
}

// FormatLog 返回f和v格式化之后的字符串。
func FormatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f := f.(type) {
	case string:
		msg = f
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			// 格式化字符串。
			msg = fmt.Sprintf(msg, v...)
		} else {
			// 不能包含格式化字符串相关字符。
			msg += strings.Repeat(" %v", len(v))
			// 使用 %v 占位符格式化字符串
			msg = fmt.Sprintf(msg, v...)
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		// msg 末尾添加%v占位符。
		msg += strings.Repeat(" %v", len(v))
		// 使用 %v 占位符格式化字符串
		msg = fmt.Sprintf(msg, v...)
	}
	return msg
}

// Trace 在记录器上记录“Trace”级别日志。
func Trace(f interface{}, v ...interface{}) {
	if Logger == nil {
		logrus.Trace(FormatLog(f, v...))
		return
	}
	Logger.Trace(FormatLog(f, v...))
}

// Debug 在记录器上记录“Debug”级别日志。
func Debug(f interface{}, v ...interface{}) {
	if Logger == nil {
		logrus.Debug(FormatLog(f, v...))
		return
	}
	Logger.Debug(FormatLog(f, v...))
}

// Info 在记录器上记录“Info”级别日志。
func Info(f interface{}, v ...interface{}) {
	if Logger == nil {
		logrus.Info(FormatLog(f, v...))
		return
	}
	Logger.Info(FormatLog(f, v...))
}

// Print 在记录器上记录“Info”级别日志。
func Print(f interface{}, v ...interface{}) {
	if Logger == nil {
		logrus.Print(FormatLog(f, v...))
		return
	}
	Logger.Print(FormatLog(f, v...))
}

// Warn 在记录器上记录“Warn”级别日志。
func Warn(f interface{}, v ...interface{}) {
	if Logger == nil {
		logrus.Warn(FormatLog(f, v...))
		return
	}
	Logger.Warn(FormatLog(f, v...))
}

// Warning 在记录器上记录“Warning”级别日志。
func Warning(f interface{}, v ...interface{}) {
	if Logger == nil {
		logrus.Warning(FormatLog(f, v...))
	}
	Logger.Warning(FormatLog(f, v...))
}

// Error 在记录器上记录“Error”级别日志。
func Error(f interface{}, v ...interface{}) {
	if Logger == nil {
		logrus.Error(FormatLog(f, v...))
		return
	}
	Logger.Error(FormatLog(f, v...))
}

// Fatal 在记录器上记录“Fatal”级别日志并且进程的退出状态码为1。
func Fatal(f interface{}, v ...interface{}) {
	if Logger == nil {
		logrus.Fatal(FormatLog(f, v...))
		return
	}
	Logger.Fatal(FormatLog(f, v...))
}

// Panic 在记录器上记录“Panic”级别日志。
func Panic(f interface{}, v ...interface{}) {
	if Logger == nil {
		logrus.Panic(FormatLog(f, v...))
		return
	}
	Logger.Panic(FormatLog(f, v...))
}
