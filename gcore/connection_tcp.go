package gcore

import (
	"context"
	"fmt"
	"gnet/iface"
	"log"
	"net"
	"time"
)

type ConnectionTcp struct {

	// 连接 id (唯一标识，每一个连接都分配一个唯一标识)
	id string

	// 连接套接字
	netConn net.Conn

	// 连接的网络方式
	network string

	// 本机地址
	loaclAddr string

	// 远程地址
	remoteAddr string

	// 当前连接的关闭状态
	isClosed bool

	// 告知该连接已经退出
	ctx    context.Context
	cancel context.CancelFunc

	// 连接管理器
	connMgr iface.IConnectionMgr
}

func NewServerConn(server iface.IServer, netConn net.Conn) iface.IConnection {

	return &ConnectionTcp{
		id:         "server111",
		netConn:    netConn,
		network:    netConn.LocalAddr().Network(),
		loaclAddr:  netConn.LocalAddr().String(),
		remoteAddr: netConn.RemoteAddr().String(),
		connMgr:    server.GetConnMgr(),
	}
}

func NewClientConn(netConn net.Conn) iface.IConnection {
	return &ConnectionTcp{
		id:         "client111",
		netConn:    netConn,
		network:    netConn.LocalAddr().Network(),
		loaclAddr:  netConn.LocalAddr().String(),
		remoteAddr: netConn.RemoteAddr().String(),
	}
}

func (c *ConnectionTcp) GetId() string {
	return c.id
}

func (c *ConnectionTcp) IsClosed() bool {
	return c.isClosed
}

func (c *ConnectionTcp) Start() {
	defer log.Println("[Info] Start() exit! id =", c.id)
	defer func() {
		if err := recover(); err != nil {
			log.Println("[Error] Start() recover", err)
		}
	}()

	log.Println("[Info] Start() start id =", c.id)
	c.ctx, c.cancel = context.WithCancel(context.Background())

	go c.StartReader()
	go c.StartWriter()

	select {
	case <-c.ctx.Done(): // 连接已经断开
		c.destory()
		return
	}
}

// 读消息Goroutine，用于从客户端中读取数据
func (c *ConnectionTcp) StartReader() {
	defer log.Println("[Info] StartReader() exit id =", c.id)
	defer c.Stop()
	defer func() {
		if err := recover(); err != nil {
			log.Println("[Error] StartReader() recover ", err)
		}
	}()

	for {
		buffer := make([]byte, 1024)

		// 连接中读取数据
		_, err := c.netConn.Read(buffer)
		if err != nil {
			log.Println("[Error] StartReader() read socket ", err)
			return
		}

		fmt.Println("收到", c.remoteAddr, "发送的数据:", string(buffer))
	}
}

// 写消息 Goroutine， 用户将数据发送给客户端
func (c *ConnectionTcp) StartWriter() {
	defer log.Println("[Info] StartWriter() exit id =", c.id)
	defer func() {
		if err := recover(); err != nil {
			log.Println("[Error] StartWriter() recover ", err)
		}
	}()

	for {

		select {
		case <-c.ctx.Done():
			return
		default:

			time.Sleep(1 * time.Second)

			if c.isClosed == true {
				return
			}

			msg := "hell world 你好！！"
			_, err := c.netConn.Write([]byte(msg))
			if err != nil {
				log.Println("[Error] StartReader() write socket ", err)
				return
			}

		}

	}
}

// 停止连接，结束当前连接状态
func (c *ConnectionTcp) Stop() {
	c.cancel()
}

func (c *ConnectionTcp) Write(msg iface.IMessage) error {

	return nil
}

func (c *ConnectionTcp) Read() {

}

// 销毁连接数据
func (c *ConnectionTcp) destory() {

	if c.isClosed == true {
		return
	}

	log.Println("[Info] destory() id =", c.id)

	// 关闭网络连接
	c.netConn.Close()

	// 从管理器中删除
	if c.connMgr != nil {
		c.connMgr.Remove(c)
	}

	c.isClosed = true
}
