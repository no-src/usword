package executor

// Executor 命令解析执行接口
type Executor interface {
	// Handle 根据参数执行相应的命令
	Handle(params ...string) (result []byte, err error)
}
