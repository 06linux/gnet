package iface

// 房间接口
type IRoom interface {

	// 房间 id
	GetId() string

	// 添加一个连接
	Add(IConnection)

	// 删除一个连接
	Del(IConnection)

	// 连接是否存在
	Has(IConnection)

	// 清除所有
	Clear()

	// 当前房间拥有连接的数量
	Len() int
}
