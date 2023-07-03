package connector

import (
	"bufio"
	"net"
	"sync/atomic"

	"github.com/no-src/log"
	"github.com/no-src/usword/res/lang"
)

// Server 网络连接服务端
type Server struct {
	// 网络通信方式
	network string
	// 服务器监听的IP
	ip net.IP
	// 服务器监听端口
	port int
	// 服务器监听器
	listener *net.TCPListener
	// 当前服务器的客户端连接数
	clientCount int32
}

// NewServer 新建一个Server实例，用于网络服务端监听
func NewServer(ip string, port int) (srv *Server) {
	srv = &Server{}
	srv.ip = net.ParseIP(ip)
	srv.port = port
	srv.network = "tcp"
	return srv
}

// Listen 开始监听
func (srv *Server) Listen() (err error) {
	addr := &net.TCPAddr{
		IP:   srv.ip,
		Port: srv.port,
	}
	srv.listener, err = net.ListenTCP(srv.network, addr)
	return err
}

// Accept 接收客户端请求进行响应处理
// process 为客户端指定响应函数，client为客户端信息，data为客户端请求消息
func (srv *Server) Accept(process func(client net.Conn, data []byte)) (err error) {
	for {
		clientConn, err := srv.listener.Accept()
		if err != nil {
			continue
		}
		srv.addClient(clientConn)

		go func() {
			reader := bufio.NewReader(clientConn)
			for {
				line, _, err := reader.ReadLine()
				if err != nil {
					clientConn.Close()
					srv.removeClient(clientConn)
					log.Error(err, lang.ConnServer_ClientClosed, clientConn.RemoteAddr().String(), srv.clientCount)
					//连接关闭 停止数据接收
					return
				} else {
					process(clientConn, line)
				}
			}

		}()
	}

	return err
}

// addClient 添加一个客户端
func (srv *Server) addClient(conn net.Conn) (clientCount int32, err error) {
	atomic.AddInt32(&srv.clientCount, 1)
	clientCount = srv.clientCount
	log.Debug(lang.ConnServer_ClientConnected, conn.RemoteAddr().String(), clientCount)
	return srv.clientCount, err
}

// removeClient 移除一个客户端
func (srv *Server) removeClient(conn net.Conn) (clientCount int32, err error) {
	atomic.AddInt32(&srv.clientCount, -1)
	clientCount = srv.clientCount
	log.Debug(lang.ConnServer_ClientRemoved, conn.RemoteAddr().String(), clientCount)
	return srv.clientCount, err
}

// GetClientCount 获取当前客户端连接总数
func (srv *Server) GetClientCount() int {
	return int(srv.clientCount)
}
