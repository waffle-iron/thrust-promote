package main

type Task struct {
    Id string
    Status string
    Name string
    Metadata string
}

func (t *Task) DeserializeMetadata(obj interface{}) {
    // deserialize metadata into obj
    if err := json.Unmarshal(t.Metadata, &obj); err != nil {
        panic(err)
    }
}