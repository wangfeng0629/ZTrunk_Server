package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gomodule/redigo/redis"

	"ZTrunk_Server/redispool"
	"ZTrunk_Server/setting"
)

type Handlers struct {
}

func (h *Handlers) ResAction(w http.ResponseWriter, req *http.Request) {
	fmt.Println("res")
	w.Write([]byte("hcg"))
}

// test
func HttpTestTask(w http.ResponseWriter, req *http.Request, c redis.Conn) {
	h_str := strings.Split(req.URL.RawQuery, "?")
	if len(h_str) == 1 {
		id_str := strings.Split(h_str[0], "=")
		if id_str[0] == "id" {
			HttpGetTask(w, req, c)
		} else if id_str[0] == "del" {
			HttpDeleteTask(w, req, c)
		}
	} else if len(h_str) == 2 {
		HttpPostTask(w, req, c)
	} else {
		fmt.Println("http error！")
		w.Write([]byte("http error！"))
	}
}

// get
func HttpGetTask(w http.ResponseWriter, req *http.Request, c redis.Conn) {
	h_str := strings.Split(req.URL.RawQuery, "?")
	id_str := strings.Split(h_str[0], "=")
	if len(id_str) != 2 {
		fmt.Println("http get error!")
		w.Write([]byte("http get error!"))
		return
	}
	id := id_str[1]
	g_str, e := redis.String(c.Do("get", id))
	if e != nil {
		fmt.Println(e)
		w.Write([]byte("get error, not this key！"))
		return
	}
	g_str1 := "get ok！" + g_str
	fmt.Println(g_str1)
	w.Write([]byte(g_str1))
}

// put/post
func HttpPostTask(w http.ResponseWriter, req *http.Request, c redis.Conn) {
	h_str := strings.Split(req.URL.RawQuery, "?")
	id_str := strings.Split(h_str[0], "=")
	v_str := strings.Split(h_str[1], "=")
	if len(id_str) != 2 || len(v_str) != 2 {
		fmt.Println("http gost/put error!")
		w.Write([]byte("http post/put error!"))
		return
	}
	_, e := c.Do("set", id_str[1], v_str[1])
	if e != nil {
		fmt.Println(e)
		p_str := "post error " + h_str[0] + "" + h_str[1]
		w.Write([]byte(p_str))
		return
	}
	p_str := "post ok! " + h_str[0] + " " + h_str[1]
	fmt.Println(p_str)
	w.Write([]byte(p_str))
}

func HttpPutTask(w http.ResponseWriter, req *http.Request, c redis.Conn) {
}

// delete
func HttpDeleteTask(w http.ResponseWriter, req *http.Request, c redis.Conn) {
	h_str := strings.Split(req.URL.RawQuery, "?")
	id_str := strings.Split(h_str[0], "=")
	if len(id_str) != 2 {
		fmt.Println("http delete error!")
		w.Write([]byte("http delete error!"))
		return
	}
	id := id_str[1]
	_, e := c.Do("del", id)
	if e != nil {
		fmt.Println(e)
		return
	}
	p_str := "delete ok! " + id
	fmt.Println(p_str)
	w.Write([]byte(p_str))
}

func RedisHttpTask(w http.ResponseWriter, req *http.Request) {
	ret, err := redispool.GetInstance().SetByString("Redis", "RedisPool")
	if err != nil {
		log.Fatalf("Redis set faial : %v", err)
	} else {
		fmt.Println(ret)
		retVal, _ := redispool.GetInstance().GetByString("Redis")
		fmt.Println(retVal)
	}
}

func say(w http.ResponseWriter, req *http.Request) {
	c := GetRedisHander()
	switch req.Method {
	case "GET":
		HttpTestTask(w, req, c)
		//HttpGetTask(w, req, c)
		RedisHttpTask(w, req)
	case "POST":
		HttpPostTask(w, req, c)
	case "PUT":
		HttpPutTask(w, req, c)
	case "DELETE":
		HttpDeleteTask(w, req, c)
	}
	defer c.Close()
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

	httpAddr := fmt.Sprintf("%s:%d", setting.HTTPIp, setting.HTTPPort)
	fmt.Printf("serving on %s:%d \n", setting.HTTPIp, setting.HTTPPort)
	http.ListenAndServe(httpAddr, nil)
}
