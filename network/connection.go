package network

type Connection interface {
	// 接收消息
	ReadMsg() ([]byte, error)

	// 发送消息
	WriteMsg([]byte)
}
