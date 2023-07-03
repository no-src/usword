package version

import (
	"fmt"

	"github.com/no-src/usword/executor"
	base_exec "github.com/no-src/usword/executor/base"
	"github.com/no-src/usword/executor/const"
	"github.com/no-src/usword/executor/help"
	"github.com/no-src/usword/res/lang"
)

// VersionExecutor Version命令执行器
type VersionExecutor struct {
	base_exec.BaseExecutor
}

// Handle 输出当前程序的编译和版本信息
func (exec *VersionExecutor) Handle(params ...string) (result []byte, err error) {
	versionInfo := fmt.Sprintf("%s\t%s\n", lang.VersionExecutor_USwordVersion, USwordVersion)
	result = []byte(versionInfo)
	return result, err
}

func init() {
	register()
	help.RegisterHelpInfo(_const.CmdExecuteVersion, lang.VersionExecutor_Desc, man)
}

func man() string {
	msg := fmt.Sprintln(lang.VersionExecutor_Welcome)
	msg += fmt.Sprintln(lang.VersionExecutor_Usage)
	return msg
}

// 注册当前IExecutor
func register() {
	executor.RegisterExecutor(_const.CmdExecuteVersion, &VersionExecutor{})
}
