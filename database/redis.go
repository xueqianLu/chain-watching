package database

import (
	"github.com/gomodule/redigo/redis"
)

type RedisContext struct {
	redis.Conn
}

func GetConnect() *RedisContext {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		// handle error
	}

	return &RedisContext{c}
}

func (c *RedisContext) Lock() {

}

func (c *RedisContext) Unlock() {

}

func (c *RedisContext) Close() {
	c.Conn.Close()
}
