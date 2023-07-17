package gcore

// web socket 连接
type ConnectionWS struct {
	// 连接 id (唯一标识，每一个连接都分配一个唯一标识)
	id string

	// 连接套接字
	conn any

	// 网络连接类型
	network string

	// 本机地址
	loaclAddr string

	// 远程地址
	remoteAddr string

	// 当前连接的关闭状态
	isClosed bool
}
