package golog

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

// File output
type FileLog struct {
	cfg *Config

	module string

	fileCount int //File count

	mutexFile *sync.Mutex //File mutex
	fh        *os.File    //Logging file handler

	checkTimer *time.Timer //Daemon timer
	writeTimer *time.Timer //Batch write timer

	writeCount int           //Number of log records waiting for write
	mutexBuff  *sync.Mutex   //Buffer mutex
	buff       *bytes.Buffer //Write buffer
}

func (logger *FileLog) start() {
	if logger.fh != nil {
		return
	}

	logger.mutexFile = new(sync.Mutex)
	logger.mutexBuff = new(sync.Mutex)
	logger.buff = new(bytes.Buffer)

	// Checking
	logger.checkFile()

	// Timer has been turned on
	if logger.cfg.timerWrite {
		go func() {
			logger.checkTimer = time.NewTimer(CheckInterval)
			logger.writeTimer = time.NewTimer(WriteInterval)
			defer logger.checkTimer.Stop()
			defer logger.writeTimer.Stop()
			for {
				select {
				case <-logger.checkTimer.C:
					logger.checkFile()
					log.Println("********** checkTimer ***********")
					break
				case <-logger.writeTimer.C:
					logger.flushBuff()
					log.Println("********** writeTimer ***********")
				}
				logger.checkTimer.Reset(CheckInterval)
				logger.writeTimer.Reset(CheckInterval)
			}
		}()
	}
}

func (logger *FileLog) Trace(args ...interface{}) {
	logger.logging(LevelTrace, args...)
}

func (logger *FileLog) Debug(args ...interface{}) {
	logger.logging(LevelDebug, args...)
}

func (logger *FileLog) Info(args ...interface{}) {
	logger.logging(LevelInfo, args...)
}

func (logger *FileLog) Warn(args ...interface{}) {
	logger.logging(LevelWarn, args...)
}

func (logger *FileLog) Error(args ...interface{}) {
	logger.logging(LevelError, args...)
}

func (logger *FileLog) Fatal(args ...interface{}) {
	logger.logging(LevelFatal, args...)
}

// Logging
func (logger *FileLog) logging(level int, args ...interface{}) {
	if level > logger.cfg.level {
		return
	}

	defer func() {
		// Try to flush buffer
		if logger.needFlushBuff() {
			logger.flushBuff()
		}
		// Checking if needs
		if logger.needCheckFile() {
			logger.checkFile()
		}
	}()

	logStr := formatLogStr(logger.module, level, args...)

	// Write to buffer
	logger.mutexBuff.Lock()
	defer logger.mutexBuff.Unlock()
	logger.buff.WriteString(logStr)
	logger.writeCount++
}

// Write to file
func (logger *FileLog) flushBuff() {
	defer logger.buff.Reset()

	logger.mutexFile.Lock()
	defer logger.mutexFile.Unlock()
	_, err := logger.fh.WriteString(logger.buff.String())
	if nil != err {
		log.Println(ErrLabel, logger.module, "WriteString error: ", err)
	}
	logger.writeCount = 0
}

// Check file before logging
func (logger *FileLog) checkFile() {
	logger.mutexFile.Lock()
	defer logger.mutexFile.Unlock()

	// Rotate as file size
	if (logger.cfg.rotate == RotateSize) && (logger.fh != nil) {
		fileInfo, err := logger.fh.Stat()
		if err != nil {
			log.Println(ErrLabel, logger.module, "Stat error: ", err)
		}
		// Increase file count
		if fileInfo.Size() >= logger.cfg.rotateSize {
			logger.fileCount++
		}
	}

	filePath := logger.newFilePath()

	fileDir := path.Dir(filePath)
	if !logger.isFileExist(fileDir) {
		err := os.MkdirAll(fileDir, FileCreateMode)
		if nil != err {
			log.Println(ErrLabel, logger.module, "MkdirAll error: ", err, fileDir)
		}
	}

	if !logger.isFileExist(filePath) {
		if nil != logger.fh {
			err := logger.fh.Close() //Previous file handler
			if nil != err {
				log.Println(ErrLabel, logger.module, "Close error: ", err, fileDir)
			}
			logger.fileCount = 0 //Reset fileCount
		}
	}

	// Open or create file
	var err error
	logger.fh, err = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, FileOperateMode)
	if nil != err {
		log.Println(ErrLabel, logger.module, "OpenFile file error: ", err, filePath)
	}
}

// Flush buffer: timer turned off || to many log record
func (logger *FileLog) needFlushBuff() bool {
	if !logger.cfg.timerWrite || (logger.cfg.timerWrite && (logger.writeCount >= FlushCount)) {
		return true
	}
	return false
}

// Check logging file: timer turned off
func (logger *FileLog) needCheckFile() bool {
	if !logger.cfg.timerWrite {
		return true
	}
	return false
}

// Is file exists
func (logger *FileLog) isFileExist(path string) bool {
	_, err := os.Stat(path)
	return nil == err || os.IsExist(err)
}

// New logging file path
func (logger *FileLog) newFilePath() string {
	if logger.cfg.rotate == RotateSize {
		return fmt.Sprintf("%s/%s.%d%s", logger.cfg.path, logger.module, logger.fileCount, FileExt)
	} else {
		return fmt.Sprintf("%s/%s/%s%s", logger.cfg.path, time.Now().Format(DateLayout), logger.module, FileExt)
	}
}
