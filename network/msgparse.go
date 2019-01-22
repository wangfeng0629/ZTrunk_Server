package network

import (
	"encoding/binary"
	"io"
)

// 消息结构
//-----------------
// | head | data |
//-----------------
// 4字节的包头

// 消息解析
type MsgParse struct {
}

// 写消息
func (msgP *MsgParse) Write(conn *TCPConnection, cmdData []byte, cmdLen uint32) {
	// 消息长度
	//TODO 检测消息的长度
	byteBuffer := make([]byte, 4+cmdLen)

	// 写入消息头
	binary.LittleEndian.PutUint32(byteBuffer, 4)

	// 写入消息体
	copy(byteBuffer[4:], cmdData)

	// 写入套接字
	conn.WriteMsg(byteBuffer)
}

// 读消息
func (msgP *MsgParse) Read(tcpConn *TCPConnection) ([]byte, error) {
	// 4字节的缓冲，用于读取消息的长度
	cmdLenBuffer := make([]byte, 4)
	if _, err := io.ReadFull(tcpConn, cmdLenBuffer); err != nil {
		return nil, err
	}

	cmdLen := binary.LittleEndian.Uint32(cmdLenBuffer)
	cmdData := make([]byte, cmdLen)
	// 读消息体
	if _, err := io.ReadFull(tcpConn, cmdData); err != nil {
		return nil, err
	}

	return cmdData, nil
}

// 新建消息解析
func newMsgParse() *MsgParse {
	msgParse := new(MsgParse)
	return msgParse
}
