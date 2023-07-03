package http

import (
	"errors"
	"fmt"

	"github.com/no-src/usword/executor"
	base_exec "github.com/no-src/usword/executor/base"
	"github.com/no-src/usword/executor/const"
	"github.com/no-src/usword/executor/help"
	"github.com/no-src/usword/http"
	http_const "github.com/no-src/usword/http/const"
	res "github.com/no-src/usword/res/lang"
)

// HttpExecutor HTTP命令执行器
type HttpExecutor struct {
	base_exec.BaseExecutor
}

// Handle 根据参数执行相应的HTTP命令
func (exec *HttpExecutor) Handle(params ...string) (result []byte, err error) {
	dict := exec.BuildArgs(params...)
	httpClient := _http.NewHttpClient()
	method := dict[CmdHttpMethod]
	protocol := dict[CmdHttpProtocol]
	url := dict[CmdHttpRequestUrl]
	contentType := dict[CmdHttpRequestContentType]
	reqData := dict[CmdHttpRequestBody]
	//header := dict[CmdHttpHeader]
	//cookies := dict[CmdHttpCookies]
	if len(method) == 0 {
		method = http_const.GET
	}
	if len(url) == 0 {
		err = errors.New(res.HttpExecutor_Error_MustSetUrl)
		result = []byte(err.Error())
		return result, err
	}
	result, err = httpClient.HttpSend(method, protocol, url, contentType, reqData, nil, nil)
	if err != nil {
		result = []byte(err.Error())
	}
	return result, err
}

func init() {
	register()
	help.RegisterHelpInfo(_const.CmdExecuteHttp, res.HttpExecutor_Desc, man)
}

func man() string {
	argsDict := argsComment()
	msg := fmt.Sprintln(res.HttpExecutor_Welcome)
	msg += fmt.Sprintln(res.HttpExecutor_Usage)
	msg += fmt.Sprintln(res.HttpExecutor_ArgDesc)
	for k, v := range argsDict {
		if len(k) > 7 {
			msg += fmt.Sprintf("\t%s\t%s\n", k, v)
		} else {
			msg += fmt.Sprintf("\t%s\t\t%s\n", k, v)
		}
	}
	return msg
}

func argsComment() map[string]string {
	argsDict := map[string]string{
		CmdHttpRequestUrl:         res.HttpExecutor_Arg_CmdHttpRequestUrl,
		CmdHttpMethod:             res.HttpExecutor_Arg_CmdHttpMethod,
		CmdHttpRequestContentType: res.HttpExecutor_Arg_CmdHttpRequestContentType,
		CmdHttpRequestBody:        res.HttpExecutor_Arg_CmdHttpRequestBody,
		CmdHttpProtocol:           res.HttpExecutor_Arg_CmdHttpProtocol,
		CmdHttpHeader:             res.HttpExecutor_Arg_CmdHttpHeader,
		CmdHttpCookies:            res.HttpExecutor_Arg_CmdHttpCookies,
		CmdHttpResponseOutput:     res.HttpExecutor_Arg_CmdHttpResponseOutput,
	}
	return argsDict
}

// 注册当前IExecutor
func register() {
	executor.RegisterExecutor(_const.CmdExecuteHttp, &HttpExecutor{})
}
