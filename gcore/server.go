package gcore

import (
	"gnet/gface"
	"gnet/gutil"
	"log"
	"net"
)

type Server struct {

	// 服务器配置
	config *Config

	// 连接池
	conns *ConnectionList

	// 房间列表
	rooms map[string]any

	// 所有链接的客户端
	clients map[string]any
}

func NewServer() *Server {
	return &Server{
		config: NewConfig(),
		conns:  NewConnectionList(),
	}
}

func (s *Server) Start() {

	log.Println("Server Start ")

	// 初始化配置

	// 链接监听
	s.listenTcp()

	// 等待
	gutil.Signal.Waiting()
	log.Printf("[SERVE]  server exit, name %s,", s.config.server.name)
}

func (server *Server) Stop() {

}

func (server *Server) Restart() {

}

// 监听客户端链接
func (s *Server) listenTcp() {

	// 获取 tcp 地址
	tcpAddr, err := net.ResolveTCPAddr(s.config.server.network, s.config.server.addr)
	if err != nil {
		log.Println("Error listenTcp ", err)
		return
	}

	// 监听服务器地址
	listener, err := net.ListenTCP(s.config.server.network, tcpAddr)
	if err != nil {
		log.Println("Error listenTcp ", err)
		panic(err)
	}

	// 开启一个新的 goroutine 等待连接
	go func() {

		// 设置服务器最大连接，如果超过最大连接，则等待

		// 阻塞等待客户端建立连接请求
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error ListenTcpConnect listener.Accept ", err)
		}

		// 创建链接对象
		connTcp := NewConnectionTcp(conn)

		// 开启一个 goroutine
		go s.startConn(connTcp)
	}()

}

// 开启一个链接
func (s *Server) startConn(conn gface.IConnection) {

	s.conns.Add(conn)
	conn.Start()
}
