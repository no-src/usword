package connector

import (
	"bufio"
	"bytes"
	"fmt"
	"net"

	"github.com/no-src/log"
	"github.com/no-src/usword/connector"
	"github.com/no-src/usword/res/lang"
)

// Client 网络连接客户端
type Client struct {
	network   string
	host      string
	port      int
	innerConn net.Conn
}

// NewClient 创建一个Client实例用于网络连接
func NewClient(host string, port int) (client *Client) {
	client = &Client{}
	client.host = host
	client.port = port
	client.network = "tcp"
	return client
}

// Connect 建立连接
func (client *Client) Connect() (err error) {
	address := fmt.Sprintf("%s:%d", client.host, client.port)
	client.innerConn, err = net.Dial(client.network, address)
	if err != nil {
		log.Error(err, lang.ConnClient_Error_ClientConnectFailed)
	}
	return err
}

// Connected 客户端是否与服务器建立连接
func (client *Client) Connected() (connected bool) {
	// todo 该实现不能准确判断服务器连接是否正常
	connected = client.innerConn != nil && client.innerConn.RemoteAddr() != nil
	return connected
}

// Write 向服务端发送数据
func (client *Client) Write(data []byte) (err error) {
	writer := bufio.NewWriter(client.innerConn)
	_, err = writer.Write(data)
	if err != nil {
		log.Error(err, lang.ConnClient_Error_ClientWriteFailed)
		return err
	}
	err = writer.Flush()
	if err != nil {
		log.Error(err, lang.ConnClient_Error_ClientFlushFailed)
		return err
	}
	return err
}

// ReadAll 读取服务端响应数据
func (client *Client) ReadAll() (result []byte, err error) {
	reader := bufio.NewReader(client.innerConn)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			return result, err
		}
		isEnd := false
		endIdentity := connector.EndIdentity
		hasError := false
		if bytes.HasSuffix(line, endIdentity) {
			isEnd = true
			if bytes.HasSuffix(line, connector.ErrorEndIdentity) {
				endIdentity = connector.ErrorEndIdentity
				hasError = true
			}
			line = line[:len(line)-len(endIdentity)]
		}

		result = append(result, line...)
		result = append(result, connector.CRLFBytes...)

		if isEnd {
			if hasError {
				err = connector.ServerExecuteError
			}
			return result, err
		}
	}
}
