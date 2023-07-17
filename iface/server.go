package iface

type IServer interface {
	Start()
	Stop()
	Restart()

	GetConnMgr() IConnectionMgr
}
