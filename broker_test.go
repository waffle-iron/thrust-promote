package main

import (
    "encoding/json"
    "time"
    "testing"
    "github.com/stretchr/testify/assert"
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

    task := NewTask("test", string(metadata))

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

func TestBrokerSendToWorkers(t *testing.T) {
    cfg := config.LoadConfig("config/config.yaml")
    redisCfg := cfg.Redis.Development
    assert.Equal(t, redisCfg.Url, "redis://localhost:6379/0")
    redisCfg.ParseUrl()

    // test instantiation
    machine := NewMachine(&redisCfg)

    taskMap := map[string]interface{}{
        "printMessage": PrintMessage}

    // register task
    machine.RegisterTasks(taskMap)
    // test queueing task in db
    payload := TestPayload{
        Message: "hello!",
    }

    metadata, err := json.Marshal(payload)

    task := NewTask("printMessage", string(metadata))
    badTask := NewTask("test", string(metadata))

    err = machine.SendTask(task)
    machine.broker.errorsChan = make(chan error)
    if assert.NoError(t, err) {
        w := NewWorker("worker-test", machine)
        tasks := make(chan *Task)
        go func() {
            tasks <- task
            tasks <- badTask
        }()
        err = machine.broker.SendToWorkers(tasks, w)
        assert.Error(t, err, "badTask should raise error")
    }

}

func TestBrokerDequeue(t *testing.T) {
    cfg := config.LoadConfig("config/config.yaml")
    redisCfg := cfg.Redis.Development
    assert.Equal(t, redisCfg.Url, "redis://localhost:6379/0")
    redisCfg.ParseUrl()

    // test instantiation
    machine := NewMachine(&redisCfg)

    taskMap := map[string]interface{}{
        "printMessage": PrintMessage}

    // register task
    machine.RegisterTasks(taskMap)
    // test queueing task in db
    payload := TestPayload{
        Message: "hello!",
    }

    metadata, err := json.Marshal(payload)

    task := NewTask("printMessage", string(metadata))
    // badTask := NewTask("test", string(metadata))

    err = machine.SendTask(task)
    assert.NoError(t, err)
    // err = machine.SendTask(badTask)
    // assert.NoError(t, err)
    machine.broker.errorsChan = make(chan error)
    w := NewWorker("worker-test", machine)
    time.AfterFunc(1*time.Second, machine.broker.StopDequeue)
    retry, err := machine.broker.Dequeue(w)
    assert.NoError(t, err, "stopChan does not raise an error")
    assert.Equal(t, retry, false)

}