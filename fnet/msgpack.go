package fnet

import (
	"bytes"
	"encoding/binary"
	face "fenginx/finterface"
	"fmt"
)

type MsgPack struct {
}

// 拆包，获取包头信息，更新到msg对象，此时未给msg.data赋值
func (p *MsgPack) UnPack(byteStream []byte) (face.IMessage, error) {

	r := bytes.NewReader(byteStream)

	msg := &Message{}
	if err := binary.Read(r, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &msg.Len); err != nil {
		return nil, err
	}

	//TODO:	判断dataLen的长度是否超出我们允许的最大包长度
	fmt.Println("pack msg:", string(msg.Data))
	return msg, nil
}

func (p *MsgPack) Pack(msg face.IMessage) ([]byte, error) {

	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//写msgID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetID()); err != nil {
		return nil, err
	}

	//写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetLength()); err != nil {
		return nil, err
	}

	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

//获取msg字节流包头字节数
func (p MsgPack) HeaderLen() uint32 {
	//这里设计包头为8字节， 其中msg id 4字节，msg data length 4字节
	return 8
}
