package main

import (
    "errors"
    "testing"
    "encoding/json"
    "github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {

    payload := TestPayload{
        Message: "hello!",
    }

    metadata, err := json.Marshal(payload)
    task := NewTask("test", string(metadata))

    var newPayload TestPayload
    err = task.DeserializeMetadata(&newPayload)
    if assert.NoError(t, err) {
        assert.Equal(t, newPayload.Message, payload.Message, "Successful deserialization")
    }
}

func TestTaskFinish(t *testing.T) {

    payload := TestPayload{
        Message: "hello!",
    }
    metadata, err := json.Marshal(payload)

    task := NewTask("test", string(metadata))

    task.FinishWithSuccess()
    assert.NotEqual(t, task.EndTimestamp, nil)
    assert.Equal(t, task.Status, "SUCCEEDED")

    err = errors.New("Test Error")

    task.FinishWithError(err)
    assert.NotEqual(t, task.EndTimestamp, nil)
    assert.Equal(t, task.Status, "FAILED")
    assert.Equal(t, task.ErrorMessage, "Test Error")
}