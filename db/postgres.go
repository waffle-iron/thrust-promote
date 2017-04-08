package db

import (
    "log"
    "database/sql"
    _ "github.com/lib/pq"
)

type ConnectionSettings struct {
        Pool int `yaml:pool`
        Url string `yaml:url`
        Timeout int `yaml:timeout`
        Host string `yaml:host`
        Port string `yaml:port`
        Database string `yaml:database`
        Password string `yaml:password`
}

type Postgres struct {
    cfg ConnectionSettings
}

func NewPostgres(config ConnectionSettings) *Postgres {
    if !config.Url {
        config.Url = fmt.Sprintf("postgres://%s:%s/%s", config.Host, config.Port, config.Database)
    }
    return &Postgres{cfg: config}
}

func (pg *Postgres) GetConn() *sql.DB {
    db, err := sql.Open("postgres", pg.cfg.Url)
    if err != nil {
        log.Fatal(err)
    }

    return db
}