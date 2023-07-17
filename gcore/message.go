package gcore

type Message struct {
	id    string // 消息 id
	room  string // 发送到那个房间
	route string // 消息路由
	data  []byte // 消息内容
}
