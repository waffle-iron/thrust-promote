package main

import (
    "fmt"
    "time"
    "encoding/json"
    "github.com/RichardKnop/uuid"
)

type Task struct {
    Id string
    StartTimestamp time.Time
    EndTimestamp time.Time
    Retries int
    Status string
    Name string
    Metadata string
    ErrorMessage string
}

func NewTask(name string, metdata string) *Task {
    // add UUID
    return &Task{
        Id:       fmt.Sprintf("task-%s-%v", name, uuid.New()),
        Status:   "Queued",
        StartTimestamp: time.Now(),
        Name:     n,
        Metadata: metadata}
}

func (t *Task) DeserializeMetadata(obj interface{}) error {
    // deserialize metadata into obj
    if err := json.Unmarshal([]byte(t.Metadata), &obj); err != nil {
        return err
    }
    return nil
}

func (t *Task) FinishWithError(err error) {
    t.EndTimestamp = time.Now()
    t.Status = "FAILED"
    t.ErrorMessage = err.Error()
}

func (t *Task) FinishWithError(err error) {
    t.EndTimestamp = time.Now()
    t.Status = "FAILED"
}