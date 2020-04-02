package iface

type IMessage interface {
	GetDataLen() uint32 //获取消息数据段长度
	GetMsgID() uint32   //获取消息id
	GetData() []byte    //获取消息内容

	SetMsgId(uint32)
	SetData([]byte)
	SetDataLen(uint32)
}
