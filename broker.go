package main

import (
    "encoding/json"
    DB "github.com/ammoses89/thrust-workers/db"
    "github.com/garyburd/redigo/redis"
)

type Broker struct {
    host string
    db string
    password string
    pool *redis.Pool
    errorChan chan error
}

func (broker *Broker) Dequeue(worker Worker) error {
    broker.pool = DB.CreatePool(broker.host)
    defer broker.pool.Close()

    broker.errorChan = make(chan error)
    tasks := make(chan []Task)
    conn = broker.pool.Get()

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

                if err := json.Unmarshal(item, &Task); err != nil {
                    broker.errorChan <- err
                    return
                }

                tasks <- Task

            }
        }
        
    }()

    if err := broker.SendToWorkers(tasks, worker); err != nil {
        broker.errorChan <- err
        return
    }
}

func (broker *Broker) SendToWorkers(tasks <-chan Task, worker Worker) error{

    for {
        select {
        case t := <- tasks:
            go func () {
                if err := Worker.Process(&t); err != nil {
                    broker.errorChan <- err;
                    return
                } 
            }()
        }
    }

    return <-broker.errorChan
}