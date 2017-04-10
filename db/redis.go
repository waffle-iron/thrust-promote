package db

import (
    "time"
    "strings"
    "github.com/garyburd/redigo/redis"
)

func CreatePool(addrStr string) *redis.Pool {
    db := "0"
    // check if a database was provided
    dbSplit := strings.Split(addrStr, "/")
    if len(dbSplit) > 1 {
        // set the addr string without the db
        addrStr = dbSplit[0]
        // set the db
        db = dbSplit[len(dbSplit)-1]
    }
    return &redis.Pool{
        MaxIdle: 2,
        IdleTimeout: 200 * time.Second,
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", addrStr)
            if err != nil {
                return nil, err
            }
            // select the db for the connection pool
            if _, err := c.Do("SELECT", db); err != nil {
                c.Close()
                return nil, err
            }
            return c, nil
        }}
}

func CreateConn(addrStr string) (redis.Conn, error) {
    db := "0"
    // check if a database was provided
    dbSplit := strings.Split(addrStr, "/")
    if len(dbSplit) > 1 {
        // set the addr string without the db
        addrStr = dbSplit[0]
        // set the db
        db = dbSplit[len(dbSplit)-1]
    }
    c, err := redis.Dial("tcp", addrStr)
    if err != nil {
        return nil, err
    }
    // select the db for the connection pool
    if _, err := c.Do("SELECT", db); err != nil {
        c.Close()
        return nil, err
    }

    return c, nil
}

