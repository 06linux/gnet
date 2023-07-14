package gcore

import (
	"fmt"
	"gnet/gface"
	"log"
	"net"
	"time"
)

type ConnectionTcp struct {

	// 连接 id (唯一标识，每一个链接都分配一个唯一标识)
	id string

	// 链接套接字
	netconn net.Conn

	// 链接的网络方式
	network string

	// 本机地址
	loaclAddr string

	// 远程地址
	remoteAddr string

	// 当前连接的关闭状态
	isClosed bool
}

func NewConnectionTcp(netConn net.Conn) *ConnectionTcp {

	return &ConnectionTcp{
		id:         "test111",
		netconn:    netConn,
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

	go c.StartReader()
	go c.StartWriter()

}

// 读消息Goroutine，用于从客户端中读取数据
func (c *ConnectionTcp) StartReader() {

	for {

		buffer := make([]byte, 1024)

		// 链接中读取数据
		_, err := c.netconn.Read(buffer)
		if err != nil {
			log.Println("Error! ConnectionTcp StartReader", err)
			return
		}

		fmt.Println("收到", c.remoteAddr, "发送的数据:", string(buffer))
	}

}

// 写消息Goroutine， 用户将数据发送给客户端
func (c *ConnectionTcp) StartWriter() {
	for {

		time.Sleep(1 * time.Second)

		msg := "hell world"
		_, err := c.netconn.Write([]byte(msg))
		if err != nil {
			log.Println("Error! ConnectionTcp StartWriter", err)
			return
		}
	}
}

func (c *ConnectionTcp) Stop() {

}

func (c *ConnectionTcp) Write(msg gface.IMessage) error {

	return nil
}

func (c *ConnectionTcp) Read() {

}
