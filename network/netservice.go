package network

import (
	"ZTrunk_Server/setting"
	"fmt"
)

var netService *NetService

type NetService struct {
	// TCPServer实例指针
	tcpServer *TCPServer
}

// 初始化服务器
func InitNetService() {
	addr := fmt.Sprintf("%s:%d", setting.HTTPIp, setting.HTTPPort)
	netService.tcpServer = newTCPServer("Super", addr)
	netService.Start()
}

// 启动服务器
func (netService *NetService) Start() {
	netService.tcpServer.Start()
}
