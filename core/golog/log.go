// Copyright 2016 The kingshard Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package golog

import (
	"fmt"
	"github.com/flike/kingshard/core/logger"
	"go.uber.org/zap"
	"sync"
)

// log level, from low to high, more high means more serious
const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

const (
	Ltime  = 1 << iota //time format "2006/01/02 15:04:05"
	Llevel             //[Trace|Debug|Info...]
)

var LevelName [6]string = [6]string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

const (
	LogOn          = "on"
	LogOff         = "off"
	TimeFormat     = "2006/01/02 15:04:05"
	maxBufPoolSize = 16
)

type Logger struct {
	sync.Mutex

	level int
	flag  int
}

// new a logger with specified handler and flag
func New(flag int) *Logger {
	var l = new(Logger)

	l.level = LevelInfo

	l.flag = flag

	return l
}

// new a default logger with specified handler and flag: Ltime|Lfile|Llevel
func NewDefault() *Logger {
	return New(Ltime | Llevel)
}

var std = NewDefault()

// set log level, any log level less than it will not log
func (l *Logger) SetLevel(level int) {
	l.level = level
}

func (l *Logger) Level() int {
	return l.level
}

func SetLevel(level int) {
	std.SetLevel(level)
}

func StdLogger() *Logger {
	return std
}

func GetLevel() int {
	return std.level
}

// 全局变量
var GlobalSysLogger *Logger = StdLogger()
var GlobalSqlLogger *Logger = GlobalSysLogger

func (l *Logger) Write(p []byte) (n int, err error) {
	logger.Info("info", zap.String("module", "web"), zap.String("method", "api"), zap.Any("msg", string(p)))
	return len(p), nil
}

func OutputSql(state string, format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)

	logger.Info("sql", zap.String("state", state), zap.String("sql", s))
}

func Trace(module string, method string, msg string, reqId uint32, args ...interface{}) {
	if LevelTrace < GlobalSysLogger.Level() {
		return
	}

	logger.Info("Trace", zap.String("module", module), zap.String("method", method), zap.String("msg", msg), zap.Uint32("reqId", reqId), zap.Any("args", args))
}

func Debug(module string, method string, msg string, reqId uint32, args ...interface{}) {
	if LevelDebug < GlobalSysLogger.Level() {
		return
	}

	logger.Debug("debug", zap.String("module", module), zap.String("method", method), zap.String("msg", msg), zap.Uint32("reqId", reqId), zap.Any("args", args))
}

func Info(module string, method string, msg string, reqId uint32, args ...interface{}) {
	if LevelInfo < GlobalSysLogger.Level() {
		return
	}

	logger.Info("info", zap.String("module", module), zap.String("method", method), zap.String("msg", msg), zap.Uint32("reqId", reqId), zap.Any("args", args))
}

func Warn(module string, method string, msg string, reqId uint32, args ...interface{}) {
	if LevelWarn < GlobalSysLogger.Level() {
		return
	}

	logger.Warn("warn", zap.String("module", module), zap.String("method", method), zap.String("msg", msg), zap.Uint32("reqId", reqId), zap.Any("args", args))
}

func Error(module string, method string, msg string, reqId uint32, args ...interface{}) {
	if LevelError < GlobalSysLogger.Level() {
		return
	}

	logger.Error("error", zap.String("module", module), zap.String("method", method), zap.String("msg", msg), zap.Uint32("reqId", reqId), zap.Any("args", args))
}

func Fatal(module string, method string, msg string, reqId uint32, args ...interface{}) {
	if LevelFatal < GlobalSysLogger.Level() {
		return
	}

	logger.Fatal("fatal", zap.String("module", module), zap.String("method", method), zap.String("msg", msg), zap.Uint32("reqId", reqId), zap.Any("args", args))
}
