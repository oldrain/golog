package golog

import (
	"fmt"
	"time"
)

// Console output
type ConsoleLog struct {
	cfg *Config
	module string
}

var (
	// Colorful characters
	consoleLevelStrMap = map[int]string{
		LevelFatal: "\033[35m[FATAL]\033[0m",
		LevelError: "\033[31m[ERROR]\033[0m",
		LevelWarn: "\033[33m[WARN]\033[0m",
		LevelInfo: "\033[32m[INFO]\033[0m",
		LevelDebug: "\033[36m[DEBUG]\033[0m",
		LevelTrace: "\033[34m[TRACE]\033[0m",
	}
)

func (logger *ConsoleLog) start() {
	// do nothing
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
	fmt.Println(logger.formatConsoleLogStr(logger.module, level, args...))
}

// Logging string
func (logger *ConsoleLog) formatConsoleLogStr(module string, level int, args ...interface{}) string {
	var levelStr = DefaultLevelStr
	if _, ok := consoleLevelStrMap[level]; ok {
		levelStr = consoleLevelStrMap[level]
	}
	return fmt.Sprintf("%s [%s] %s %s", time.Now().Format(TimeLayout), module, levelStr, fmt.Sprint(args...))
}
