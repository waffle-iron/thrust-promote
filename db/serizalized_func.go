package main

import (
    "fmt"
    "reflect"
    "encoding/json"
)

type Task struct {
    Name string
}

func Hello() string {
    return "Hello!"
}

func main() {
    taskMap := make(map[string]interface{})
    taskMap["hello"] = Hello
    task := Task{Name: "hello"}

    emptyArgs := make([]reflect.Value, 0)
    reflectedTask := reflect.ValueOf(taskMap[task.Name])
    fmt.Println(reflectedTask.Call(emptyArgs))

    serialized, err := json.Marshal(task)

    if err != nil {
        panic(err)
    }

    fmt.Println("serialized data: ", string(serialized))

    var deseralized Task

    if err := json.Unmarshal(serialized, &deseralized); err != nil {
        panic(err);
    }
    reflectedTask = reflect.ValueOf(taskMap[deseralized.Name])
    fmt.Println(reflectedTask.Call(emptyArgs))
}

