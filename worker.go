package main

import (
    "log"
    "reflect"
)

type Worker struct {
    Name string
    machine *Machine
}

func (w *Worker) Process(task *Task) error {
    // will take a task and run task with args
    taskFunc := w.machine.GetRegisteredTask(task.Name)
    metadataArg := make([]reflect.Value, 1)
    reflectedTask := reflect.ValueOf(taskFunc)

    results := reflectedTask.Call(metadataArg)
    log.Println("FUnc called successfully")
    if !results[1].IsNil() {
        return results[1].Interface().(error)
    }
    return nil
}


func (w *Worker) Run() error {
    broker := w.machine.GetBroker()
    errChan := make(chan error)

    go func() {
        for {
            err := broker.Dequeue(w)
            if err != nil {
                errChan <- err
                return
            }
        }
    }()

    return <-errChan

}

