package main

import (
    "log"
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

func (mach *Machine) LaunchWorker(name string) error {
    w := Worker{Name: name}
    if err := w.Run(); err != nil {
        log.Fatalf("Worker failed to start: %v", err)
        panic(err)
    }
}

func (mach *Machine) LaunchWorkers() error {
    for i := range WORKER_COUNT {
        w := Worker{Name: "worker-" + i}
        if err := w.Run(); err != nil {
            log.Fatalf("Worker failed to start: %v", err)
            panic(err)
        }
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
    taskMap := map[string]interface{}{
        "transcode_audio": TranscodeAudio,
        "transcode_video": TranscodeVideo,
        "social_send": SocialSend,
        "event_send": EventSend,
        "release_send": ReleaseSend,
    }
    machine := Machine{}
    log.Info("Registering Tasks...")
    machine.RegisterTasks(taskMap)
    log.Info("Launching Workers...")
    if err := machine.LaunchWorkers(); err != nil {
        log.Fatalf("Failed to launch workers: %v", err)
        panic(err)
    }
}