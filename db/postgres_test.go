package db

import (
    "testing"
    config "github.com/ammoses89/thrust-workers/config"
    "github.com/stretchr/testify/assert"
)

func TestPostgres(t *testing.T) {
    cfg := config.LoadConfig("../config/config.yaml")
    //TODO create a test db for this
    pgCfg := cfg.Db.Development
    pg := NewPostgres(&pgCfg)
    db, err := pg.GetConn()
    assert.NoError(t, err, "Connection successful")
    defer db.Close()
    err = db.Ping()
    assert.NoError(t, err, "Connection Ping successful")
}
