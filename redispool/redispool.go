package redispool

import (
	"ZTrunk_Server/logger"
	"ZTrunk_Server/setting"

	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// 连接池
type ConnPool struct {
	redisPool *redis.Pool
}

var pool = &ConnPool{}

// 初始化Redis连接池
func InitRedisPool() bool {
	redisAddr := fmt.Sprintf("%s:%d", setting.RedisIP, setting.RedisPort)
	pool.redisPool = newPool(redisAddr, "", 0, setting.MaxOpenConn, setting.MaxIdleConn)
	if _, err := pool.DoCmd("PING"); err != nil {
		logger.Fatal("Init Redis Poll Failed !!! %v", err.Error())
		return false
	}
	logger.Info("初始化 Redis [%s] 成功", redisAddr)
	return true
}

// 删除连接池
func FreePool() error {
	err := pool.redisPool.Close()
	logger.Fatal("free redis poll")
	return err
}

// 新建连接池
func newPool(host, password string, database, maxOpenConns, maxIdleConns int) *redis.Pool {
	return &redis.Pool{
		MaxActive:   maxOpenConns,
		MaxIdle:     maxIdleConns,
		IdleTimeout: 120 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			if len(password) > 0 {
				if _, err := conn.Do("AUTH", password); err != nil {
					conn.Close()
					return nil, err
				}
			}
			if _, err := conn.Do("select", database); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		TestOnBorrow: func(conn redis.Conn, time time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}

// 获取池
func GetRedis() *ConnPool {
	return pool
}

// 执行指令
func (connPool *ConnPool) DoCmd(command string, args ...interface{}) (interface{}, error) {
	conn := connPool.redisPool.Get()
	defer conn.Close()
	return conn.Do(command, args...)
}

// 通过字符存值
func (connPool *ConnPool) SetByString(key string, value interface{}) (interface{}, error) {
	conn := connPool.redisPool.Get()
	defer conn.Close()
	return conn.Do("SET", key, value)
}

// 通过字符取值
func (connPool *ConnPool) GetByString(key string) (string, error) {
	conn := connPool.redisPool.Get()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}
