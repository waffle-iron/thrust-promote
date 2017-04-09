package main

import (
    "fmt"
    "testing"
    "encoding/json"
    "github.com/stretchr/testify/assert"
    "github.com/RichardKnop/uuid"
)

func TestTask(t *testing.T) {

    payload := TestPayload{
        Message: "hello!",
    }

    metadata, err := json.Marshal(payload)

    task := &Task{
        Id:       fmt.Sprintf("task-%v", uuid.New()),
        Status:   "Queued",
        Name:     "test",
        Metadata: string(metadata)}

    var newPayload TestPayload
    err = task.DeserializeMetadata(&newPayload)
    if assert.NoError(t, err) {
        assert.Equal(t, newPayload.Message, payload.Message, "Successful deserialization")
    }
}