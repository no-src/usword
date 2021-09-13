package main

import (
	"github.com/no-src/log"
	"github.com/no-src/usword/executor"
	_ "github.com/no-src/usword/executor/client"
	_ "github.com/no-src/usword/executor/help"
	_ "github.com/no-src/usword/executor/http"
	_ "github.com/no-src/usword/executor/multiclient"
	_ "github.com/no-src/usword/executor/proxy"
	_ "github.com/no-src/usword/executor/server"
	_ "github.com/no-src/usword/executor/version"
	"github.com/no-src/usword/res/lang"
	"os"
	"strings"
	"time"
)

var LogPath = "usword_log" // 日志目录

func main() {
	args := os.Args[1:]
	initLogger(args...)
	ed := executor.ExecDistribute{}
	exec := ed.GetExecutor(args...)
	result, err := exec.Handle(args...)
	if err != nil {
		log.Error(err, lang.USword_Error_IExecutor_ExecuteFailed+" Result="+string(result))
	} else {
		log.Log("%s", string(result))
	}
	// 等待日志写入完毕
	time.Sleep(time.Millisecond * 10)
}

// initLogger 初始化日志组件
func initLogger(args ...string) {
	for _, v := range args {
		if strings.Index(v, "log=") == 0 {
			logPath := v[4:]
			if len(logPath) > 0 {
				LogPath = logPath
			}
		}
	}
	log.InitDefaultLogger(log.NewMultiLogger(log.NewConsoleLogger(log.DebugLevel), log.NewFileLogger(log.DebugLevel, LogPath, "usword")))
}
