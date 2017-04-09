package db

import (
    "fmt"
    "log"
    "database/sql"
    config "github.com/ammoses89/thrust-workers/config"
    _ "github.com/lib/pq"
)

type Postgres struct {
    cfg *config.ConnectionSettings
}

func NewPostgres(cfg *config.ConnectionSettings) *Postgres {
    if cfg.Url == "" {
        cfg.Url = fmt.Sprintf("postgres://%s:%s/%s", cfg.Host, cfg.Port, cfg.Database)
    }
    return &Postgres{cfg: cfg}
}

func (pg *Postgres) GetConn() *sql.DB {
    db, err := sql.Open("postgres", pg.cfg.Url)
    if err != nil {
        log.Fatal(err)
    }

    return db
}