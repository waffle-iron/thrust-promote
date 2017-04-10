package main

import (
    "fmt"
    "log"
    "time"
    "reflect"
    "encoding/json"
    "testing"
    "github.com/stretchr/testify/assert"
    config "github.com/ammoses89/thrust-workers/config"
)

func PrintMessage(task *Task) (bool, error) {
    var payload TestPayload
    err := task.DeserializeMetadata(&payload)
    if err != nil {
        log.Fatalf("Failed to deserialize payload: %v", err)
        return false, nil
    }
    fmt.Println(payload.Message)
    return true, nil
}

func TestMachine(t *testing.T) {
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

    fn := machine.GetRegisteredTask("printMessage")

    // test queueing task in db
    payload := TestPayload{
        Message: "hello!",
    }

    metadata, err := json.Marshal(payload)

    task := NewTask("test", string(metadata))
    taskArg := []reflect.Value{reflect.ValueOf(task)}
    reflectedTask := reflect.ValueOf(fn)
    results := reflectedTask.Call(taskArg)
    ok := results[0].Interface().(bool)
    err = nil
    if !results[1].IsNil() {
        err = results[1].Interface().(error)
    }

    if assert.NoError(t, err) {
        expectedOk, _ := PrintMessage(task) 
        assert.Equal(t, ok, expectedOk)
    }

    err = machine.SendTask(task)
    assert.NoError(t, err, "Task sent successfully")
}


func TestMachineSendTaskResult(t *testing.T) {

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

    task := NewTask("test", string(metadata))
    err = machine.SendTaskResult(task)
    assert.NoError(t, err, "Task sent successfully")
}

func TestMachineLaunchWorkers(t *testing.T) {
    cfg := config.LoadConfig("config/config.yaml")
    redisCfg := cfg.Redis.Development
    assert.Equal(t, redisCfg.Url, "redis://localhost:6379/0")
    redisCfg.ParseUrl()

    // test instantiation
    machine := NewMachine(&redisCfg)
    machine.broker.DeleteQueue()

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
    err = machine.SendTask(task)
    assert.NoError(t, err)
    err = machine.SendTask(task)
    assert.NoError(t, err)
    err = machine.SendTask(task)
    assert.NoError(t, err)
    time.AfterFunc(3*time.Second, machine.broker.StopDequeue)
    err = machine.LaunchWorkers(2)
    assert.NoError(t, err, "stopChan does not raise an error")
    machine.broker.DeleteQueue()
}