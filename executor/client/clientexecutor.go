package client

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/no-src/log"
	default_conn "github.com/no-src/usword/connector"
	"github.com/no-src/usword/connector/client"
	"github.com/no-src/usword/console"
	"github.com/no-src/usword/executor"
	base_exec "github.com/no-src/usword/executor/base"
	"github.com/no-src/usword/executor/const"
	"github.com/no-src/usword/executor/help"
	"github.com/no-src/usword/res/lang"
)

// ClientExecutor 网络连接客户端执行器
type ClientExecutor struct {
	base_exec.BaseExecutor
	client *connector.Client
	// 是否为离线模式
	isOffline bool
	// 连接的服务器地址
	Server string
	// 连接的服务器端口
	Port int
	// 客户端运行模式
	Mode string
	// 尝试重新连接的次数 <0则不进行重试，0则使用默认值DefaultRetryCount
	RetryCount int
}

// ClientExecError 执行错误信息
type ClientExecError error

// ConnAbortErr 连接中断异常
var ConnAbortErr ClientExecError = errors.New(lang.ClientExecutor_Error_ConnAbortErr)

const (
	DefaultRetryCount     = 3         // 默认尝试重新连接的次数
	ClientArgServer       = "server"  // 客户端要连接的服务器主机地址，默认是127.0.0.1
	ClientArgPort         = "port"    // 客户端要连接的服务器主机端口，默认是9091
	ClientRunMode         = "mode"    // 客户端运行模式
	ClientRunModeOnline   = "online"  // 在线模式，与远程服务器建立连接（默认模式）
	ClientRunModeOffline  = "offline" // 离线模式，本机自行接管所有命令，不进行远程服务器连接
	ClientInteractiveExec = "-i"      // 交互执行
	ClientExec            = "exec"    // 提交给客户端执行的命令，如果是online模式，会转发给server执行
	ClientExit            = "exit"    // 退出客户端交互
	ClientClear           = "clear"   // 清空当前界面
)

// Handle 运行USword客户端
func (exec *ClientExecutor) Handle(params ...string) (result []byte, err error) {
	dict := exec.BuildArgs(params...)
	exec.Server = dict[ClientArgServer]
	if len(exec.Server) == 0 {
		exec.Server = default_conn.DefaultServerHost
	}

	port, parseError := strconv.Atoi(dict[ClientArgPort])
	if parseError != nil {
		exec.Port = default_conn.DefaultServerPort
	} else {
		exec.Port = port
	}

	exec.Mode = dict[ClientRunMode]
	if len(exec.Mode) == 0 {
		exec.Mode = ClientRunModeOnline
	}

	err = exec.Connect()
	if err != nil {
		errMsg := fmt.Sprintf("%s[%s:%d] mode=%s", lang.ClientExecutor_Error_ConnAbortErr, exec.Server, exec.Port, exec.Mode)
		result = []byte(errMsg)
		log.Error(err, errMsg)
		return result, err
	}

	cmd := dict[ClientExec]

	// 判断是否为交互执行
	if dict[ClientInteractiveExec] == _const.ParamExist {
		// 交互循环执行
		for {
			log.Log("%s,[%s] %s,[%s] %s", lang.ClientInteractive_InputTips, ClientExit, lang.ClientInteractive_Exit, ClientClear, lang.ClientInteractive_Clear)
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
			if cmd == ClientExit {
				result = []byte(lang.ClientInteractive_Exit)
				return result, nil
			} else if cmd == ClientClear {
				console.Clear()
				continue
			}

			result, err = exec.Exec(cmd)
			if err != nil {
				log.Error(err, "[%s]%s,Result:\r\n%s", cmd, lang.ClientInteractive_Error_ExecError, string(result))
				continue
			} else {
				log.Debug("[%s]%s,Result:\r\n%s", cmd, lang.ClientInteractive_Finish, string(result))
			}
		}
	} else {
		// 单次执行
		result, err = exec.Exec(cmd)
		if err != nil {
			log.Error(err, "[%s]%s,Result:\r\n%s", cmd, lang.ClientInteractive_Error_ExecError, string(result))
		} else {
			log.Debug("[%s]%s，Result=%s", cmd, lang.ClientInteractive_Finish, string(result))
		}
	}
	return result, err
}

// Connect 连接远程服务器
func (exec *ClientExecutor) Connect() (err error) {
	if exec.Mode == ClientRunModeOffline {
		exec.isOffline = true
		return nil
	}
	if exec.client == nil {
		exec.client = connector.NewClient(exec.Server, exec.Port)
	}
	err = exec.client.Connect()
	if err != nil {
		log.Error(err, "%s[%s:%d]", lang.ClientExecutor_ServerConnect_Failed, exec.Server, exec.Port)
	} else {
		log.Debug("%s[%s:%d]", lang.ClientExecutor_ServerConnect_Success, exec.Server, exec.Port)
	}
	return err
}

