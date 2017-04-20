package social

/*
Here will be all the logic for making twitter API calls
We'll use a struct to make sure all the necessary componens are
supplied.
This will also provide flexibility
*/

import (
    "fmt"
    "log"
    "net/http"
    "net/url"
    "io/ioutil"
    config "github.com/ammoses89/thrust-workers/config"
    dbc "github.com/ammoses89/thrust-workers/db"
    "github.com/dghubble/oauth1"
)

const TwitterURL = "https://api.twitter.com/1.1"

type Twitter struct {
    ConsumerKey string
    ConsumerSecret string
}

func MakeTwitter(consumerKey string, consumerSecret string) *Twitter {
    return &Twitter{ConsumerKey: consumerKey, ConsumerSecret: consumerSecret} 
}

func (twitter *Twitter) BuildTwitterClient(userAccessToken string, userTokenSecret string) *http.Client {
    clientCfg := oauth1.NewConfig(twitter.ConsumerKey, twitter.ConsumerSecret)
    token := oauth1.NewToken(userAccessToken, userTokenSecret)
    // httpClient will automatically authorize http.Request's
    httpClient := clientCfg.Client(oauth1.NoContext, token)
    return httpClient
}

func (twitter *Twitter) SendMessage(message string, socialID int) (string, error) {
    // add filename to database
    cfg := config.LoadConfig("config/config.yaml")
    //TODO create a test db for this
    pgCfg := cfg.Db.Development
    pg := dbc.NewPostgres(&pgCfg)
    db, err := pg.GetConn()
    if err != nil {
        return "", err
    }
    
    endpoint := fmt.Sprintf("%s%s", TwitterURL, "/statuses/update.json")
    var accessToken, tokenSecret string
    err = db.QueryRow(`
        SELECT oauth_token, token_secret 
        FROM socials WHERE id = $1`, 
        socialID).Scan(&accessToken, &tokenSecret)

    if pg.IsNoResultsErr(err) {
        log.Println("No results found")
        return "", err
    }

    httpClient := twitter.BuildTwitterClient(accessToken, tokenSecret)

    status := url.Values{"status": {message}}
    resp, err := httpClient.PostForm(endpoint, status) 

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