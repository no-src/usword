package help

import (
	"fmt"
	"github.com/no-src/usword/executor"
	base_exec "github.com/no-src/usword/executor/base"
	"github.com/no-src/usword/executor/const"
	"github.com/no-src/usword/res/lang"
)

// HelpExecutor Help命令执行器
type HelpExecutor struct {
	base_exec.BaseExecutor
}

var helpInfo map[string]func() string
var helpInfoDesc map[string]string

// RegisterHelpInfo 注册帮助信息
// cmd 命令名称
// desc 命令简述信息
// helpFunc 帮助信息输出函数
func RegisterHelpInfo(cmd string, desc string, helpFunc func() string) {
	if helpInfo[cmd] != nil {
		panic(lang.HelpExecutor_Error_RegiterHelpRepeat + " cmd=" + cmd)
	}
	helpInfo[cmd] = helpFunc
	helpInfoDesc[cmd] = desc
}

// Handle 根据参数执行相应的Help命令
func (exec *HelpExecutor) Handle(params ...string) (result []byte, err error) {
	cmd := ""
	//	示例：usword help http
	if len(params) > 1 {
		cmd = params[1]
	}
	helpFunc := helpInfo[cmd]
	if helpFunc != nil {
		result = []byte(helpFunc())
		return result, nil
	}
	argsDict := argsComment()
	msg := fmt.Sprintln(lang.HelpExecutor_USword_Welcome)
	msg += fmt.Sprintln(lang.HelpExecutor_USword_Usage)
	msg += fmt.Sprintln(lang.HelpExecutor_USword_HelpUsage)
	msg += fmt.Sprintln(lang.HelpExecutor_USword_HelpCmdSupport)
	for k, v := range argsDict {
		msg += fmt.Sprintf("\t%s\t%s\n", k, v)
	}
	result = []byte(msg)
	return result, err
}

func init() {
	register()
	helpInfo = make(map[string]func() string)
	helpInfoDesc = make(map[string]string)
	RegisterHelpInfo(_const.CmdExecuteHelp, lang.HelpExecutor_Desc, man)
}

func man() string {
	argsDict := argsComment()
	msg := fmt.Sprintln(lang.HelpExecutor_Welcome)
	msg += fmt.Sprintln(lang.HelpExecutor_Usage)
	msg += fmt.Sprintln(lang.HelpExecutor_HelpCmdSupport)
	for k, v := range argsDict {
		msg += fmt.Sprintf("\t%s\t%s\n", k, v)
	}
	return msg
}

func argsComment() map[string]string {
	return helpInfoDesc
}

// 注册当前IExecutor
func register() {
	executor.RegisterExecutor(_const.CmdExecuteHelp, &HelpExecutor{})
}
