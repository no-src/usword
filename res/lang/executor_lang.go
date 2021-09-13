package lang

// res HttpExecutor
var (
	HttpExecutor_Error_MustSetUrl              = "请设置要访问的地址:url"
	HttpExecutor_Desc                          = "解析执行http命令"
	HttpExecutor_Welcome                       = "欢迎使用USword HTTP组件!"
	HttpExecutor_Usage                         = "使用方式: usword http [arg=value]..."
	HttpExecutor_ArgDesc                       = "arg: 当前支持的参数如下"
	HttpExecutor_Arg_CmdHttpRequestUrl         = "指定http的请求url，默认为空"
	HttpExecutor_Arg_CmdHttpMethod             = "指定http的method，默认为GET"
	HttpExecutor_Arg_CmdHttpRequestContentType = "指定http请求的Content-Type请求头"
	HttpExecutor_Arg_CmdHttpRequestBody        = "指定http的请求body，默认为空"
	HttpExecutor_Arg_CmdHttpProtocol           = "指定http的protocol，默认为HTTP/1.1"
	HttpExecutor_Arg_CmdHttpHeader             = "指定http的header，默认为空"
	HttpExecutor_Arg_CmdHttpCookies            = "指定http的cookies，默认为空"
	HttpExecutor_Arg_CmdHttpResponseOutput     = "http响应结果的输出路径，默认在控制台输出"
)

// res VersionExecutor
var (
	VersionExecutor_USwordVersion = "USword版本"
	VersionExecutor_USwordBuild   = "构建于"
	VersionExecutor_USwordCommit  = "Commit"
	VersionExecutor_Desc          = "查看当前程序版本信息"
	VersionExecutor_Welcome       = "欢迎使用USword Version组件!"
	VersionExecutor_Usage         = "使用方式: usword version"
)

// res ClientExecutor
var (
	ClientExecutor_Error_ConnAbortErr         = "远程服务器连接异常，操作中断"
	ClientExecutor_ServerConnect_Success      = "连接服务器成功"
	ClientExecutor_ServerConnect_Failed       = "连接服务器失败"
	ClientExecutor_ServerConnectRetry_Success = "第[%d]次尝试重新连接服务器成功"
	ClientExecutor_ServerConnectRetry_Failed  = "第[%d]次尝试重新连接服务器失败"
	ClientExecutor_ServerWrite_Failed         = "向远程服务器发送数据异常，操作中断"
	ClientExecutor_ServerRead_Failed          = "读取远程服务器响应结果异常，操作中断"
	ClientExecutor_Desc                       = "运行一个usword客户端"
	ClientExecutor_Welcome                    = "欢迎使用USword Client组件!"
	ClientExecutor_Usage                      = "使用方式: usword client [arg=value]..."
	ClientExecutor_ArgDesc                    = "arg: 当前支持的参数如下"
	ClientExecutor_Arg_ClientArgServer        = "客户端要连接的服务器主机地址，默认是127.0.0.1"
	ClientExecutor_Arg_ClientArgPort          = "客户端要连接的服务器主机端口，默认是9091"
	ClientExecutor_Arg_ClientRunMode          = "客户端运行模式,可选值为[%s] [%s],默认值为[%s]"
	ClientExecutor_Arg_ClientInteractiveExec  = "在交互式界面执行命令，程序默认在执行完一次命令后退出"
	ClientExecutor_Arg_ClientExec             = "提交给客户端执行的命令，如果是online模式，会转发给server执行"
)

// res ClientInteractive
var (
	ClientInteractive_InputTips       = "请输入要执行的命令"
	ClientInteractive_Exit            = "退出客户端"
	ClientInteractive_Clear           = "清空控制台"
	ClientInteractive_Error_Input     = "获取用户输入命令异常"
	ClientInteractive_Error_ExecError = "命令执行异常"
	ClientInteractive_Cmd             = "命令"
	ClientInteractive_Finish          = "命令执行完成"
)

// res HelpExecutor
var (
	HelpExecutor_Error_RegiterHelpRepeat = "重复注册帮助信息"
	HelpExecutor_USword_Welcome          = "欢迎使用USword !"
	HelpExecutor_USword_Usage            = "使用方式: usword [cmd] [args...]"
	HelpExecutor_USword_HelpUsage        = "您可以使用usword help [cmd] 命令查询具体cmd的使用方法"
	HelpExecutor_USword_HelpCmdSupport   = "cmd: 当前支持的命令如下"
	HelpExecutor_Desc                    = "解析执行help相关命令"
	HelpExecutor_Welcome                 = "欢迎使用USword Help组件!"
	HelpExecutor_Usage                   = "使用方式: usword help [cmd]"
	HelpExecutor_HelpCmdSupport          = "cmd: 当前支持的参数如下"
)

