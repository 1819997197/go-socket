package impl

import "go-socket/ch12/iface"

type Message struct {
	Id      uint32 //消息id
	DataLen uint32 //消息长度
	Data    []byte //消息内容
}

func NewMessage(id uint32, data []byte) iface.IMessage {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message) GetMsgID() uint32 {
	return m.Id
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}

func (m *Message) SetDataLen(len uint32) {
	m.DataLen = len
}
