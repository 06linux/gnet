package gcore

import (
	"fmt"
	"gnet/gbase"
	"gnet/iface"
	"net"
)

type Client struct {
	name string
	ver  string

	network string
	conn    iface.IConnection
}

func NewClient() *Client {
	return &Client{
		name:    "测试客户端",
		ver:     "v0.1",
		network: "tcp",
	}
}

func (c *Client) Connect(ip string, port int) {

	gbase.Log.Init()

	tcpAddr := &net.TCPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
		Zone: "", //for ipv6, ignore
	}

	conn, err := net.DialTCP(c.network, nil, tcpAddr)
	if err != nil {
		fmt.Printf("[Client %s] 连接错误, %v", c.ver, err)
		return
	}

	clientConn := NewClientConn(conn)
	go clientConn.Start()
}
