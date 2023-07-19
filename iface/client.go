package iface

type IClient interface {
	// 连接服务器
	Connect(addr string)
	Close()
}
