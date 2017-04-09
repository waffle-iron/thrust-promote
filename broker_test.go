package main

import (
    "fmt"
    "encoding/json"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/RichardKnop/uuid"
    config "github.com/ammoses89/thrust-workers/config"
    db "github.com/ammoses89/thrust-workers/db"
)

func TestBroker(t *testing.T) {
    cfg := config.LoadConfig("config/config.yaml")
    redisCfg := cfg.Redis.Development
    assert.Equal(t, redisCfg.Url, "redis://localhost:6379/0")
    redisCfg.ParseUrl()

    // test instantiation
    broker := NewBroker(&redisCfg)
    assert.Equal(t, broker.host, "localhost")
    assert.Equal(t, broker.database, "0")
    assert.Equal(t, broker.port, 6379)
    assert.Equal(t, broker.user, "")
    assert.Equal(t, broker.password, "")

    // test queueing task in db
    payload := TestPayload{
        Message: "hello!",
    }

    metadata, err := json.Marshal(payload)

    task := &Task{
        Id:       fmt.Sprintf("task-%v", uuid.New()),
        Status:   "Queued",
        Name:     "test",
        Metadata: string(metadata)}

    err = broker.QueueTask(task)
    if assert.NoError(t, err) {
        addString := broker.BuildAddrString()
        broker.pool = db.CreatePool(addString)
        defer broker.pool.Close()

        conn := broker.pool.Get()
        _, err := conn.Do("BLPOP", "thrust-default", "1")
        assert.NoError(t, err)
    }

}