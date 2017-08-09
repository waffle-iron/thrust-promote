package main

import (
    "fmt"
    "log"
    "errors"
    "encoding/json"
    "net/http"
    "io/ioutil"
    config "github.com/ammoses89/thrust-promote/config"
    dbc "github.com/ammoses89/thrust-promote/db"
    social "github.com/ammoses89/thrust-promote/social"
)

func CreateSocialSendTask(rw http.ResponseWriter, req *http.Request, 
                              machine *Machine, pg *dbc.Postgres) string {
    // TODO add task to worker
    var payload SocialSendPayload
    res, err := ioutil.ReadAll(req.Body)
    if err := json.Unmarshal(res, &payload); err != nil {
        fmt.Println("Could not parse JSON: %v", err)
    }

    metadata, err := json.Marshal(payload)

    if err != nil {
        fmt.Println("Error ocurred: %v", err)
    }

    task := NewTask("social_send", string(metadata))
    machine.SendTask(task)
    fmt.Println("Save task")
    return "{\"status\": 200}"
}

func SocialSend(task *Task) (bool, error) {
    var payload SocialSendPayload
    err := task.DeserializeMetadata(&payload)
    if err != nil {
        log.Fatalf("Failed to deserialize payload: %v", err)
        return false, err
    }

    var videoTargetFilename string
    if payload.VideoUrl != "" {
        videoTargetFilename = fmt.Sprintf("/tmp/video_ss_%s.mp4", task.Id)
        fmt.Println(videoTargetFilename)
        _, err = DownloadFromGCS(payload.VideoUrl, videoTargetFilename)
        if err != nil {
            return false, err
        }

    }

    cfg := config.LoadConfig("config/config.yaml")
    fmt.Println(payload.Service)
    switch payload.Service {
    case "twitter":
        twit := social.MakeTwitter(cfg.Twitter.ConsumerKey, 
                                   cfg.Twitter.ConsumerSecret)
        _, err = twit.SendMessage(payload.Message, payload.SocialID)
        if err != nil {
            return false, err
        }
    case "facebook":
        // TODO send facebook post
        fb := social.MakeFacebook()
        _, err = fb.SendMessage(payload.Message, payload.SocialID)
        fmt.Printf("Error occured: %v\n", err)
        if err != nil {
            return false, err
        }
    case "youtube":
        yt := social.MakeYoutube(cfg.Youtube.ClientID, cfg.Youtube.ClientSecret)
        _, err = yt.SendVideo(payload.Title, payload.Description, 
            videoTargetFilename, payload.SocialID)
        fmt.Printf("Error occured: %v\n", err)
        if err != nil {
            return false, err
        }
    case "soundcloud":
        return true, nil
    default:
        log.Fatalf("Social services is not supported %s", payload.Service)
        msg := fmt.Sprintf("Social services is not supported %s", payload.Service)
        return false, errors.New(msg)
    }
    return true, nil
}