package network

type TCPServer struct {
	// 服务器名字
	name string

	// 服务器地址
	addr string
}

// 启动
func (tcpServer *TCPServer) Start() {
	tcpServer.init(tcpServer.addr)
}

// 初始化
func (tcpServer *TCPServer) init(addr string) {
	acceptor := NewAcceptor()
	acceptor.Start(addr)
}

// 新建服务器TCPServer实例
func newTCPServer(name, addr string) *TCPServer {
	tcpServer := new(TCPServer)
	tcpServer.name = name
	tcpServer.addr = addr
	return tcpServer
}
