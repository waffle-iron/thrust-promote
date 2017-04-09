package main

import (
    "encoding/json"
)

type Task struct {
    Id string
    Status string
    Name string
    Metadata string
}

func (t *Task) DeserializeMetadata(obj interface{}) error {
    // deserialize metadata into obj
    if err := json.Unmarshal([]byte(t.Metadata), &obj); err != nil {
        return err
    }
    return nil
}