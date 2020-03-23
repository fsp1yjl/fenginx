package face

type IMessage interface {
	GetData() []byte
	GetID() uint32
	GetLength() uint32

	SetData([]byte)
	SetID(uint32)
	SetLength(uint32)
}
