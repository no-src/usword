package connector

import (
	"errors"

	"github.com/no-src/usword/res/lang"
)

const (
	DefaultServerHost = "127.0.0.1"
	DefaultServerPort = 9091
)

var (
	// EndIdentity 特殊字符串，标识一条通讯消息的结尾，一般其末尾会还有一个CRLFBytes
	EndIdentity = []byte("_$#END#$_")
	// ErrorIdentity 特殊字符串，标识响应结果为错误消息，出现在EndIdentity的前面
	ErrorIdentity = []byte("_$#ERROR#$_")
	// ErrorEndIdentity 特殊字符串,包含错误的结尾标识
	ErrorEndIdentity = append(ErrorIdentity, EndIdentity...)
	// CRLFBytes 换行字符的byte数组
	CRLFBytes = []byte(CRLF)
)

const (
	// CRLF 换行字符串
	CRLF = "\r\n"
)

// ServerExecuteError 服务器端执行器执行异常
var ServerExecuteError = errors.New(lang.ConnClient_Error_ServerExecuteFailed)
