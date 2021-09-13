package res

import (
	. "github.com/no-src/usword/res/lang"
)

func init() {
	// rewrite HttpExecutor
	HttpExecutor_Error_MustSetUrl = "must set up param :url"
	HttpExecutor_Desc = "execute http command"
	HttpExecutor_Welcome = "Welcome To Use USword HTTP !"
	HttpExecutor_Usage = "Usage: usword http [arg=value]..."
	HttpExecutor_ArgDesc = "arg: current supported args"
	HttpExecutor_Arg_CmdHttpRequestUrl = "set up http url,default is empty"
	HttpExecutor_Arg_CmdHttpMethod = "set up http method,default is GET"
	HttpExecutor_Arg_CmdHttpRequestContentType = "set up http request header:Content-Type"
	HttpExecutor_Arg_CmdHttpRequestBody = "set up http body,default is empty"
	HttpExecutor_Arg_CmdHttpProtocol = "set up http protocol,default is HTTP/1.1"
	HttpExecutor_Arg_CmdHttpHeader = "set up http header,default is empty"
	HttpExecutor_Arg_CmdHttpCookies = "set up http cookies,default is empty"
	HttpExecutor_Arg_CmdHttpResponseOutput = "http response output type，default is console"

	// rewrite VersionExecutor
	VersionExecutor_USwordVersion = "USword Version"
	VersionExecutor_USwordBuild = "Build From"
	VersionExecutor_USwordCommit = "Commit"
	VersionExecutor_Desc = "check current version info"
	VersionExecutor_Welcome = "Welcome To Use USword Version !"
	VersionExecutor_Usage = "Usage: usword version"

	// rewrite res ClientExecutor
	ClientExecutor_Error_ConnAbortErr = "remote server connect failed,operator abort"
	ClientExecutor_ServerConnect_Success = "remote server connect succeed"
	ClientExecutor_ServerConnect_Failed = "remote server connect failed"
	ClientExecutor_ServerConnectRetry_Success = "the[%d]times try to connect server succeed"
	ClientExecutor_ServerConnectRetry_Failed = "the[%d]times try to connect server failed"
	ClientExecutor_ServerWrite_Failed = "write to server failed,operator abort"
	ClientExecutor_ServerRead_Failed = "read from server failed,operator abort"
	ClientExecutor_Desc = "run a USword client"
	ClientExecutor_Welcome = "Welcome to Use USword Client !"
	ClientExecutor_Usage = "Usage: usword client [arg=value]..."
	ClientExecutor_ArgDesc = "arg: current supported args"
	ClientExecutor_Arg_ClientArgServer = "the server host client connected ,default is 127.0.0.1"
	ClientExecutor_Arg_ClientArgPort = "the server port client listened ,default is 9091"
	ClientExecutor_Arg_ClientRunMode = "client run mode,optional:[%s] [%s],default is [%s]"
	ClientExecutor_Arg_ClientInteractiveExec = "interactive to execute command,in default ,program wil exit when execute finished once"
	ClientExecutor_Arg_ClientExec = "the command line will executed,if in online mode,the command will send to server"

	// rewrite ClientInteractive
	ClientInteractive_InputTips = "please input your command"
	ClientInteractive_Exit = "exit client"
	ClientInteractive_Clear = "clear console"
	ClientInteractive_Error_Input = "get user input error"
	ClientInteractive_Error_ExecError = "command execute failed"
	ClientInteractive_Cmd = "command"
	ClientInteractive_Finish = "command execute finished"

	// rewrite HelpExecutor
	HelpExecutor_Error_RegiterHelpRepeat = "register same HelpExecutor repeat"
	HelpExecutor_USword_Welcome = "Welcome to Use USword !"
	HelpExecutor_USword_Usage = "Usage: usword [cmd] [args...]"
	HelpExecutor_USword_HelpUsage = "you can use this command get detail info:usword help [cmd] "
	HelpExecutor_USword_HelpCmdSupport = "cmd: current supported cmd  as follow"
	HelpExecutor_Desc = "execute help command"
	HelpExecutor_Welcome = "Welcome to Use USword Help !"
	HelpExecutor_Usage = "Usage: usword help [cmd]"
	HelpExecutor_HelpCmdSupport = "cmd: current supported cmd  as follow"

	// rewrite MultiClientExecutor
	MultiClientExecutor_Error_ClientArgServers_NotFound = "not found servers param,please check param : "
	MultiClientExecutor_Error_ClientArgServers_NotValid = "servers not valid,please check param :"
	MultiClientExecutor_Error_ConnectServerFailed = "connect server failed"
	MultiClientExecutor_ConnectServerSuccess = "connect server succeed"
	MultiClientExecutor_Desc = "run a usword client,connect to multi server"
	MultiClientExecutor_Welcome = "Welcome to Use USword MultiClient !"
	MultiClientExecutor_Usage = "Usage: usword multiclient [arg=value]..."
	MultiClientExecutor_ArgDesc = "arg: current supported args"
	MultiClientExecutor_Arg_ClientArgServers = "the servers client will connected,format like 127.0.0.1:8185?mode=online,when you have multi server,just use english comma(,) to split"
	MultiClientExecutor_Arg_ClientRunMode = "client run mode,optionals:[%s] [%s],default is [%s]"
	MultiClientExecutor_Arg_ClientInteractiveExec = "execute command interactive，on default,after program executed command once,and just exit "
	MultiClientExecutor_Arg_ClientExec = "the command will send to client executed，if on online mode，the command will send to server,and executed"

	// rewrite ProxyClientExecutor and ProxyServerExecutor
	ProxyExecutor_Error_InvalidProxyServer = "invalid proxy server :"
	ProxyExecutor_Error_InvalidProxyPort = "invalid proxy server port :"
	ProxyExecutor_Error_InvalidServer = "invalid server :"
	ProxyExecutor_Error_InvalidPort = "invalid server port :"
	ProxyExecutor_Error_ProxyConfigIsEmpty = "proxy config is empty"
	ProxyExecutor_Error_NotFoundLegalProxyConfig = "can't found legal proxy config"
	ProxyExecutor_Error_NotFoundRealServer = "can't found legal real server config"
	ProxyExecutor_Error_ExecuteClientCommandFailed = "execute client [%s] command [%s]failed"
	ProxyExecutor_Error_ServerReponseFailed = "response client[%s] command [%s]failed"
	ProxyExecutor_Info_ServerReponseSuccess = "response client[%s] command [%s]succeed"

	ProxyClientExecutor_Desc = "run a usword proxy client"
	ProxyClientExecutor_Welcome = "Welcome to Use USword ProxyClient !"
	ProxyClientExecutor_Usage = "Usage: usword proxy_client [arg=value]..."
	ProxyClientExecutor_ArgDesc = "arg: current supported args"

	ProxyExecutor_Arg_ProxyClientProxyServer = "proxy server"
	ProxyExecutor_Arg_ProxyClientProxyPort = "proxy server port"
	ProxyExecutor_Arg_ProxyClientServer = "real server"
	ProxyExecutor_Arg_ProxyClientPort = "real server port"
	MultiClientExecutor_Arg_ClientInteractiveExec = ClientExecutor_Arg_ClientInteractiveExec
	MultiClientExecutor_Arg_ClientExec = ClientExecutor_Arg_ClientExec

	ProxyServerExecutor_ReceiveClientCommand = "received client[%s]command:[%s]"
	ProxyServerExecutor_Desc = "run a usword proxy server"
	ProxyServerExecutor_Welcome = "Welcome to Use USword ProxyServer !"
	ProxyServerExecutor_Usage = "Usage: usword proxy_server [arg=value]..."
	ProxyServerExecutor_ArgDesc = "arg: current supported args"
	ProxyServerExecutor_Arg_ProxyClientProxyServer = "proxy server bind address"
	ProxyServerExecutor_Arg_ProxyClientProxyPort = "proxy server listen port,default is 9091"

	// rewrite ServerExecutor
	ServerExecutor_Error_ListenFailed = "server[%s:%d]start listen failed"
	ServerExecutor_ListenSucceed = "server[%s:%d]start listen succeed！"
	ServerExecutor_Error_AcceptFailed = "server accept failed ！"
	ServerExecutor_ReceiveClientCommand = ProxyServerExecutor_ReceiveClientCommand
	ServerExecutor_Error_ExecuteClientCommandFailed = ProxyExecutor_Error_ExecuteClientCommandFailed
	ServerExecutor_Error_ServerReponseFailed = ProxyExecutor_Error_ServerReponseFailed
	ServerExecutor_Info_ServerReponseSuccess = ProxyExecutor_Info_ServerReponseSuccess
	ServerExecutor_Desc = "run a usword server"
	ServerExecutor_Welcome = "Welcome to Use USword Server !"
	ServerExecutor_Usage = "Usage: usword server [arg=value]..."
	ServerExecutor_ArgDesc = "arg: current supported args"
	ServerExecutor_Arg_ServerArgServerHost = "server bind address"
	ServerExecutor_Arg_ServerArgPort = "server listen port,default is 9091"

	// rewrite distribute
	Distribute_RegisterIExecutorRepeat = "IExecutor Register Repeat"
}
