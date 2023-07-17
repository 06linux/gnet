package gcore

import (
	"errors"
	"fmt"
	"gnet/iface"
	"log"
	"sync"
)

// 连接管理
type ConnectionMgr struct {

	// 所有连接
	list map[string]iface.IConnection

	// 读写锁
	lock sync.RWMutex
}

func NewConnectionMgr() *ConnectionMgr {
	return &ConnectionMgr{
		list: make(map[string]iface.IConnection),
	}
}

func (mgr *ConnectionMgr) Add(conn iface.IConnection) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	mgr.list[conn.GetId()] = conn
	log.Println("ConnectionMgr Add ", conn.GetId())
}

// 通过 连接 id 进行删除
func (mgr *ConnectionMgr) Remove(conn iface.IConnection) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	delete(mgr.list, conn.GetId())
	log.Println("ConnectionMgr Remove ", conn.GetId())
}

// 获取一个连接
func (mgr *ConnectionMgr) Get(connId string) (iface.IConnection, error) {
	mgr.lock.RLocker()
	defer mgr.lock.RUnlock()

	if conn, ok := mgr.list[connId]; ok {
		return conn, nil
	}

	return nil, errors.New(fmt.Sprintln("ConnectionMgr.Get", connId, "not found"))
}

// 连接池大小
func (mgr *ConnectionMgr) Len() int {
	mgr.lock.RLocker()
	defer mgr.lock.RUnlock()

	len := len(mgr.list)
	return len
}
