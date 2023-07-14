package gcore

import (
	"errors"
	"fmt"
	"gnet/gface"
	"log"
	"sync"
)

// 连接管理
type ConnectionManager struct {

	// 所有链接
	list map[string]gface.IConnection

	// 读写锁
	lock sync.RWMutex
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		list: make(map[string]gface.IConnection),
	}
}

func (mgr *ConnectionManager) Add(conn gface.IConnection) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	mgr.list[conn.GetId()] = conn
	log.Println("ConnectionManager Add ", conn.GetId())
}

// 通过 链接 id 进行删除
func (mgr *ConnectionManager) Remove(conn gface.IConnection) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	delete(mgr.list, conn.GetId())
	log.Println("ConnectionManager Remove ", conn.GetId())
}

// 获取一个链接
func (mgr *ConnectionManager) Get(connId string) (gface.IConnection, error) {
	mgr.lock.RLocker()
	defer mgr.lock.RUnlock()

	if conn, ok := mgr.list[connId]; ok {
		return conn, nil
	}

	return nil, errors.New(fmt.Sprintln("ConnectionManager.Get", connId, "not found"))
}

// 连接池大小
func (mgr *ConnectionManager) Len() int {
	mgr.lock.RLocker()
	defer mgr.lock.RUnlock()

	len := len(mgr.list)
	return len
}
