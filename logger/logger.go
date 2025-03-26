package logger

import (
	"log"
	"log/syslog"
)

var (
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
)

// 初始化 syslog
func init() {
	sysLogger, err := syslog.New(syslog.LOG_INFO|syslog.LOG_DAEMON, "snmp_agentx")
	if err != nil {
		log.Fatalf("Failed to initialize syslog: %v", err)
	}

	debugLogger = log.New(sysLogger, "DEBUG: ", log.LstdFlags)
	infoLogger = log.New(sysLogger, "INFO: ", log.LstdFlags)
	warnLogger = log.New(sysLogger, "WARN: ", log.LstdFlags)
	errorLogger = log.New(sysLogger, "ERROR: ", log.LstdFlags)
}

// Debug 级别日志
func Debug(v ...interface{}) {
	debugLogger.Println(v...)
}

// Info 级别日志
func Info(v ...interface{}) {
	infoLogger.Println(v...)
}

// Warn 级别日志
func Warn(v ...interface{}) {
	warnLogger.Println(v...)
}

// Error 级别日志
func Error(v ...interface{}) {
	errorLogger.Println(v...)
}
