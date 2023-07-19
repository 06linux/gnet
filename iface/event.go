package iface

type EventId string
type EventHandle func(IMessage) error

type IEvent interface {

	// 注册事件
	On(EventId, EventHandle)

	// 关闭事件
	Off(EventId, EventHandle)

	// 发送事件数据
	Emit(EventId, IMessage)
}
