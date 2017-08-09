package db

import (
    "fmt"
    "log"
    "database/sql"
    config "github.com/ammoses89/thrust-promote/config"
    _ "github.com/lib/pq"
)

type Postgres struct {
    cfg *config.ConnectionSettings
}

func NewPostgres(cfg *config.ConnectionSettings) *Postgres {
    if cfg.Url == "" {
        cfg.Url = fmt.Sprintf("postgres://%s:%d/%s?sslmode=disable", cfg.Host, cfg.Port, cfg.Database)
    }
    return &Postgres{cfg: cfg}
}

func (pg *Postgres) GetConn() (*sql.DB, error) {
    db, err := sql.Open("postgres", pg.cfg.Url)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    return db, nil
}

func (pg *Postgres) IsNoResultsErr(err error) bool {
    if err == sql.ErrNoRows {
        return true
    }
    return false
}