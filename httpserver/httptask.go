package main

import (
	"ZTrunk_Server/logger"
	"ZTrunk_Server/setting"
	"fmt"
	"net/http"
	"time"
)

var (
	server = &http.Server{
		Addr:           ":http",
		Handler:        &ppserver{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	handlersMap = make(map[string]HandlersFunc)
)

type ppserver struct {
}

type HandlersFunc func(http.ResponseWriter, *http.Request)

func (*ppserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		logger.Error("客户端请求错误【%s】，只能请求Post", r.Method)
		w.Write([]byte("错误请求模式"))
		return
	}
	setHeader(w)
	if h, ok := handlersMap[r.URL.Path]; ok {
		h(w, r)
	}
}

// 设置访问域
func setHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
}

// 初始化网页链接列表（http消息列表） 后期要搞成配置 客户端、服务器公用
func InitHandler() {
	server.Addr = fmt.Sprintf("%s:%d", setting.HTTPIp, setting.HTTPPort)
	handlersMap["/get"] = HttpGetTask
	handlersMap["/set"] = HttpSetTask
}

// 启动http服务
func HttpStartServer() bool {
	InitHandler()
	logger.Info("[启动] Http监听端口 [%d]", setting.HTTPPort)
	server.ListenAndServe()
	return true
}
