package util

import "github.com/astaxie/beego/logs"

var ConsoleLogs *logs.BeeLogger
var FileLogs *logs.BeeLogger

func InitLogs(){
	ConsoleLogs = logs.NewLogger(1000)
	ConsoleLogs.SetLogger("console")
	FileLogs = logs.NewLogger(1000)
	FileLogs.SetLogger("file",`{"filename":”logs/error.log"}`)
}