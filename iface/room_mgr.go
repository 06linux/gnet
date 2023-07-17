package iface

// 房间管理
type IRoomMgr interface {

	// 添加一个房间
	Add(IRoom)

	// 删除一个房间
	Del(IRoom)

	// 房间是否存在
	Has(IRoom)

	// 房间的数量
	Len() int
}
