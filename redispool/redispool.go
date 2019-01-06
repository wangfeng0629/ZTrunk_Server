package redispool

import (
	"github.com/garyburd/redigo/redis"

	"log"
	"time"
)

// 连接池
type ConnPool struct {
	redisPool *redis.Pool
}

func InitRedisPool(host, password string, database, maxOpenConns, maxIdleConns int) *ConnPool {
	connPool := &ConnPool{}
	connPool.redisPool = newPool(host, password, database, maxOpenConns, maxIdleConns)
	if _, err := connPool.DoCmd("PING"); err != nil {
		log.Fatal("Init Redis Poll Failed !!!", err.Error())
	}
	return connPool
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

// 关闭连接池
func (connPool *ConnPool) Close() error {
	err := connPool.redisPool.Close()
	return err
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
