package gface

type IServer interface {
	Start()
	Stop()
	Restart()
}