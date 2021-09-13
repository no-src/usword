package proxy

import (
	. "github.com/no-src/usword/res/lang"
)

const (
	ProxyClientProxyServer = "proxy_server"
	ProxyClientProxyPort   = "proxy_port"
	ProxyClientServer      = "server"
	ProxyClientPort        = "port"
)

var (
	InvalidProxyServer = ProxyExecutor_Error_InvalidProxyServer + ProxyClientProxyServer
	InvalidProxyPort   = ProxyExecutor_Error_InvalidProxyPort + ProxyClientProxyPort
	InvalidServer      = ProxyExecutor_Error_InvalidServer + ProxyClientServer
	InvalidPort        = ProxyExecutor_Error_InvalidPort + ProxyClientPort
)
