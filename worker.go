package main

import (
    "fmt"
    "errors"
    "log"
    "reflect"
)

type Worker struct {
    Name string
    machine *Machine
}

func NewWorker(name string, mach *Machine) *Worker {
    return &Worker{Name: name, machine: mach}
}

func (w *Worker) Process(task *Task) error {
    // will take a task and run task with args
    taskFunc := w.machine.GetRegisteredTask(task.Name)
    if taskFunc == nil {
        errMsg := fmt.Sprintf("Invalid Task Name: %v", task.Name)
        return errors.New(errMsg)
    }
    reflectedTask := reflect.ValueOf(taskFunc)

    taskArg := []reflect.Value{reflect.ValueOf(task)}
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
            retry, err := broker.Dequeue(w)
            if retry {
                fmt.Println("Retrying...")
            } else {
                errChan <- err
                return
            }
        }
    }()

    return <-errChan
}

func (w *Worker) Stop() {
    broker := w.machine.GetBroker()
    broker.StopDequeue()
}

