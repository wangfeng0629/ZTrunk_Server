package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"reflect"
	"strings"
)

type Handlers struct {
}

func (h *Handlers) ResAction(w http.ResponseWriter, req *http.Request) {
	fmt.Println("res")
	w.Write([]byte("hcg"))
}

type RedisHander struct {
}

func say(w http.ResponseWriter, req *http.Request) {
	//fmt.Println(req)
	c := GetRedisHander()
	switch req.Method {
	case "GET":
		if c == nil {
			fmt.Println("redis connet failed")
		}
		v, e := c.Do("set", string(111111), req)
		if e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(v)
			fmt.Println(redis.String(c.Do("get", string(111111))))
		}
	}
	defer c.Close()
	return
	fmt.Println("Method:", req.Method)
	pathInfo := strings.Trim(req.URL.Path, "/")
	fmt.Println("pathInfo:", pathInfo)

	parts := strings.Split(pathInfo, "/")
	fmt.Println("parts:", parts)

	var action = "ResAction"
	fmt.Println(strings.Join(parts, "|"))
	if len(parts) > 1 {
		fmt.Println("22222222")
		action = strings.Title(parts[1]) + "Action"
	}
	fmt.Println("action:", action)
	handle := &Handlers{}
	controller := reflect.ValueOf(handle)
	method := controller.MethodByName(action)
	r := reflect.ValueOf(req)
	wr := reflect.ValueOf(w)
	method.Call([]reflect.Value{wr, r})
}

func GetRedisHander() redis.Conn {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("test error")
		return nil
	}
	return c
}

func main() {
	http.HandleFunc("/", say)
	http.Handle("/hcg/", http.HandlerFunc(say))
	http.ListenAndServe(":8001", nil)

	//select {} //阻塞进程
}
