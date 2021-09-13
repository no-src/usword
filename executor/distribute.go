package executor

import (
	"github.com/no-src/usword/executor/const"
	"github.com/no-src/usword/res/lang"
)

// Distribute 命令分发器接口
type Distribute interface {
	// GetExecutor 根据参数解析出相应的执行器
	GetExecutor(params ...string) Executor
}

// 内部所有提供服务的执行器
var executors map[string]Executor

type ExecDistribute struct {
}

// RegisterExecutor 注册执行器
func RegisterExecutor(name string, exec Executor) {
	if executors[name] != nil {
		panic(lang.Distribute_RegisterIExecutorRepeat + ":" + name)
	}
	executors[name] = exec
}

// GetExecutor 根据参数解析出相应的执行器
func (ed *ExecDistribute) GetExecutor(params ...string) (exec Executor) {
	cmdExec := ""
	if len(params) == 0 {
		cmdExec = _const.CmdExecuteHelp
	} else {
		cmdExec = params[0]
	}
	exec = executors[cmdExec]

	// 如果没有找到指定的执行器 则使用默认执行器help
	if exec == nil {
		exec = executors[_const.CmdExecuteHelp]
	}
	return exec
}

func init() {
	executors = make(map[string]Executor)
}
