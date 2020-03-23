package fnet

type Message struct {
	Data []byte // 消息内容
	ID   uint32 // 消息类型编号
	Len  uint32 // 消息内容长度
}

func (c Message) GetData() []byte {
	return c.Data
}

func (c Message) GetID() uint32 {
	return c.ID
}

func (c Message) GetLength() uint32 {
	return c.Len
}

func (c *Message) SetData(data []byte) {
	c.Data = data
}

func (c *Message) SetID(id uint32) {
	c.ID = id
}

func (c *Message) SetLength(len uint32) {
	c.Len = len
}
