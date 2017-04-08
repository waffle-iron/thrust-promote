package db

import (
    "time"
    "github.com/garyburd/redigo/redis"
)

func CreatePool(addrStr string) *redis.Pool {
    return &redis.Pool{
        MaxIdle: 2,
        IdleTimeout: 200 * time.Second,
        Dial: func() (redis.Conn, error) {
            return redis.Dial("tcp", addrStr)
        }}
}