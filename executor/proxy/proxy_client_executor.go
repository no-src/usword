package proxy

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/no-src/log"
	"github.com/no-src/usword/console"
	"github.com/no-src/usword/executor"
	base_exec "github.com/no-src/usword/executor/base"
	"github.com/no-src/usword/executor/client"
	"github.com/no-src/usword/executor/const"
	"github.com/no-src/usword/executor/help"
	"github.com/no-src/usword/res/lang"
)

type ProxyClientExecutor struct {
	base_exec.BaseExecutor
	client *client.ClientExecutor
	// 代理服务器地址
	ProxyServer string
	// 代理服务器端口
	ProxyPort int
	// 真实服务器地址
	Server string
	// 真实服务器端口
	Port int
}

// Handle 运行代理客户端
func (exec *ProxyClientExecutor) Handle(params ...string) (result []byte, err error) {
	dict := exec.BuildArgs(params...)
	// 校验代理服务器地址
	exec.ProxyServer = dict[ProxyClientProxyServer]
	if len(exec.ProxyServer) == 0 {
		err = errors.New(InvalidProxyServer)
		result = []byte(err.Error())
		return result, err
	}

	// 校验代理服务器端口
	proxyPort, parseError := strconv.Atoi(dict[ProxyClientProxyPort])
	if parseError != nil {
		err = errors.New(InvalidProxyPort)
		result = []byte(err.Error())
		return result, err
	} else {
		exec.ProxyPort = proxyPort
	}

	// 校验服务器地址
	exec.Server = dict[ProxyClientServer]
	if len(exec.Server) == 0 {
		err = errors.New(InvalidServer)
		result = []byte(err.Error())
		return result, err
	}

	// 校验服务器端口
	port, parseError := strconv.Atoi(dict[ProxyClientPort])
	if parseError != nil {
		err = errors.New(InvalidPort)
		result = []byte(err.Error())
		return result, err
	} else {
		exec.Port = port
	}

	// 使用ClientExecutor与ProxyServer建立连接，并且发送服务器连接信息
	exec.client = &client.ClientExecutor{}
	proxyExec := fmt.Sprintf("exec=server=%s port=%d", exec.Server, exec.Port)
	var clientParams []string
	clientParams = append(clientParams, "server="+exec.ProxyServer)
	clientParams = append(clientParams, "port="+strconv.Itoa(exec.ProxyPort))
	clientParams = append(clientParams, proxyExec)
	result, err = exec.client.Handle(clientParams...)

	// 接收用户输入，发送命令到ProxyServer
	cmd := dict[client.ClientExec]

	// 判断是否为交互执行
	if dict[client.ClientInteractiveExec] == _const.ParamExist {
		// 交互循环执行
		for {
			log.Log("%s,[%s] %s,[%s] %s", lang.ClientInteractive_InputTips, client.ClientExit, lang.ClientInteractive_Exit, client.ClientClear, lang.ClientInteractive_Clear)
			inputReader := bufio.NewReader(os.Stdin)
			line, _, err := inputReader.ReadLine()
			if err != nil {
				if err == io.EOF {
					result = []byte(lang.ClientInteractive_Exit)
					return result, nil
				} else {
					log.Error(err, lang.ClientInteractive_Error_Input)
					continue
				}
			}
			cmd = string(line)
			if cmd == client.ClientExit {
				result = []byte(lang.ClientInteractive_Exit)
				return result, nil
			} else if cmd == client.ClientClear {
				console.Clear()
				continue
			}

			result, err = exec.client.Exec(cmd)
			if err != nil {
				log.Error(err, "[%s]%s,Result:\r\n%s", cmd, lang.ClientInteractive_Error_ExecError, string(result))
				continue
			} else {
				log.Debug("[%s]%s，Result=%s", cmd, lang.ClientInteractive_Finish, string(result))
			}
		}
	} else {
		// 单次执行
		result, err = exec.client.Exec(cmd)
		if err != nil {
			log.Error(err, "[%s]%s,Result:\r\n%s", cmd, lang.ClientInteractive_Error_ExecError, string(result))
		} else {
			log.Debug("[%s]%s，Result=%s", cmd, lang.ClientInteractive_Finish, string(result))
		}
	}

	return result, err
}

func init() {
	register()
	help.RegisterHelpInfo(_const.CmdExecuteProxyClient, lang.ProxyClientExecutor_Desc, man)
}

func man() string {
	argsDict := argsComment()
	msg := fmt.Sprintln(lang.ProxyClientExecutor_Welcome)
	msg += fmt.Sprintln(lang.ProxyClientExecutor_Usage)
	msg += fmt.Sprintln(lang.ProxyClientExecutor_ArgDesc)
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
		ProxyClientProxyServer:       lang.ProxyExecutor_Arg_ProxyClientProxyServer,
		ProxyClientProxyPort:         lang.ProxyExecutor_Arg_ProxyClientProxyPort,
		ProxyClientServer:            lang.ProxyExecutor_Arg_ProxyClientServer,
		ProxyClientPort:              lang.ProxyExecutor_Arg_ProxyClientPort,
		client.ClientInteractiveExec: lang.ProxyExecutor_Arg_ClientInteractiveExec,
		client.ClientExec:            lang.ProxyExecutor_Arg_ClientExec,
	}
	return argsDict
}

// 注册当前IExecutor
func register() {
	executor.RegisterExecutor(_const.CmdExecuteProxyClient, &ProxyClientExecutor{})
}
