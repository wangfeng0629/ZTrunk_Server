package network

import "net"

// 消息缓存大小
const msgBufferSize = 100

type TCPConnection struct {
	// TCP原始套接字
	conn net.Conn

	// 唯一ID
	id int64

	// 消息队列
	messageQueueChan chan []byte
}

func (tcpConn *TCPConnection) Read(b []byte) (int, error) {
	return tcpConn.conn.Read(b)
}

func (tcpConn *TCPConnection) WriteMsg([]byte) {

}

func (tcpConn *TCPConnection) ReadMsg() ([]byte, error) {

}

func (tcpConn *TCPConnection) Run() {
	go tcpConn.writeLoop()

	go tcpConn.readLoop()
}

func (tcpConn *TCPConnection) writeLoop() {
	for b := range tcpConn.messageQueueChan {
		if b == nil {
			break
		}
		_, err := tcpConn.conn.Write(b)
		if err != nil {
			break
		}
	}
	tcpConn.conn.Close()
}

func (tcpConn *TCPConnection) readLoop() {
}

func (tcpConn *TCPConnection) GetID() int64 {
	return tcpConn.id
}

func (tcpConn *TCPConnection) SetID(id int64) {
	tcpConn.id = id
}

func newTCPConnection(conn net.Conn) *TCPConnection {
	tcpConn := new(TCPConnection)
	tcpConn.conn = conn
	tcpConn.messageQueueChan = make(chan []byte, msgBufferSize)
	return tcpConn
}
