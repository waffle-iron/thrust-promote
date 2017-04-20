package main

import (
    "time"
    "testing"
    "encoding/json"
    "github.com/stretchr/testify/assert"
    config "github.com/ammoses89/thrust-workers/config"
    dbc "github.com/ammoses89/thrust-workers/db"
)

func TestTranscodeVideo(t *testing.T) {
    sourceUrlPath := "test/unstaged/audio/test.flac"
    sourceImageUrlPath := "test/unstaged/image/test.jpg"
    targetUrlPath := "test/staged/video/test.mp4"

    // insert into DB
    cfg := config.LoadConfig("config/config.yaml")
    //TODO create a test db for this
    pgCfg := cfg.Db.Development
    pg := dbc.NewPostgres(&pgCfg)
    db, err := pg.GetConn()
    assert.NoError(t, err)

    var trackID int
    err = db.QueryRow(`
        INSERT INTO tracks(title, description, image_path, audio_path, created_at, updated_at) 
        VALUES($1, $2, $3, $4, $5, $6) returning id;
    `, "fun song", "its a song", sourceImageUrlPath, sourceUrlPath, time.Now(), time.Now()).Scan(&trackID)

    assert.NoError(t, err)

    payload := VideoTranscodePayload{
        SourceUrl: sourceUrlPath,
        TargetUrl: targetUrlPath,
        ImageUrl: sourceImageUrlPath,
        TranscodeType: "wav",
        TrackID: trackID,
    }

    metadata, err := json.Marshal(payload)
    task := NewTask("transcode_video", string(metadata))
    status, err := TranscodeVideo(task)
    if assert.NoError(t, err) {
        assert.Equal(t, status, true, "Successful transcode")
    }

    _, err = db.Exec(`
        DELETE FROM tracks 
        WHERE id = $1
    `, trackID)
    assert.NoError(t, err)

    _, err = db.Exec(`
        DELETE FROM asset_files 
        WHERE track_id = $1
    `, trackID)
    assert.NoError(t, err)
}