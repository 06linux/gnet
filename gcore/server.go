package gcore

import (
	"gnet/gbase"
	"gnet/iface"
	"log"
	"net"
)

type Server struct {

	// 服务器配置
	config *Config

	// 连接管理器
	connMgr *ConnectionMgr

	// 房间管理器
	roomMgr map[string]any
}

func NewServer() *Server {
	return &Server{
		config:  NewConfig(),
		connMgr: NewConnectionMgr(),
	}
}

func (s *Server) Start() {

	gbase.Log.Init()
	log.Println("[Server] Start ")

	// 初始化配置

	// 连接监听
	s.listenTcp()
}

func (s *Server) Stop() {

}

func (s *Server) Restart() {

}

func (s *Server) GetConnMgr() iface.IConnectionMgr {
	return s.connMgr
}

// 监听客户端连接
func (s *Server) listenTcp() {

	// 获取 tcp 地址
	tcpAddr, err := net.ResolveTCPAddr(s.config.server.network, s.config.server.addr)
	if err != nil {
		log.Println("[Error] listenTcp() ", err)
		return
	}

	// 监听服务器地址
	listener, err := net.ListenTCP(s.config.server.network, tcpAddr)
	if err != nil {
		log.Println("[Error] listenTcp() ", err)
		panic(err)
	}

	// 开启一个新的 goroutine 等待连接
	go func() {

		for {
			// 设置服务器最大连接，如果超过最大连接，则等待

			// 阻塞等待客户端建立连接请求
			conn, err := listener.Accept()
			if err != nil {
				log.Println("[Error] listenTcp()  Accept ", err)
			}

			// 创建连接对象
			serverConn := NewServerConn(s, conn)

			// 开启一个 goroutine
			go s.startConn(serverConn)
		}

	}()

}

// 开启一个连接
func (s *Server) startConn(conn iface.IConnection) {

	log.Println("[Info] startConn()  ")

	s.connMgr.Add(conn)
	conn.Start()
}
