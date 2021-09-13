package lang

var (
	// res client
	ConnClient_Error_ClientConnectFailed = "客户端Connect异常"
	ConnClient_Error_ClientWriteFailed   = "客户端Write异常"
	ConnClient_Error_ClientFlushFailed   = "客户端Flush异常"
	ConnClient_Error_ServerExecuteFailed = "服务端执行器执行异常"

	// res server
	ConnServer_ClientClosed    = "客户端[%s]连接已关闭，当前客户端连接数：%d"
	ConnServer_ClientConnected = "客户端[%s]连接成功，当前客户端连接数：%d"
	ConnServer_ClientRemoved   = "客户端[%s]连接移除，当前客户端连接数：%d"
)
