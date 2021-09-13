package server

import (
	"bufio"
	"fmt"
	"github.com/no-src/log"
	default_conn "github.com/no-src/usword/connector"
	"github.com/no-src/usword/connector/server"
	"github.com/no-src/usword/executor"
	base_exec "github.com/no-src/usword/executor/base"
	"github.com/no-src/usword/executor/const"
	"github.com/no-src/usword/executor/help"
	"github.com/no-src/usword/res/lang"
	"net"
	"strconv"
	"strings"
)

type ServerExecutor struct {
	base_exec.BaseExecutor
}

const (
	ServerArgServerHost = "host" // 服务器主机绑定地址
	ServerArgPort       = "port" // 服务器主机监听端口，默认是9091
)

// HandleFunc 运行USword服务端,并指定处理函数
func (exec *ServerExecutor) HandleFunc(process func(client net.Conn, data []byte), params ...string) (result []byte, err error) {
	dict := exec.BuildArgs(params...)
	ip := dict[ServerArgServerHost]

	port, parseError := strconv.Atoi(dict[ServerArgPort])
	if parseError != nil {
		port = default_conn.DefaultServerPort
	}
	srv := connector.NewServer(ip, port)
	err = srv.Listen()
	if err != nil {
		log.Error(err, lang.ServerExecutor_Error_ListenFailed, ip, port)
		return result, err
	}
	log.Debug(lang.ServerExecutor_ListenSucceed, ip, port)
	err = srv.Accept(process)
	if err != nil {
		log.Error(err, lang.ServerExecutor_Error_AcceptFailed)
	}
	return result, err
}

// Handle 运行USword服务端
func (exec *ServerExecutor) Handle(params ...string) (result []byte, err error) {
	processFunc := func(client net.Conn, data []byte) {
		receiveMsg := string(data)
		log.Debug(lang.ServerExecutor_ReceiveClientCommand, client.RemoteAddr().String(), receiveMsg)
		result, err = exec.Exec(receiveMsg)
		if err != nil {
			result = append(result, default_conn.ErrorIdentity...)
			log.Error(err, lang.ServerExecutor_Error_ExecuteClientCommandFailed, client.RemoteAddr().String(), receiveMsg)
		}
		writer := bufio.NewWriter(client)
		result = append(result, default_conn.EndIdentity...)
		result = append(result, default_conn.CRLFBytes...)
		_, err = writer.Write(result)
		if err != nil {
			log.Error(err, lang.ServerExecutor_Error_ServerReponseFailed, client.RemoteAddr().String(), receiveMsg)
		} else {
			log.Debug(lang.ServerExecutor_Info_ServerReponseSuccess+", Result=\r\n%s", client.RemoteAddr().String(), receiveMsg, string(result))
		}
		writer.Flush()
	}
	result, err = exec.HandleFunc(processFunc, params...)
	return result, err
}

// Exec 执行命令
func (exec *ServerExecutor) Exec(cmd string) (result []byte, err error) {
	ed := executor.ExecDistribute{}
	params := strings.Fields(cmd)
	nextExec := ed.GetExecutor(params...)
	result, err = nextExec.Handle(params...)
	return result, err
}

func init() {
	register()
	help.RegisterHelpInfo(_const.CmdExecuteServer, lang.ServerExecutor_Desc, man)
}

func man() string {
	argsDict := argsComment()
	msg := fmt.Sprintln(lang.ServerExecutor_Welcome)
	msg += fmt.Sprintln(lang.ServerExecutor_Usage)
	msg += fmt.Sprintln(lang.ServerExecutor_ArgDesc)
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
		ServerArgServerHost: lang.ServerExecutor_Arg_ServerArgServerHost,
		ServerArgPort:       lang.ServerExecutor_Arg_ServerArgPort,
	}
	return argsDict
}

// 注册当前IExecutor
func register() {
	executor.RegisterExecutor(_const.CmdExecuteServer, &ServerExecutor{})
}
