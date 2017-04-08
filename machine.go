package main

import (
    "log"
)

type Machine struct {
    broker *Broker
    TaskMap map[string]interface{}
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
    w := Worker{Name: name}
    if err := w.Run(); err != nil {
        log.Fatalf("Worker failed to start: %v", err)
        panic(err)
    }
}

func (mach *Machine) LaunchWorkers(worker_count int) error {
    for i := range worker_count {
        w := Worker{Name: "worker-" + i}
        if err := w.Run(); err != nil {
            log.Fatalf("Worker failed to start: %v", err)
            panic(err)
        }
    }
}

func (mach *Machine) SendTask(task *Task) {
    broker = mach.GetBroker()
    broker.QueueTask(task)
}

func (mach *Machine) CreateBroker() {

    mach.broker := &Broker{
        host:"redis://127.0.0.0.1:6379"}
}

func (mach *Machine) GetBroker() *Broker {
    if mach.broker == nil {
        mach.broker := mach.CreateBroker()
    }
    return mach.broker
}
