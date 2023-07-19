package gcore

import "gnet/iface"

// 请求数据
type Request struct {

	// 请求的链接
	conn iface.IConnection

	// 请求的消息
	msg iface.IMessage
}
