package gcore

import (
	"fmt"
	"gnet/gface"
	"gnet/gutil"
	"net"
)

type Client struct {
	name string
	ver  string

	network string
	conn    gface.IConnection
}

func NewClient() *Client {
	return &Client{
		name:    "测试客户端",
		ver:     "v0.1",
		network: "tcp",
	}
}

func (c *Client) Connect(ip string, port int) {

	tcpAddr := &net.TCPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
		Zone: "", //for ipv6, ignore
	}

	conn, err := net.DialTCP(c.network, nil, tcpAddr)
	if err != nil {
		fmt.Printf("[Client %s] 链接错误, %v", c.ver, err)
		return
	}

	connTcp := NewConnectionTcp(conn)
	go connTcp.Start()

	gutil.Signal.Waiting()
}
