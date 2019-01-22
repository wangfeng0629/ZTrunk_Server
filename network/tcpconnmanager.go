package network

import (
	"sync"
	"sync/atomic"
)

type TCPConnManager struct {
	// 连接容器
	TCPConnMap sync.Map

	// 当前已生成连接的ID
	curTCPConnID int64

	// 连接数量
	curNum int64
}

func (tcpConnM *TCPConnManager) Add(tcpConn *TCPConnection) {
	// 原子操作自增当前连接ID
	id := atomic.AddInt64(&tcpConnM.curTCPConnID, 1)

	// 原子操作自增当前连接数量
	tcpConnM.curNum = atomic.AddInt64(&tcpConnM.curNum, 1)

	// 设置连接ID
	tcpConn.SetID(id)

	tcpConnM.TCPConnMap.Store(id, tcpConn)
}

func (tcpConnM *TCPConnManager) Remove(tcpConn *TCPConnection) {
	tcpConnM.TCPConnMap.Delete(tcpConn.GetID())

	tcpConnM.curNum = atomic.AddInt64(&tcpConnM.curNum, -1)
}

// 通过ID获得一个连接
func (tcpConnM *TCPConnManager) GetTCPConnByID(id int64) *TCPConnection {
	if v, ok := tcpConnM.TCPConnMap.Load(id); ok {
		return v.(*TCPConnection)
	}
	return nil
}
