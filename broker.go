package main

import (
    "encoding/json"
    "github.com/garyburd/redigo/redis"
    db "github.com/ammoses89/thrust-workers/db"
)

type Broker struct {
    host string
    database string
    password string
    port int
    pool *redis.Pool
    errorChan chan error
}

func (broker *Broker) QueueTask(task *Task) error {
    broker.pool = db.CreatePool(broker.host)
    defer broker.pool.Close()

    broker.errorChan = make(chan error)
    conn := broker.pool.Get()

    serializedTask, err := json.Marshal(task)
    if err != nil {
        panic(err)
    }

    _, err = conn.Do("RPUSH", "default", serializedTask)
    return err
}

func (broker *Broker) Dequeue(worker *Worker) error {
    broker.pool = db.CreatePool(broker.host)
    defer broker.pool.Close()

    broker.errorChan = make(chan error)
    tasks := make(chan Task)
    conn := broker.pool.Get()

    go func() {
        for {
            select {
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

                tasks <- task

            }
        }
        
    }()

    if err := broker.SendToWorkers(tasks, worker); err != nil {
        broker.errorChan <- err
    }

    return <-broker.errorChan
}

func (broker *Broker) SendToWorkers(tasks <-chan Task, worker *Worker) error{

    for {
        select {
        case t := <- tasks:
            go func () {
                if err := worker.Process(&t); err != nil {
                    broker.errorChan <- err;
                    return
                } 
            }()
        }
    }

    return <-broker.errorChan
}