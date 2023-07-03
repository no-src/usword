package base

import (
	"strings"

	"github.com/no-src/usword/executor/const"
)

// BaseExecutor 命令执行器
type BaseExecutor struct {
}

// BuildArgs 将接收的参数构造成字典
func (exec *BaseExecutor) BuildArgs(args ...string) map[string]string {
	dict := make(map[string]string)
	for _, item := range args {
		// kv参数，例如mode=offline
		equalIndex := strings.Index(item, "=")
		if equalIndex > 0 {
			k := item[:equalIndex]
			v := item[equalIndex+1:]
			dict[k] = v
		} else {
			// 单个值的参数，例如-i
			dict[item] = _const.ParamExist
		}
	}
	return dict
}
