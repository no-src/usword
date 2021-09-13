package res

import (
	. "github.com/no-src/usword/res/lang"
)

func init() {
	// rewrite http
	HTTP_Error_MethodNotSupported = "http method not supported"
	HTTP_Error_RequestFailed = "request failed"
	HTTP_Error_ReadBodyFailed = "read response body failed"
}
