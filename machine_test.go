package main

import (
    "fmt"
    "log"
    "reflect"
    "encoding/json"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/RichardKnop/uuid"
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

    val := machine.GetRegisteredTask("printMessage")

    // test queueing task in db
    payload := TestPayload{
        Message: "hello!",
    }

    metadata, err := json.Marshal(payload)

    task := NewTask("test", string(metadata))
    taskArg := make([]reflect.Value, 1)
    taskArg[0] = reflect.ValueOf(task)
    reflectedTask := reflect.ValueOf(val)
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