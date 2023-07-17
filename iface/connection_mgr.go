package iface

// 连接管理器接口
type IConnectionMgr interface {

	// 添加一个连接
	Add(IConnection)

	// 删除一个连接
	Remove(IConnection)
	Get(string) (IConnection, error)
	Len() int
}
