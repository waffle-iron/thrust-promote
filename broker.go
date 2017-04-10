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
    retry bool
    pool *redis.Pool
    stopChan chan int
    errorsChan chan error
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

    conn := broker.pool.Get()

    serializedTask, err := json.Marshal(task)
    if err != nil {
        return err
    }

    _, err = conn.Do("RPUSH", "thrust-result-default", serializedTask)
    return err
}

func (broker *Broker) Dequeue(worker *Worker) (bool, error) {
    addString := broker.BuildAddrString()
    broker.pool = db.CreatePool(addString)
    defer broker.pool.Close()

    broker.errorsChan = make(chan error)
    broker.stopChan = make(chan int)
    tasks := make(chan *Task)
    conn := broker.pool.Get()

    broker.wg.Add(1)

    go func() {
        defer broker.wg.Done()

        for {
            select {
            case <-broker.stopChan:
                return 
            default:
                itemBytes, err := conn.Do("BLPOP", "thrust-default", "1")

                if err != nil {
                    fmt.Printf("Error occured: %v \n", err)
                    broker.errorsChan <- err
                    return
                }

                if itemBytes == nil {
                    fmt.Println("NIL Bytes now continue")
                    continue
                }

                items, err := redis.ByteSlices(itemBytes, nil)
                if err != nil {
                    broker.errorsChan <- err
                    return
                }

                item := items[1]
                var task Task
                if err := json.Unmarshal(item, &task); err != nil {
                    broker.errorsChan <- err
                    return
                }

                tasks <- &task
            }
        }
        
    }()

    if err := broker.SendToWorkers(tasks, worker); err != nil {
        return broker.retry, err
    }

    if broker.retry == false {
        fmt.Printf("Worker %s not retrying\n", worker.Name)
    }

    return broker.retry, nil
}

func (broker *Broker) SendToWorkers(tasks <-chan *Task, worker *Worker) error{

    for {
        select {
        case err := <-broker.errorsChan:
            return err
        case t := <- tasks:
            fmt.Println("Task received ", t.Name)
            go func () {
                if err := worker.Process(t); err != nil {
                    fmt.Printf("Error occured: %v \n", err)
                    broker.errorsChan <- err;
                    return
                } 
            }()
        case <-broker.stopChan:
            return nil
        }
    }
}

func (broker *Broker) StopDequeue() {
    fmt.Println("Stop chan triggered")
    broker.retry = false
    broker.stopChan <- 1
    broker.wg.Wait()
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

func (broker *Broker) DeleteQueue() error {
    addString := broker.BuildAddrString()
    broker.pool = db.CreatePool(addString)
    defer broker.pool.Close()

    conn := broker.pool.Get()
    _, err := conn.Do("DEL", "thrust-default")
    if err != nil {
        return err
    }
    return nil
}