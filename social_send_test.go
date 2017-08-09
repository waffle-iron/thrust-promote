package main
/*
    We want to be able to find an access token in the db
    but should we need to refresh? - Buffer delays the send
    and notifies user that a refresh is needed
    but if we have a refresh token we should be able to 
    make that request
*/

import (
    "time"
    "testing"
    "encoding/json"
    "github.com/stretchr/testify/assert"
    dbc "github.com/ammoses89/thrust-promote/db"
    config "github.com/ammoses89/thrust-promote/config"
)

func TestSocialTwitterSend(t *testing.T) {
    /* 
    find the access token from the db
    and send the message
    - if the publishedAt date is in the future
    schedule the post
    */

    accessToken := "836439866992779265-D0t5sPjtuNgh0TdCebK1T37ofBX6rDG"
    tokenSecret := "836439866992779265-D0t5sPjtuNgh0TdCebK1T37ofBX6rDG"
    // insert into DB
    cfg := config.LoadConfig("config/config.yaml")
    //TODO create a test db for this
    pgCfg := cfg.Db.Development
    pg := dbc.NewPostgres(&pgCfg)
    db, err := pg.GetConn()
    assert.NoError(t, err)

    var socialID int
    err = db.QueryRow(`
        INSERT INTO socials(provider, name, oauth_token, token_secret, created_at, updated_at) 
        VALUES($1, $2, $3, $4, $5, $6) returning id;
    `, "twitter", "marsmoses", accessToken, tokenSecret, time.Now(), time.Now()).Scan(&socialID)

    assert.NoError(t, err)
    payload := SocialSendPayload{
        Service: "twitter",
        Message: "thrust bot test",
        SocialID: socialID,
    }

    metadata, err := json.Marshal(payload)
    task := NewTask("social_send", string(metadata))
    status, err := SocialSend(task)
    if assert.NoError(t, err) {
        assert.Equal(t, status, true, "Successful send")
    }

    _, err = db.Exec(`
        DELETE FROM socials 
        WHERE id = $1
    `, socialID)
    assert.NoError(t, err)
}


func TestSocialFacebookPageMessageSend(t *testing.T) {
    /* 
    find the access token from the db
    and send the message
    - if the publishedAt date is in the future
    schedule the post
    */

    accessToken := "836439866992779265-D0t5sPjtuNgh0TdCebK1T37ofBX6rDG"
    pageID := 11930113092581
    // insert into DB
    cfg := config.LoadConfig("config/config.yaml")
    //TODO create a test db for this
    pgCfg := cfg.Db.Development
    pg := dbc.NewPostgres(&pgCfg)
    db, err := pg.GetConn()
    assert.NoError(t, err)

    var socialID int
    err = db.QueryRow(`
        INSERT INTO socials(provider, name, oauth_token, page_id, created_at, updated_at) 
        VALUES($1, $2, $3, $4, $5, $6) returning id;
    `, "facebook", "marsmoses", accessToken, pageID, time.Now(), time.Now()).Scan(&socialID)

    assert.NoError(t, err)
    payload := SocialSendPayload{
        Service: "facebook",
        Message: "thrust bot test",
        SocialID: socialID,
    }

    metadata, err := json.Marshal(payload)
    task := NewTask("social_send", string(metadata))
    status, err := SocialSend(task)
    if assert.NoError(t, err) {
        assert.Equal(t, status, true, "Successful send")
    }

    _, err = db.Exec(`
        DELETE FROM socials 
        WHERE id = $1
    `, socialID)
    assert.NoError(t, err)
}

func TestSocialYoutubeUpload(t *testing.T) {
    /* 
    find the access token from the db
    and send the message
    - if the publishedAt date is in the future
    schedule the post
    */

    accessToken := "ya29.GlsyBLzXD4G0HNpYUrhMEFeuS6PUsaLk9NY4WCodLXzjwzJLp15GqNMmo1Vm5up8mqOUzZ2Db7GXMKeAHb9uUCojdZEBETG8Gp3591PUjIv0MNhUlOY-yB2xHONU"
    // insert into DB
    cfg := config.LoadConfig("config/config.yaml")
    //TODO create a test db for this
    pgCfg := cfg.Db.Development
    pg := dbc.NewPostgres(&pgCfg)
    db, err := pg.GetConn()
    assert.NoError(t, err)

    var socialID int
    err = db.QueryRow(`
        INSERT INTO socials(provider, name, oauth_token, created_at, updated_at) 
        VALUES($1, $2, $3, $4, $5) returning id;
    `, "youtube", "marsmoses", accessToken, time.Now(), time.Now()).Scan(&socialID)

    assert.NoError(t, err)
    payload := SocialSendPayload{
        Service: "youtube",
        Title: "test video",
        Description: "short links and description",
        VideoUrl: "test/staged/video/test.mp4",
        SocialID: socialID,
    }

    metadata, err := json.Marshal(payload)
    task := NewTask("social_send", string(metadata))
    status, err := SocialSend(task)
    if assert.NoError(t, err) {
        assert.Equal(t, status, true, "Successful send")
    }

    _, err = db.Exec(`
        DELETE FROM socials 
        WHERE id = $1
    `, socialID)
    assert.NoError(t, err)
}