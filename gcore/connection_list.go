package gcore

import (
	"errors"
	"fmt"
	"gnet/gface"
	"log"
	"sync"
)

// 连接池
type ConnectionList struct {

	// 所有链接
	list map[string]gface.IConnection

	// 读写锁
	lock sync.RWMutex
}

func NewConnectionList() *ConnectionList {
	return &ConnectionList{
		list: make(map[string]gface.IConnection),
	}
}

func (pool *ConnectionList) Add(conn gface.IConnection) {

	pool.lock.Lock()
	defer pool.lock.Unlock()

	pool.list[conn.GetId()] = conn
	log.Println("ConnectionList Add ", conn.GetId())
}

// 通过 链接 id 进行删除
func (pool *ConnectionList) Remove(connId string) {

	pool.lock.Lock()
	defer pool.lock.Unlock()

	delete(pool.list, connId)
	log.Println("ConnectionList Remove ", connId)
}

// 获取一个链接
func (pool *ConnectionList) GetConn(connId string) (gface.IConnection, error) {

	pool.lock.RLocker()
	defer pool.lock.RUnlock()

	if conn, ok := pool.list[connId]; ok {
		return conn, nil
	}

	return nil, errors.New(fmt.Sprintln("ConnectionList.GetConn", connId, "not found"))
}

// 连接池大小
func (pool *ConnectionList) Len(connId string) int {
	pool.lock.RLocker()
	defer pool.lock.RUnlock()

	len := len(pool.list)

	return len
}
