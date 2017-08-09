package main

import (
    "log"
    config "github.com/ammoses89/thrust-promote/config"
    "fmt"
)

type Machine struct {
    cfg *config.ConnectionSettings
    broker *Broker
    TaskMap map[string]interface{}
}

func NewMachine(cfg *config.ConnectionSettings) *Machine {
    machine := &Machine{cfg: cfg}
    machine.TaskMap = make(map[string]interface{})
    machine.CreateBroker()
    return machine
}

func (mach *Machine) GetRegisteredTask(taskName string) interface{} {
    val, ok := mach.TaskMap[taskName]
    if !ok {
        return nil
    }
    return val
}

func (mach *Machine) RegisterTask(taskName string, taskFunc interface{}) {
    mach.TaskMap[taskName] = taskFunc
}

func (mach *Machine) RegisterTasks(taskMap map[string]interface{}) {
    for taskName, taskFunc := range taskMap {
        mach.RegisterTask(taskName, taskFunc)
    }
}

func (mach *Machine) LaunchWorker(name string) error {
    w := NewWorker(name, mach)
    if err := w.Run(); err != nil {
        log.Fatalf("Worker failed to start: %v", err)
        return err
    }
    return nil
}

func (mach *Machine) LaunchWorkers(worker_count int) error {
    workerErrorsChan := make(chan error)
    for i := 1; i <= worker_count; i++ {
        fmt.Println("worker ", i)
        workerName := fmt.Sprintf("worker-%d", i)
        log.Printf("Starting worker %s\n", workerName)
        w := NewWorker(workerName, mach)
        go func() {
            err := w.Run(); 
            if err != nil {
                log.Fatalf("Worker failed: %v", err)
            }
            workerErrorsChan <- err
            w.Stop()
            return
        }()
    }
    return <-workerErrorsChan
}

func (mach *Machine) SendTask(task *Task) error {
    broker := mach.GetBroker()
    return broker.QueueTask(task)
}

func (mach *Machine) SendTaskResult(task *Task) error {
    // A task result is a task and two other params
    // the start date
    // end date
    broker := mach.GetBroker()
    return broker.QueueTaskResult(task)
}

func (mach *Machine) CreateBroker() {
    mach.cfg.ParseUrl()
    mach.broker = NewBroker(mach.cfg)
}

func (mach *Machine) GetBroker() *Broker {
    if mach.broker == nil {
        mach.CreateBroker()
    }
    return mach.broker
}
