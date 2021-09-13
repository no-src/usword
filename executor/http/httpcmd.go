package http

const (
	CmdHttpRequestUrl         = "url"          // 指定http的请求url，默认为空
	CmdHttpMethod             = "method"       // 指定http的method，默认为GET
	CmdHttpRequestContentType = "content_type" // 指定http请求的Content-Type请求头
	CmdHttpRequestBody        = "body"         // 指定http的请求body，默认为空
	CmdHttpProtocol           = "protocol"     // 指定http的protocol，默认为HTTP/1.1
	CmdHttpHeader             = "header"       // 指定http的header，默认为空
	CmdHttpCookies            = "cookies"      // 指定http的cookies，默认为空
	CmdHttpResponseOutput     = "out"          // http响应结果的输出路径，默认在控制台输出
)
