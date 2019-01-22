package network

import (
	"ZTrunk_Server/logger"

	"net"
)

// 连接接受器
type Acceptor struct {
	// 侦听地址
	addr string

	// 侦听器
	listener net.Listener

	// 连接管理器
	TCPConnManager
}

func (a *Acceptor) Start(addr string) {
	a.addr = addr

	// 开始侦听
	go a.listen(addr)
}

func (a *Acceptor) listen(addr string) {
	var err error

	// 侦听器在给定地址侦听
	a.listener, err = net.Listen("tcp", addr)
	if err != nil {
		logger.Fatal("%s", err.Error())
		return
	}

	logger.Info("侦听 [%s]", a.addr)

	// 侦听循环，接受连接
	for {
		conn, err := a.listener.Accept()

		// 侦听错误，跳出循环
		if err != nil {
			// TODO 或者持续侦听？
			break
		}

		// 处理到来的连接
		go a.onNewConnection(conn)
	}
}

func (a *Acceptor) onNewConnection(conn net.Conn) {
	tcpConnection := newTCPConnection(conn)

	tcpConnection.Run()

	// 加入管理器中
	a.Add(tcpConnection)
}

// 新建一个接受器
func NewAcceptor() *Acceptor {
	acceptor := new(Acceptor)
	return acceptor
}
