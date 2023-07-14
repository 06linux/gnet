package gface

// 链接管理器接口
type IConnectionManager interface {

	// 添加一个链接
	Add(IConnection)

	// 删除一个链接
	Remove(IConnection)              // Remove connection
	Get(string) (IConnection, error) // Get a connection by ConnID
	Len() int
}
