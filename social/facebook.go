package social

import (
    config "github.com/ammoses89/thrust-workers/config"
    dbc "github.com/ammoses89/thrust-workers/db"
    "errors"
    "io/ioutil"
    "net/http"
    "os"
    "net/url"
    "bufio"
    "fmt"
    "log"
    // "golang.org/x/oauth2"
)

const FacebookURL = "https://graph.facebook.com"

type Facebook struct {
}

func MakeFacebook() *Facebook {
    return &Facebook{}
}

func (facebook *Facebook) SendMessage(message string, socialID int) (string, error) {

    cfg := config.LoadConfig("config/config.yaml")
    //TODO create a test db for this
    pgCfg := cfg.Db.Development
    pg := dbc.NewPostgres(&pgCfg)
    db, err := pg.GetConn()
    if err != nil {
        return "", err
    }

    var accessToken string
    var pageID int
    err = db.QueryRow(`
        SELECT access_token, page_id 
        FROM socials WHERE id = $1`, 
        socialID).Scan(&accessToken, &pageID)

    if pg.IsNoResultsErr(err) {
        log.Println("No results found")
        return "", err
    }

    if pageID == 0 {
        return "", errors.New("No page id found")
    }

    endpoint := fmt.Sprintf(FacebookURL, "/", pageID, "/", "feed")
    params := url.Values{"message": {message},
                         "access_token": {accessToken}}
    resp, err := http.PostForm(endpoint, params) 

    if err != nil {
        log.Fatalf("Failed to send: %v", err)
        return "", err
    }
    defer resp.Body.Close()
    // what to do with body?
    body, err := ioutil.ReadAll(resp.Body)
    log.Printf("Response: %s\n", body)
    return string(body), err
}

func (facebook *Facebook) SendVideo(videoUrl string, videoFilename string, socialID int) (string, error) {

    cfg := config.LoadConfig("config/config.yaml")
    //TODO create a test db for this
    pgCfg := cfg.Db.Development
    pg := dbc.NewPostgres(&pgCfg)
    db, err := pg.GetConn()
    if err != nil {
        return "", err
    }

    var accessToken string
    var pageID int
    err = db.QueryRow(`
        SELECT access_token, page_id 
        FROM socials WHERE id = $1`, 
        socialID).Scan(&accessToken, &pageID)

    if pg.IsNoResultsErr(err) {
        log.Println("No results found")
        return "", err
    }

    if pageID == 0 {
        return "", errors.New("No page id found")
    }

    filePtr, err := os.Open(videoFilename)

    if err != nil {
        log.Fatalf("Failed to open file: %v", err)
        return "", err
    }


    buf := bufio.NewReader(filePtr)

    if err != nil {
        log.Fatalf("Failed to read file: %v", err)
        return "", err
    }

    params := url.Values{"access_token": {accessToken}}
    endpoint := fmt.Sprintf(FacebookURL, "/", pageID, "/", "videos", params.Encode())
    resp, err := http.Post(endpoint, "video/mp4", buf) 

    if err != nil {
        log.Fatalf("Failed to send: %v", err)
        return "", err
    }
    defer resp.Body.Close()
    // what to do with body?
    body, err := ioutil.ReadAll(resp.Body)
    log.Printf("Response: %s\n", body)
    return string(body), err
}