package main

import (
    "fmt"
    "log"
    "net/url"
    "io/ioutil"
    config "github.com/ammoses89/thrust-workers/config"
    "github.com/dghubble/oauth1"
    // "golang.org/x/oauth2"
)

const FacebookURL = "https://graph.facebook.com"
const TwitterURL = "https://api.twitter.com/1.1/statuses/update.json"

func CreateSocialSendTask() string {
    // TODO add task to worker
    fmt.Println("Save task")
    return "{\"status\": 200}"
}

func BuildTwitterClient(message string, userAccessToken string, userTokenSecret string, cfg *config.Config) (*http.Client, error) {
    clientCfg := oauth1.NewConfig(cfg.Twitter.ConsumerKey, cfg.Twitter.ConsumerSecret)
    token := oauth1.NewToken(userAccessToken, userTokenSecret)

    // httpClient will automatically authorize http.Request's
    httpClient := clientCfg.Client(oauth1.NoContext, token)

    status := url.Values{"status": {message}}
    return httpClient.PostForm(TwitterURL, status) 
}

func SocialSend(task *Task) (bool, error) {
    var payload SocialSendPayload
    err := task.DeserializeMetadata(&payload)
    if err != nil {
        log.Fatalf("Failed to deserialize payload: %v", err)
        return false, err
    }

    if payload.Service == "twitter" {
        resp, err := BuildTwitterClient(payload.Message)
    } else {
        // resp, err := BuildTwitterClient(payload.Message)

    }

    if err != nil {
        log.Fatalf("Failed to send: %v", err)
        return false, err
    }
    defer resp.Body.Close()
    // what to do with body?
    body, _ := ioutil.ReadAll(resp.Body)
    log.Printf("Response: %s\n", body)
    return true, nil
}