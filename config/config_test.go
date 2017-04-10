package config

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
    cfg := LoadConfig("config.yaml")
    pgCfg := cfg.Db.Development
    assert.Equal(t, pgCfg.Pool, 5)
    assert.Equal(t, pgCfg.Timeout, 5000)
    assert.Equal(t, pgCfg.Host, "localhost")
    assert.Equal(t, pgCfg.Database, "thrust")
    assert.Equal(t, pgCfg.Port, 5432)
    assert.Equal(t, pgCfg.User, "adrian")
    assert.Equal(t, pgCfg.Password, "")
}

func TestParseUrl(t *testing.T) {
    cfg := LoadConfig("config.yaml")
    redisCfg := cfg.Redis.Development
    assert.Equal(t, redisCfg.Url, "redis://localhost:6379/0")
    redisCfg.ParseUrl()
    assert.Equal(t, redisCfg.Host, "localhost")
    assert.Equal(t, redisCfg.Database, "0")
    assert.Equal(t, redisCfg.Port, 6379)
    assert.Equal(t, redisCfg.User, "")
    assert.Equal(t, redisCfg.Password, "")
}