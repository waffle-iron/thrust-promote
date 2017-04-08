package db

import (
    "database/sql"
    _ "github.com/lib/pq"
)

type Postgres struct {}

func (pg *Postgres) GetConn() *sql.DB {
    db, err := sql.Open("postgres", "connection url")
    if err != nil {
        log.Fatal(err)
    }

    return db
}