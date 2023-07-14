package gface

type IMessage interface {
	Pack() ([]byte, error)     // 封包方法
	Unpack() (IMessage, error) //拆包
}
