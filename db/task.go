package main

type Task struct {
    Name string
    Metadata string
    PayloadType string
}

func (t *Task) DeserializeMetadata(obj interface{}) {
    // deserialize metadata into obj
    if err := json.Unmarshal(t.Metadata, &obj); err != nil {
        panic(err)
    }
}