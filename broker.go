package main

import (
    "fmt"
    "sync"
    "encoding/json"
    "github.com/garyburd/redigo/redis"
    config "github.com/ammoses89/thrust-workers/config"
    db "github.com/ammoses89/thrust-workers/db"
)

type Broker struct {
    host string
    database string
    url string
    user string
    password string
    port int
    pool *redis.Pool
    stopChan chan int
    errorChan chan error
    wg sync.WaitGroup
}

func NewBroker(cfg *config.ConnectionSettings) *Broker {
    return &Broker{
        host: cfg.Host,
        url: cfg.Url,
        database: cfg.Database,
        user: cfg.User,
        password: cfg.Password,
        port: cfg.Port}
}

func (broker *Broker) BuildAddrString() string {
    //TODO add support for Authentication
    return fmt.Sprintf("%s:%d/%s", broker.host, broker.port, broker.database)
}

func (broker *Broker) QueueTask(task *Task) error {
    addString := broker.BuildAddrString()
    broker.pool = db.CreatePool(addString)
    defer broker.pool.Close()

    conn := broker.pool.Get()

    serializedTask, err := json.Marshal(task)
    if err != nil {
        return err
    }

    _, err = conn.Do("RPUSH", "thrust-default", serializedTask)
    return err
}


func (broker *Broker) QueueTaskResult(task *Task) error {
    addString := broker.BuildAddrString()
    broker.pool = db.CreatePool(addString)
    defer broker.pool.Close()

    broker.errorChan = make(chan error)
    conn := broker.pool.Get()

    serializedTask, err := json.Marshal(task)
    if err != nil {
        return err
    }

    _, err = conn.Do("RPUSH", "thrust-result-default", serializedTask)
    return err
}

func (broker *Broker) Dequeue(worker *Worker) error {
    broker.pool = db.CreatePool(broker.host)
    defer broker.pool.Close()

    broker.errorChan = make(chan error)
    tasks := make(chan *Task)
    conn := broker.pool.Get()

    go func() {
        for {
            select {
            case <-broker.stopChan:
                return 
            default:
                itemBytes, err := conn.Do("BLOP", "default", "1")
                if err != nil {
                    broker.errorChan <- err
                    return
                }

                items, err := redis.ByteSlices(itemBytes, nil)
                if err != nil {
                    broker.errorChan <- err
                    return
                }

                item := items[1]

                var task Task
                if err := json.Unmarshal(item, &task); err != nil {
                    broker.errorChan <- err
                    return
                }

                tasks <- &task

            }
        }
        
    }()

    if err := broker.SendToWorkers(tasks, worker); err != nil {
        broker.errorChan <- err
    }

    return <-broker.errorChan
}

func (broker *Broker) SendToWorkers(tasks <-chan *Task, worker *Worker) error{

    for {
        select {
        case err := <-broker.errorChan:
            return err
        case t := <- tasks:
            fmt.Println("Task received ", t.Name)
            go func () {
                if err := worker.Process(t); err != nil {
                    fmt.Printf("Error occured: %v \n", err)
                    broker.errorChan <- err;
                    return
                } 
            }()
        }
    }
}


func (broker *Broker) GetQueuedTasks() ([]*Task, error) {
    addString := broker.BuildAddrString()
    broker.pool = db.CreatePool(addString)
    defer broker.pool.Close()

    conn := broker.pool.Get()

    itemBytes, err := conn.Do("LRANGE", "thrust-default", 0, -1)

    items, err := redis.ByteSlices(itemBytes, nil)
    if err != nil {
        return nil, err
    }

    var tasks []*Task
    for _, value := range items {
        var task Task
        if err := json.Unmarshal(value, &task); err != nil {
            return nil, err
        }
        tasks = append(tasks, &task)
    }
    return tasks, nil
}

func (broker *Broker) GetFinishedTasks() ([]*Task, error) {
    addString := broker.BuildAddrString()
    broker.pool = db.CreatePool(addString)
    defer broker.pool.Close()

    conn := broker.pool.Get()

    itemBytes, err := conn.Do("LRANGE", "thrust-result-default", 0, -1)

    items, err := redis.ByteSlices(itemBytes, nil)
    if err != nil {
        return nil, err
    }

    var tasks []*Task
    for _, value := range items {
        var task Task
        if err := json.Unmarshal(value, &task); err != nil {
            return nil, err
        }
        tasks = append(tasks, &task)
    }
    return tasks, nil

}