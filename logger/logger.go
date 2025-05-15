package logger

import (
	"log"
	"log/syslog"
)

// 日志级别常量
const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

var (
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
)

// 初始化 syslog
func init() {
	debugWriter, err := syslog.New(syslog.LOG_DEBUG|syslog.LOG_DAEMON, "snmp_agentx")
	if err != nil {
		return
	}
	infoWriter, _ := syslog.New(syslog.LOG_INFO|syslog.LOG_DAEMON, "snmp_agentx")
	warnWriter, _ := syslog.New(syslog.LOG_WARNING|syslog.LOG_DAEMON, "snmp_agentx")
	errorWriter, _ := syslog.New(syslog.LOG_ERR|syslog.LOG_DAEMON, "snmp_agentx")

	debugLogger = log.New(debugWriter, "", log.Lmsgprefix)
	infoLogger = log.New(infoWriter, "", log.Lmsgprefix)
	warnLogger = log.New(warnWriter, "", log.Lmsgprefix)
	errorLogger = log.New(errorWriter, "", log.Lmsgprefix)
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
