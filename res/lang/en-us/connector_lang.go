package res

import (
	. "github.com/no-src/usword/res/lang"
)

func init() {
	// rewrite client
	ConnClient_Error_ClientConnectFailed = "client connect failed"
	ConnClient_Error_ClientWriteFailed = "client write failed"
	ConnClient_Error_ClientFlushFailed = "client flush failed"
	ConnClient_Error_ServerExecuteFailed = "server executor execute failed"

	// rewrite server
	ConnServer_ClientClosed = "client[%s]conn closed,current client connect count:%d"
	ConnServer_ClientConnected = "client[%s]conn succeed，current client connect count:%d"
	ConnServer_ClientRemoved = "client[%s]conn removed，current client connect count:%d"
}
