package impl

import (
	"bytes"
	"encoding/binary"
	"go-socket/ch13/iface"
)

type DataPack struct{}

func NewDataPack() iface.IDataPack {
	return &DataPack{}
}

func (d *DataPack) GetHeadLen() uint32 {
	// Id uint32(4字节) + DataLen uint32(4字节)
	return 8
}

// 封包
func (d *DataPack) Pack(msg iface.IMessage) ([]byte, error) {
	dataBuffer := bytes.NewBuffer([]byte{})

	// 写dataLen
	if err := binary.Write(dataBuffer, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	// 写msgId
	if err := binary.Write(dataBuffer, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}

	// 写data
	if err := binary.Write(dataBuffer, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuffer.Bytes(), nil
}

// 解包
func (d *DataPack) Unpack(binaryData []byte) (iface.IMessage, error) {
	dataBuffer := bytes.NewReader(binaryData)

	msg := &Message{}

	// 读dataLen
	if err := binary.Read(dataBuffer, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// 读msgId
	if err := binary.Read(dataBuffer, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	return msg, nil
}
