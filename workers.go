package main

import (
    "log"
    "reflect"
)

type Worker struct {
    Name string
    machine *Machine
}

func (w *Worker) Process(task Task) int {
    // will take a task and run task with args
    taskFunc := w.machine.GetRegisteredTask(task.Name)
    metadataArg := make([]reflect.Value, 1)
    reflectedTask := reflect.ValueOf(taskFunc)
    if err := reflectedTask.Call(metadataArg); err != nil {
        log.Fatalf("Error occured with running task: %v", err)
        panic(err)
    }
}


func (w *Worker) Run() error {
    broker := w.machine.GetBroker()
    errChan := make(chan error)

    go func() {
        for {
            err := broker.Dequeue(w)
            if err {
                errChan <- err
                return
            }
        }
    }()

    return <-errChan

}

