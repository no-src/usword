package _const

const (
	CmdLogPath = "log"      // 日志路径
	CmdLogMode = "log_mode" //日志模式，默认在控制台输出
)

const (
	CmdExecuteHelp        = "help"         // 解析执行help相关命令
	CmdExecuteHttp        = "http"         // 解析执行http命令
	CmdExecuteClient      = "client"       // 运行一个usword客户端
	CmdExecuteMultiClient = "multiclient"  // 运行一个usword客户端，连接多个服务器
	CmdExecuteServer      = "server"       // 运行一个usword服务端
	CmdExecuteVersion     = "version"      // 查看当前程序版本信息
	CmdExecuteProxyClient = "proxy_client" // 运行一个usword代理客户端
	CmdExecuteProxyServer = "proxy_server" // 运行一个usword代理服务器
)

const (
	ParamExist = "$exist$" // 参数存在标志
)