// ReConnect 尝试重新连接
func (exec *ClientExecutor) ReConnect() (err error) {
	tryCount := exec.RetryCount
	if tryCount == 0 {
		exec.RetryCount = DefaultRetryCount
		tryCount = exec.RetryCount
	}
	for i := 0; i < tryCount; i++ {
		err = exec.Connect()
		if err == nil {
			log.Debug(lang.ClientExecutor_ServerConnectRetry_Success+"[%s:%d]", i+1, exec.Server, exec.Port)
			return err
		} else {
			log.Error(err, lang.ClientExecutor_ServerConnectRetry_Failed+"[%s:%d]", i+1, exec.Server, exec.Port)
		}
	}
	return err
}

// Exec 执行命令
func (exec *ClientExecutor) Exec(cmd string) (result []byte, err error) {
	result, err = exec.exec(cmd)
	if err != nil {
		errObj := interface{}(err)
		switch errObj.(type) {
		case *net.OpError:
			opErr := errObj.(*net.OpError)
			if strings.Contains(opErr.Error(), "An existing connection was forcibly closed by the remote host") ||
				strings.Contains(opErr.Error(), "An established connection was aborted by the software in your host machine") {
				// 尝试重新连接，如果连接成功，并且是再发送消息的时候连接异常，则重新发送命令
				err = exec.ReConnect()
				// 判断opErr.Op == "write"，防止发送成功，但是接收数据时连接中断，重连成功后，重复发起一次请求，产生潜在的风险
				if err == nil && opErr.Op == "write" {
					result, err = exec.exec(cmd)
				}
			}
			break
		case ClientExecError:
			clientErr := errObj.(ClientExecError)
			if clientErr == ConnAbortErr {
				// 尝试重新连接
				err = exec.ReConnect()
				if err == nil {
					// 当前没有发送过命令，直接重新发送命令
					result, err = exec.exec(cmd)
				}
			}
			break
		default:
			break
		}
	}
	return result, err
}

// exec 执行命令
func (exec *ClientExecutor) exec(cmd string) (result []byte, err error) {
	// 如果在非离线模式下，则将命令发送到远程服务器，否则直接执行命令
	if exec.isOffline == false {
		if exec.client != nil && exec.client.Connected() {
			cmd += default_conn.CRLF
			err = exec.client.Write([]byte(cmd))
			if err != nil {
				log.Error(err, lang.ClientExecutor_ServerWrite_Failed)
				return result, err
			}
			result, err = exec.client.ReadAll()
			if err != nil {
				if err != default_conn.ServerExecuteError {
					log.Error(err, lang.ClientExecutor_ServerRead_Failed)
				}
				return result, err
			}
		} else {
			err = ConnAbortErr
			log.Error(err, lang.ClientExecutor_Error_ConnAbortErr)
			return result, err
		}
	} else {
		result, err = exec.ExecOffline(cmd)
	}
	return result, err
}

// ExecOffline offline模式，直接执行执行的命令
func (exec *ClientExecutor) ExecOffline(cmd string) (result []byte, err error) {
	ed := executor.ExecDistribute{}
	params := strings.Fields(cmd)
	nextExec := ed.GetExecutor(params...)
	result, err = nextExec.Handle(params...)
	return result, err
}

func init() {
	register()
	help.RegisterHelpInfo(_const.CmdExecuteClient, lang.ClientExecutor_Desc, man)
}

func man() string {
	argsDict := argsComment()
	msg := fmt.Sprintln(lang.ClientExecutor_Welcome)
	msg += fmt.Sprintln(lang.ClientExecutor_Usage)
	msg += fmt.Sprintln(lang.ClientExecutor_ArgDesc)
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
		ClientArgServer:       lang.ClientExecutor_Arg_ClientArgServer,
		ClientArgPort:         lang.ClientExecutor_Arg_ClientArgPort,
		ClientRunMode:         fmt.Sprintf(lang.ClientExecutor_Arg_ClientRunMode, ClientRunModeOnline, ClientRunModeOffline, ClientRunModeOnline),
		ClientInteractiveExec: lang.ClientExecutor_Arg_ClientInteractiveExec,
		ClientExec:            lang.ClientExecutor_Arg_ClientExec,
	}
	return argsDict
}

// 注册当前IExecutor
func register() {
	executor.RegisterExecutor(_const.CmdExecuteClient, &ClientExecutor{})
}
