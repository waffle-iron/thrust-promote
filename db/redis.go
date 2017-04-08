package db

import (
    "time"
    "github.com/garyburd/redigo/redis"
)

func CreatePool(addrStr string) *redis.Pool {
    return &redis.Pool{
        MaxIdle: 2,
        IdleTimeout: 60 * time.Second,
        Dial: func() (redis.Conn, err) {
            return redis.Dial("tcp", addrStr)
        }
    }
}