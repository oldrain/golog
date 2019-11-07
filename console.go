// Copyright 2019 golog Author. All Rights Reserved.
// License that can be found in the LICENSE file.

package golog

import (
	"fmt"
	"time"
)

// Console output
type ConsoleLog struct {
	cfg *Config
	module string

	head string
	tail string
}

var (
	// Colorful characters
	consoleLevelStrMap = map[int]string{
		LevelFatal: "\033[35mFATAL\033[0m",
		LevelError: "\033[31mERROR\033[0m",
		LevelWarn: "\033[33mWARN\033[0m",
		LevelInfo: "\033[32mINFO\033[0m",
		LevelDebug: "\033[36mDEBUG\033[0m",
		LevelTrace: "\033[34mTRACE\033[0m",
	}
)

func (logger *ConsoleLog) start() {
	// do nothing
}

func (logger *ConsoleLog) AppendHead(args ...interface{}) {
	logger.setHead(logger.append(logger.head, args...))
}

func (logger *ConsoleLog) AppendTail(args ...interface{}) {
	logger.setTail(logger.append(logger.tail, args...))
}

func (logger *ConsoleLog) EraseHead() {
	logger.setHead(EmptyString)
}

func (logger *ConsoleLog) EraseTail() {
	logger.setTail(EmptyString)
}

func (logger *ConsoleLog) Erase() {
	logger.EraseHead()
	logger.EraseTail()
}

func (logger *ConsoleLog) Trace(args ...interface{}) {
	logger.logging(LevelTrace, args...)
}

func (logger *ConsoleLog) Debug(args ...interface{}) {
	logger.logging(LevelDebug, args...)
}

func (logger *ConsoleLog) Info(args ...interface{}) {
	logger.logging(LevelInfo, args...)
}

func (logger *ConsoleLog) Warn(args ...interface{}) {
	logger.logging(LevelWarn, args...)
}

func (logger *ConsoleLog) Error(args ...interface{}) {
	logger.logging(LevelError, args...)
}

func (logger *ConsoleLog) Fatal(args ...interface{}) {
	logger.logging(LevelFatal, args...)
}

// Logging
func (logger *ConsoleLog) logging(level int, args ... interface{}) {
	if level > logger.cfg.level {
		return
	}
	fmt.Println(logger.formatConsoleLogStr(logger.module, level, logger.head, args, logger.tail))
}

func (logger *ConsoleLog) append(s string, args ...interface{}) string {
	return fmt.Sprintf("%s%s", s, fmt.Sprint(args...))
}

func (logger *ConsoleLog) setHead(args ...interface{}) {
	logger.head = fmt.Sprintf("%s", fmt.Sprint(args...))
}

func (logger *ConsoleLog) setTail(args ...interface{}) {
	logger.tail = fmt.Sprintf("%s", fmt.Sprint(args...))
}

// Logging string
func (logger *ConsoleLog) formatConsoleLogStr(module string, level int, args ...interface{}) string {
	var levelStr = DefaultLevelStr
	if _, ok := consoleLevelStrMap[level]; ok {
		levelStr = consoleLevelStrMap[level]
	}
	return fmt.Sprintf("\033[33m%s\033[0m \033[36m%s\033[0m %s %s", time.Now().Format(TimeLayout), module, levelStr, fmt.Sprint(args...))
}
