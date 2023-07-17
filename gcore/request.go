package gcore

import "gnet/iface"

// 请求数据
type Request struct {

	// 请求的链接
	conn iface.IConnection

	// 请求的消息
	msg iface.IMessage

	id   string // 消息 id
	room string // 发送到那个房间
	data []byte // 消息内容
}