// res MultiClientExecutor
var (
	MultiClientExecutor_Error_ClientArgServers_NotFound = "未指定要连接的服务器地址，请检查参数设置："
	MultiClientExecutor_Error_ClientArgServers_NotValid = "未解析到正确的服务器列表，请检查参数设置："
	MultiClientExecutor_Error_ConnectServerFailed       = "连接远程服务器异常"
	MultiClientExecutor_ConnectServerSuccess            = "连接服务器成功"
	MultiClientExecutor_Desc                            = "运行一个usword客户端，连接多个服务器"
	MultiClientExecutor_Welcome                         = "欢迎使用USword MultiClient组件 !"
	MultiClientExecutor_Usage                           = "使用方式: usword multiclient [arg=value]..."
	MultiClientExecutor_ArgDesc                         = "arg: 当前支持的参数如下"
	MultiClientExecutor_Arg_ClientArgServers            = "客户端要连接的1个或多个服务器主机地址及参数设置，格式如下127.0.0.1:8185?mode=online，多个地址间用英文逗号分隔"
	MultiClientExecutor_Arg_ClientRunMode               = "客户端运行模式,可选值为[%s] [%s],默认值为[%s]"
	MultiClientExecutor_Arg_ClientInteractiveExec       = ClientExecutor_Arg_ClientInteractiveExec
	MultiClientExecutor_Arg_ClientExec                  = ClientExecutor_Arg_ClientExec
)

// res ProxyClientExecutor and ProxyServerExecutor
var (
	ProxyExecutor_Error_InvalidProxyServer         = "无效的代理服务器地址:"
	ProxyExecutor_Error_InvalidProxyPort           = "无效的代理服务器端口:"
	ProxyExecutor_Error_InvalidServer              = "无效的服务器地址:"
	ProxyExecutor_Error_InvalidPort                = "无效的服务器端口:"
	ProxyExecutor_Error_ProxyConfigIsEmpty         = "代理配置信息为空"
	ProxyExecutor_Error_NotFoundLegalProxyConfig   = "为获取到合法的代理配置"
	ProxyExecutor_Error_NotFoundRealServer         = "未获取到真实服务器信息"
	ProxyExecutor_Error_ExecuteClientCommandFailed = "执行客户端[%s]的命令[%s]发生异常"
	ProxyExecutor_Error_ServerReponseFailed        = "响应客户端[%s]命令[%s]执行结果发生异常"
	ProxyExecutor_Info_ServerReponseSuccess        = "响应客户端[%s]命令[%s]执行结果成功"

	ProxyClientExecutor_Desc    = "运行一个usword代理客户端"
	ProxyClientExecutor_Welcome = "欢迎使用USword ProxyClient组件 !"
	ProxyClientExecutor_Usage   = "使用方式: usword proxy_client [arg=value]..."
	ProxyClientExecutor_ArgDesc = "arg: 当前支持的参数如下"

	ProxyExecutor_Arg_ProxyClientProxyServer = "代理服务器地址"
	ProxyExecutor_Arg_ProxyClientProxyPort   = "代理服务器端口"
	ProxyExecutor_Arg_ProxyClientServer      = "真实服务器地址"
	ProxyExecutor_Arg_ProxyClientPort        = "真实服务器端口"
	ProxyExecutor_Arg_ClientInteractiveExec  = "在交互式界面执行命令，程序默认在执行完一次命令后退出"
	ProxyExecutor_Arg_ClientExec             = "提交给服务器执行的命令"

	ProxyServerExecutor_ReceiveClientCommand       = "接收到来自客户端[%s]的命令:[%s]"
	ProxyServerExecutor_Desc                       = "运行一个usword代理服务器"
	ProxyServerExecutor_Welcome                    = "欢迎使用USword ProxyServer组件 !"
	ProxyServerExecutor_Usage                      = "使用方式: usword proxy_server [arg=value]..."
	ProxyServerExecutor_ArgDesc                    = "arg: 当前支持的参数如下"
	ProxyServerExecutor_Arg_ProxyClientProxyServer = "代理服务器主机绑定地址"
	ProxyServerExecutor_Arg_ProxyClientProxyPort   = "代理服务器主机监听端口，默认是9091"
)

// res ServerExecutor
var (
	ServerExecutor_Error_ListenFailed               = "服务器[%s:%d]启动监听异常"
	ServerExecutor_ListenSucceed                    = "服务器[%s:%d]启动成功！"
	ServerExecutor_Error_AcceptFailed               = "服务器Accept异常！"
	ServerExecutor_ReceiveClientCommand             = ProxyServerExecutor_ReceiveClientCommand
	ServerExecutor_Error_ExecuteClientCommandFailed = ProxyExecutor_Error_ExecuteClientCommandFailed
	ServerExecutor_Error_ServerReponseFailed        = ProxyExecutor_Error_ServerReponseFailed
	ServerExecutor_Info_ServerReponseSuccess        = ProxyExecutor_Info_ServerReponseSuccess
	ServerExecutor_Desc                             = "运行一个usword服务端"
	ServerExecutor_Welcome                          = "欢迎使用USword Server组件 !"
	ServerExecutor_Usage                            = "使用方式: usword server [arg=value]..."
	ServerExecutor_ArgDesc                          = "arg: 当前支持的参数如下"
	ServerExecutor_Arg_ServerArgServerHost          = "服务器主机绑定地址"
	ServerExecutor_Arg_ServerArgPort                = "服务器主机监听端口，默认是9091"
)

// res distribute
var (
	Distribute_RegisterIExecutorRepeat = "不能注册重复的IExecutor执行器"
)
