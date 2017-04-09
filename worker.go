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
    taskArg := make([]reflect.Value, 1)
    reflectedTask := reflect.ValueOf(taskFunc)
    taskArg[0] = reflect.ValueOf(task)

    results := reflectedTask.Call(taskArg)
    log.Println("Func called successfully")
    if !results[1].IsNil() {
        // add to results queue as uncessful
        err := results[1].Interface().(error)
        task.FinishWithError(err)
        w.machine.SendTaskResult(task)
        return err 
    }
    // add to resuls queue as successful
    task.FinishWithSuccess()
    w.machine.SendTaskResult(task)
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

