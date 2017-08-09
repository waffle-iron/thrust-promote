package main

import (
    "encoding/json"
    "testing"
    "github.com/stretchr/testify/assert"
    config "github.com/ammoses89/thrust-promote/config"
)

func TestWorkerProgress(t *testing.T) {
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

    // Launch worker
    w := NewWorker("worker-test", machine)
    err = w.Process(task)
    assert.NoError(t, err, "Worker should process without error")

}

func TestWorkerRun(t *testing.T) {
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
    assert.NoError(t, err)
    err = machine.SendTask(badTask)
    assert.NoError(t, err)
    machine.broker.errorsChan = make(chan error)
    w := NewWorker("worker-test", machine)
    // time.AfterFunc(1*time.Second, machine.broker.StopDequeue)
    err = w.Run()
    assert.Error(t, err, "bad task should throw error")
}