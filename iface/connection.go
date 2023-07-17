package iface

// 一个连接
type IConnection interface {

	// 连接 id
	GetId() string

	// 启动连接，让当前连接开始工作
	Start()

	// 停止连接，结束当前连接状态
	Stop()

	// 连接是否已经关闭
	IsClosed() bool

	// 写消息
	Write(msg IMessage) error

	// 读信息
	Read()
}
