package gcore

import (
	"gnet/iface"
	"sync"
)

// 房间管理
type RoomMgr struct {

	// 所有房间
	list map[string]iface.IRoom

	// 读写锁
	lock sync.RWMutex
}
