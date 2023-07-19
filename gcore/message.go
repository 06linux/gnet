package gcore

type Message struct {

	// 消息 id (唯一)
	id string

	// 事件 id
	eventId string

	// 房间 id (如果 room id 存在则发送到指定房间)
	roomId string

	// 消息内容
	data []byte
}

func NewMessage() *Message {
	return &Message{
		id: "xxxxx",
	}
}

func (msg *Message) Pack() ([]byte, error) {

	return nil, nil
}

func (msg *Message) Unpack([]byte) error {

	return nil
}
