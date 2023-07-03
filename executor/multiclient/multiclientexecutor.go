package multiclient

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/no-src/log"
	"github.com/no-src/usword/console"
	"github.com/no-src/usword/executor"
	base_exec "github.com/no-src/usword/executor/base"
	"github.com/no-src/usword/executor/client"
	"github.com/no-src/usword/executor/const"
	"github.com/no-src/usword/executor/help"
	"github.com/no-src/usword/res/lang"
)

type MultiClientExecutor struct {
	base_exec.BaseExecutor
	clients []*client.ClientExecutor
}

const (
	ClientArgServers = "servers" // 客户端要连接的1个或多个服务器主机地址及参数设置，格式如下127.0.0.1:8185?mode=online，多个地址间用英文逗号分隔
)

// parseClients 解析servers字符串,例如：127.0.0.1:8185?mode=online,127.0.0.1:8187?mode=online
func (exec *MultiClientExecutor) parseClients(servers string) (clients []*client.ClientExecutor) {
	if len(servers) == 0 {
		return clients
	}
	serverList := strings.Split(servers, ",")
	for _, server := range serverList {
		if len(server) == 0 {
			continue
		}
		c := &client.ClientExecutor{}
		// 如果是离线模式，直接处理，不需要解析主机端口
		if strings.Contains(server, "mode=offline") {
			c.Mode = client.ClientRunModeOffline
			clients = append(clients, c)
			continue
		}
		colonIndex := strings.Index(server, ":")
		if colonIndex <= 0 {
			continue
		}

		c.Server = server[:colonIndex]

		qusIndex := strings.Index(server, "?")
		if qusIndex > 0 {
			port := server[colonIndex+1 : qusIndex]
			c.Port, _ = strconv.Atoi(port)
			params := server[qusIndex+1:]
			values, err := url.ParseQuery(params)
			if err == nil {
				modes := values[client.ClientRunMode]
				if len(modes) > 0 {
					c.Mode = modes[0]
				}
			}
		} else {
			port := server[colonIndex+1:]
			c.Port, _ = strconv.Atoi(port)
		}
		if len(c.Mode) == 0 {
			c.Mode = client.ClientRunModeOnline
		}
		clients = append(clients, c)
	}
	return clients
}

// Handle 运行USword客户端
func (exec *MultiClientExecutor) Handle(params ...string) (result []byte, err error) {
	dict := exec.BuildArgs(params...)
	servers := dict[ClientArgServers]
	if len(servers) == 0 {
		result = []byte(lang.MultiClientExecutor_Error_ClientArgServers_NotFound + ClientArgServers)
		return result, nil
	}
	exec.clients = exec.parseClients(servers)
	if len(exec.clients) == 0 {
		result = []byte(lang.MultiClientExecutor_Error_ClientArgServers_NotValid + ClientArgServers)
		return result, nil
	}

	err = exec.Connect()
	if err != nil {
		result = []byte(err.Error())
		return result, err
	}

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
			exec.Exec(cmd)
		}
	} else {
		// 单次执行
		exec.Exec(cmd)
	}

	return result, err
}

// Connect 连接远程服务器
func (exec *MultiClientExecutor) Connect() (err error) {
	for i, v := range exec.clients {
		err = v.Connect()
		if err != nil {
			errMsg := fmt.Sprintf("[%d]%s[%s:%d] mode=%s", i, lang.MultiClientExecutor_Error_ConnectServerFailed, v.Server, v.Port, v.Mode)
			log.Error(err, errMsg)
			return err
		} else {
			log.Debug("[%d]%s[%s:%d]mode=%s", i, lang.MultiClientExecutor_ConnectServerSuccess, v.Server, v.Port, v.Mode)
		}
	}
	return nil
}

// Exec 执行命令
func (exec *MultiClientExecutor) Exec(cmd string) {
	for i, c := range exec.clients {
		result, err := c.Exec(cmd)
		if err != nil {
			log.Error(err, "[%s]%s,Result:\r\n%s", cmd, lang.ClientInteractive_Error_ExecError, string(result))
			continue
		} else {
			log.Debug("[%d][%s:%d]%s[%s]%s，Result=%s", i, c.Server, c.Port, lang.ClientInteractive_Cmd, cmd, lang.ClientInteractive_Finish, string(result))
		}
	}
}
func init() {
	register()
	help.RegisterHelpInfo(_const.CmdExecuteMultiClient, lang.MultiClientExecutor_Desc, man)
}

func man() string {
	argsDict := argsComment()
	msg := fmt.Sprintln(lang.MultiClientExecutor_Welcome)
	msg += fmt.Sprintln(lang.MultiClientExecutor_Usage)
	msg += fmt.Sprintln(lang.MultiClientExecutor_ArgDesc)
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
		ClientArgServers:             lang.MultiClientExecutor_Arg_ClientArgServers,
		client.ClientRunMode:         fmt.Sprintf(lang.MultiClientExecutor_Arg_ClientRunMode, client.ClientRunModeOnline, client.ClientRunModeOffline, client.ClientRunModeOnline),
		client.ClientInteractiveExec: lang.MultiClientExecutor_Arg_ClientInteractiveExec,
		client.ClientExec:            lang.MultiClientExecutor_Arg_ClientExec,
	}
	return argsDict
}

// 注册当前IExecutor
func register() {
	executor.RegisterExecutor(_const.CmdExecuteMultiClient, &MultiClientExecutor{})
}
