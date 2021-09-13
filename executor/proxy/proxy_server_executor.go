package proxy

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/no-src/log"
	default_conn "github.com/no-src/usword/connector"
	"github.com/no-src/usword/executor"
	base_exec "github.com/no-src/usword/executor/base"
	"github.com/no-src/usword/executor/client"
	"github.com/no-src/usword/executor/const"
	"github.com/no-src/usword/executor/help"
	"github.com/no-src/usword/executor/server"
	"github.com/no-src/usword/res/lang"
	"net"
	"strconv"
	"strings"
)

type ProxyServerExecutor struct {
	base_exec.BaseExecutor
	server *server.ServerExecutor
	// 代理服务器地址
	ProxyServer string
	// 代理服务器端口
	ProxyPort int
	// 代理字典
	ProxyMap map[string]*client.ClientExecutor
}

// parseProxyConfig 解析代理配置 server=127.0.0.1 port=8002
func (exec *ProxyServerExecutor) parseProxyConfig(cmd string) (server string, port int, err error) {
	if len(cmd) == 0 {
		err = errors.New(lang.ProxyExecutor_Error_ProxyConfigIsEmpty)
		return
	}
	fields := strings.Fields(cmd)
	for _, k := range fields {
		kv := strings.Split(k, "=")
		if len(kv) == 2 {
			if kv[0] == "server" {
				server = kv[1]
			} else if kv[0] == "port" {
				portStr := kv[1]
				portTmp, parseErr := strconv.Atoi(portStr)
				if parseErr == nil {
					port = portTmp
				}
			}
		}
	}
	if len(server) > 0 && port > 0 {
		return server, port, nil
	} else {
		err = errors.New(lang.ProxyExecutor_Error_NotFoundLegalProxyConfig)
	}
	return server, port, err
}

// Handle 运行USword代理服务端
func (exec *ProxyServerExecutor) Handle(params ...string) (result []byte, err error) {
	dict := exec.BuildArgs(params...)
	exec.ProxyServer = dict[ProxyClientProxyServer]
	proxyPort, parseError := strconv.Atoi(dict[ProxyClientProxyPort])
	if parseError == nil {
		exec.ProxyPort = proxyPort
	}
	exec.server = &server.ServerExecutor{}
	var serverParams []string
	serverParams = append(serverParams, "server="+exec.ProxyServer)
	serverParams = append(serverParams, "port="+strconv.Itoa(exec.ProxyPort))
	processFunc := func(clientConn net.Conn, data []byte) {
		receiveMsg := string(data)
		log.Debug(lang.ProxyServerExecutor_ReceiveClientCommand, clientConn.RemoteAddr().String(), receiveMsg)
		// 直接转发到真实服务器
		// 解析接收到的消息，是否包含真是服务器的配置信息
		proxyKey := fmt.Sprintf("%s:%d", exec.ProxyServer, exec.ProxyPort)
		clientExec := exec.ProxyMap[proxyKey]
		server, port, parseErr := exec.parseProxyConfig(receiveMsg)
		if parseErr == nil {
			// 解析成功 重置代理信息
			clientExec = &client.ClientExecutor{}
			var clientParams []string
			clientParams = append(clientParams, "server="+server)
			clientParams = append(clientParams, "port="+strconv.Itoa(port))
			result, err = clientExec.Handle(clientParams...)
			if err == nil {
				if exec.ProxyMap == nil {
					exec.ProxyMap = make(map[string]*client.ClientExecutor)
				}
				exec.ProxyMap[proxyKey] = clientExec
			}
		} else {
			// 解析失败 则尝试把客户端接收到命令发给真实服务器执行
			if clientExec == nil {
				err = errors.New(lang.ProxyExecutor_Error_NotFoundRealServer)
				result = []byte(err.Error())
			} else {
				result, err = clientExec.Exec(receiveMsg)
			}

		}
		if err != nil {
			result = append(result, default_conn.ErrorIdentity...)
			log.Error(err, lang.ProxyExecutor_Error_ExecuteClientCommandFailed, clientConn.RemoteAddr().String(), receiveMsg)
		}
		writer := bufio.NewWriter(clientConn)
		result = append(result, default_conn.EndIdentity...)
		result = append(result, default_conn.CRLFBytes...)
		_, err = writer.Write(result)
		if err != nil {
			log.Error(err, lang.ProxyExecutor_Error_ServerReponseFailed, clientConn.RemoteAddr().String(), receiveMsg)
		} else {
			log.Debug(lang.ProxyExecutor_Info_ServerReponseSuccess+"，Result=\r\n%s", clientConn.RemoteAddr().String(), receiveMsg, string(result))
		}
		writer.Flush()
	}
	result, err = exec.server.HandleFunc(processFunc, serverParams...)
	return result, err
}

func init() {
	registerServer()
	help.RegisterHelpInfo(_const.CmdExecuteProxyServer, lang.ProxyServerExecutor_Desc, serverMan)
}

func serverMan() string {
	argsDict := argsCommentServer()
	msg := fmt.Sprintln(lang.ProxyServerExecutor_Welcome)
	msg += fmt.Sprintln(lang.ProxyServerExecutor_Usage)
	msg += fmt.Sprintln(lang.ProxyServerExecutor_ArgDesc)
	for k, v := range argsDict {
		if len(k) > 7 {
			msg += fmt.Sprintf("\t%s\t%s\n", k, v)
		} else {
			msg += fmt.Sprintf("\t%s\t\t%s\n", k, v)
		}
	}
	return msg
}

func argsCommentServer() map[string]string {
	argsDict := map[string]string{
		ProxyClientProxyServer: lang.ProxyServerExecutor_Arg_ProxyClientProxyServer,
		ProxyClientProxyPort:   lang.ProxyServerExecutor_Arg_ProxyClientProxyPort,
	}
	return argsDict
}

// 注册当前IExecutor
func registerServer() {
	executor.RegisterExecutor(_const.CmdExecuteProxyServer, &ProxyServerExecutor{})
}
