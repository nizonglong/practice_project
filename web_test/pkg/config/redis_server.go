package config

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

// RedisServer 表示 redis 服务器配置
type RedisServer struct {
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

// NewPool 初始化redis连接
func (c RedisServer) NewPool(maxIdle int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", c.Addr, redis.DialDatabase(c.DB), redis.DialPassword(c.Password))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
