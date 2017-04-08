package main

import (
    "log"
    "fmt"
)

type Machine struct {
    cfg ConnectionSettings
    broker *Broker
    TaskMap map[string]interface{}
}

func (mach *Machine) NewMachine(cfg ConnectionSettings) *Machine {
    machine := &Machine{cfg: ConnectionSettings}
    machine.CreateBroker()
    return &machine
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
    w := Worker{Name: name}
    if err := w.Run(); err != nil {
        log.Fatalf("Worker failed to start: %v", err)
        return err
    }
    return nil
}

func (mach *Machine) LaunchWorkers(worker_count int) error {
    for i := 1; i <= worker_count; i++ {
        w := Worker{Name: fmt.Sprintf("worker-%d", i)}
        if err := w.Run(); err != nil {
            log.Fatalf("Worker failed to start: %v", err)
            return err
        }
    }
    return nil
}

func (mach *Machine) SendTask(task *Task) {
    broker := mach.GetBroker()
    broker.QueueTask(task)
}

func (mach *Machine) CreateBroker() {
    mach.cfg.ParseUrl()
    mach.broker = &Broker{
        host: mach.cfg.Host,
        database: mach.cfg.Database,
        port: mach.cfg.Port}
}

func (mach *Machine) GetBroker() *Broker {
    if mach.broker == nil {
        mach.CreateBroker()
    }
    return mach.broker
}
