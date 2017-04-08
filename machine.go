package main

import (
    "fmt"
)

const WORKER_COUNT = 5;

type Machine struct {
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

func (mach *Machine) LaunchWorker(name string) {
    w := Worker{name: name}
    if err := w.Run(); err != nil {

    }
}

func (mach *Machine) LaunchWorkers() error {
    errorChan = make(chan error)
    for i := range WORKER_COUNT {
        go func() {
            w := Worker{name: "worker-" + i}
            if err := w.Run(); err != nil {
                errorChan <- err
                return
            }
        }()
    }
}

func (mach *Machine) CreateBroker() *Broker {

    return &Broker{
        host:"127.0.0.0.1:6379"
    }
}

func (mach *Machine) GetBroker() *Broker {
    broker := mach.CreateBroker()
    return broker
}


func main() {
    /**
    This will be the server
    1) it will register tasks at start
    2) it will then launch workers
    3) the workes will run go routines
       that will exhaust the redis queue
    **/
    // First register tasks
}