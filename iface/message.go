package iface

type IMessage interface {
	Pack() ([]byte, error) // 封包
	Unpack([]byte) error   // 拆包
}
