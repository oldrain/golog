// Copyright 2019 golog Author. All Rights Reserved.
// License that can be found in the LICENSE file.

package golog

import (
	"fmt"
	"time"
)

// Log level
const (
	LevelOff = iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
	LevelAll
)

// Size unit
const (
	_        = iota
	KB int64 = 1 << (iota * 10)
	MB
	GB
	TB
)

const (
	DefaultPath = "logs" //Default log file path
	FileExt     = ".log" //File suffix

	ErrLabel = "[LOGGER]" //Standard log label

	DefaultLevelStr = "[UNKNOWN]" //Unknown level

	ModeConsole = 1 //Console log
	ModeFile    = 2 //File log

	RotateDate = 1 //Rotate as date
	RotateSize = 2 //Rotate as size

	FileOperateMode = 0755 //Write Permissions
	FileCreateMode  = 0755 //Create Permissions

	CheckInterval = 900 * time.Millisecond //Checking duration
	WriteInterval = 1200 * time.Millisecond //Batch write duration

	FlushCount = 100 //Trigger flush buffer

	DateLayout = "2006-01-02"
	TimeLayout = "2006-01-02 15:04:05"

	EmptyString = ""
)

type Logger interface {
	AppendHead(args ...interface{})
	AppendTail(args ...interface{})
	EraseHead()
	EraseTail()
	Erase()

	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})

	start()
	logging(level int, args ...interface{})
}

var (
	loggerContainer = make(map[string]Logger)

	logMode = ModeFile

	// Global config
	logConfig *Config

	levelStrMap = map[int]string{
		LevelOff: "[OFF]",
		LevelFatal: "[FATAL]",
		LevelError: "[ERROR]",
		LevelWarn: "[WARN]",
		LevelInfo: "[INFO]",
		LevelDebug: "[DEBUG]",
		LevelTrace: "[TRACE]",
		LevelAll: "[ALL]",
	}
)

type Config struct {
	level      int

	path string
	rotate int
	rotateSize int64

	timerWrite bool
}

func (config *Config) SetLevel(level int) {
	config.level = level
}

func (config *Config) SetPath(path string) {
	config.path = path
}

func (config *Config) SetRotate(rotate int) {
	config.rotate = rotate
}

func (config *Config) SetRotateSize(size int64) {
	config.rotateSize = size
}

func (config *Config) SetTimerWrite(timerWrite bool) {
	config.timerWrite = timerWrite
}


func SetLogMode(mode int) {
	logMode = mode
}

func SetGlobalConfig(config *Config) {
	logConfig = config
}

func GetLogger(moduleName string) Logger {
	if hasMode(logMode, ModeConsole) {
		return GetConsoleLogger(moduleName)
	}

	if hasMode(logMode, ModeFile) {
		return GetFileLogger(moduleName)
	}

	// Default
	return GetFileLogger(moduleName)
}

func GetConsoleLogger(moduleName string) Logger {
	var logger Logger

	logger = getFromContainer(moduleName)
	if logger == nil {
		config := getConfig()
		logger = newConsoleLogger(moduleName, config)
	}

	set2Container(moduleName, logger)

	return logger
}

func ConsoleLogger(moduleName string, config *Config) Logger {
	var logger Logger

	logger = getFromContainer(moduleName)
	if logger == nil {
		logger = newConsoleLogger(moduleName, config)
	}

	set2Container(moduleName, logger)

	return logger
}

func GetFileLogger(moduleName string) Logger {
	var logger Logger

	logger = getFromContainer(moduleName)
	if logger == nil {
		logger = newFileLogger(moduleName, logConfig)
	}

	set2Container(moduleName, logger)

	return logger
}

func FileLogger(moduleName string, config *Config) Logger {
	var logger Logger

	logger = getFromContainer(moduleName)
	if logger == nil {
		logger = newFileLogger(moduleName, config)
	}

	set2Container(moduleName, logger)

	return logger
}

func newConsoleLogger(moduleName string, config *Config) Logger {
	logger := Logger(&ConsoleLog{
		cfg:    config,
		module: moduleName,
	})

	logger.start()

	return logger
}

func newFileLogger(moduleName string, config *Config) Logger {
	logger := Logger(&FileLog{
		cfg: config,
		module: moduleName,
	})

	logger.start()

	return logger
}

// Global or default
func getConfig() *Config {
	if logConfig != nil {
		return logConfig
	} else {
		return &Config{
			level: LevelInfo,
			path: DefaultPath,
			rotate: RotateDate,
			rotateSize: 10 * MB,
			timerWrite: false,
		}
	}
}

func getFromContainer(moduleName string) Logger {
	if _, ok := loggerContainer[moduleName]; ok {
		return loggerContainer[moduleName]
	} else {
		return nil
	}
}

func set2Container(moduleName string, logger Logger) {
	if logger != nil {
		loggerContainer[moduleName] = logger
	}
}

func hasMode(logMode, mode int) bool {
	return (logMode & mode) != 0
}

func formatLogStr(module string, level int, args ...interface{}) string {
	var levelStr = DefaultLevelStr
	if _, ok := levelStrMap[level]; ok {
		levelStr = levelStrMap[level]
	}
	return fmt.Sprintf("%s [%s] %s %s\n", time.Now().Format(TimeLayout), module, levelStr, fmt.Sprint(args...))
}
