package db

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestPoolAndConn(t *testing.T) {
    pool := CreatePool("localhost:6379/0")
    conn := pool.Get()
    defer conn.Close()

    _, err := conn.Do("PING")
    assert.NoError(t, err, "Connection Successful")
}


func TestConn(t *testing.T) {
    conn, err := CreateConn("localhost:6379/0")
    assert.NoError(t, err, "Connection Successful")
    defer conn.Close()

    _, err = conn.Do("PING")
    assert.NoError(t, err, "Connection Ping Successful")
}

