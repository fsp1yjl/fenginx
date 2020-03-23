package face

type IMsgPack interface {
	UnPack([]byte) (IMessage, error)
	Pack(msg IMessage) ([]byte, error)
	HeaderLen() uint32
}
