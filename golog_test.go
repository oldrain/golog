package golog

import (
	"testing"
)

func Test_GetLogger(t *testing.T) {
	// console
	config1 := new(Config)
	config1.SetLevel(LevelInfo)
	SetLogMode(ModeConsole)
	SetGlobalConfig(config1)
	logger1 := GetLogger("consoleGetLogger")
	logger1.Trace("trace console")
	logger1.Debug("debug console")
	logger1.Info("info console")
	logger1.Warn("warn console")
	logger1.Error("error console")
	logger1.Fatal("fatal console")

	// file, date
	config2 := new(Config)
	config2.SetLevel(LevelInfo)
	config2.SetPath("logs")
	config2.SetRotate(RotateDate)
	SetLogMode(ModeFile)
	SetGlobalConfig(config2)
	logger2 := GetLogger("fileDateGetLogger")
	logger2.Trace("trace console")
	logger2.Debug("debug console")
	logger2.Info("info console")
	logger2.Warn("warn console")
	logger2.Error("error console")
	logger2.Fatal("fatal console")

	// file, size
	config3 := new(Config)
	config3.SetLevel(LevelInfo)
	config3.SetPath("logs")
	config3.SetRotate(RotateSize)
	config3.SetRotateSize(10 * KB)
	SetLogMode(ModeFile)
	SetGlobalConfig(config3)
	logger3 := GetLogger("fileSizeGetLogger")
	logger3.Trace("trace console")
	logger3.Debug("debug console")
	logger3.Info("info console")
	logger3.Warn("warn console")
	logger3.Error("error console")
	logger3.Fatal("fatal console")
}

func Test_ConsoleLogger(t *testing.T) {
	config := new(Config)
	config.SetLevel(LevelAll)

	logger := ConsoleLogger("consoleConsoleLogger", config)
	logger.Trace("trace console")
	logger.Debug("debug console")
	logger.Info("info console")
	logger.Warn("warn console")
	logger.Error("error console")
	logger.Fatal("fatal console")
}

func Test_FileLoggerDate(t *testing.T) {
	config := new(Config)
	config.SetLevel(LevelInfo)
	config.SetPath("logs")
	config.SetRotate(RotateDate)

	logger := FileLogger("fileDateFileLoggerDate", config)
	logger.Trace("trace console")
	logger.Debug("debug console")
	logger.Info("info console")
	logger.Warn("warn console")
	logger.Error("error console")
	logger.Fatal("fatal console")
	logger2org := logger.(*FileLog)
	if logger2org.fh == nil {
		t.Error("fh nil")
	}
}

func Test_FileLoggerSize(t *testing.T) {
	config := new(Config)
	config.SetLevel(LevelInfo)
	config.SetPath("logs")
	config.SetRotate(RotateSize)
	config.SetRotateSize(10 * KB)

	logger := FileLogger("fileSizeFileLoggerSize", config)
	logger.Trace("trace console")
	logger.Debug("debug console")
	logger.Info("info console")
	logger.Warn("warn console")
	logger.Error("error console")
	logger.Fatal("fatal console")
	logger2org := logger.(*FileLog)
	if logger2org.fh == nil {
		t.Error("fh nil")
	}
}

func Test_FileLoggerTimer(t *testing.T) {
	config := new(Config)
	config.SetLevel(LevelInfo)
	config.SetPath("logs")
	config.SetRotate(RotateDate)
	config.SetTimerWrite(true)

	logger := FileLogger("fileDateFileLoggerTimer", config)
	logger.Trace("trace console")
	logger.Debug("debug console")
	logger.Info("info console")
	logger.Warn("warn console")
	logger.Error("error console")
	logger.Fatal("fatal console")
	logger2org := logger.(*FileLog)
	if logger2org.fh == nil {
		t.Error("fh nil")
	}
}

func Test_LoggerSameModule(t *testing.T) {
	var moduleName = "theOneAndOnly"
	// file
	configA := new(Config)
	configA.SetLevel(LevelInfo)
	configA.SetPath("logs")
	configA.SetRotate(RotateDate)
	configA.SetTimerWrite(true)
	SetLogMode(ModeFile)
	SetGlobalConfig(configA)
	loggerA := GetLogger(moduleName)

	// console
	configB := new(Config)
	configB.SetLevel(LevelInfo)
	loggerB := ConsoleLogger(moduleName, configB)

	if loggerA != loggerB {
		t.Error("Not the same")
	}
}
