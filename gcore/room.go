package gcore

import (
	"gnet/iface"
	"sync"
)

// 房间
type Room struct {

	// 房间ID (唯一性)
	id string

	// 所有连接
	list map[string]iface.IConnection

	// 读写锁
	lock sync.RWMutex
}
