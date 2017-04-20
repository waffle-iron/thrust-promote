package social

import (
    "os"
    "log"
    "net/http"
    "google.golang.org/api/youtube/v3"
    "golang.org/x/oauth2"
    config "github.com/ammoses89/thrust-workers/config"
    dbc "github.com/ammoses89/thrust-workers/db"
)

// https://github.com/youtube/api-samples/blob/master/go/upload_video.go
// https://developers.google.com/youtube/v3/code_samples/go
type Youtube struct {
    ClientID string
    ClientSecret string
}

func MakeYoutube(clientID string, clientSecret string) *Youtube {
    return &Youtube{ClientID: clientID, ClientSecret: clientSecret}
}

func (yt *Youtube) BuildYoutubeClient(accessToken string) *http.Client {
    config := &oauth2.Config{
                ClientID:     yt.ClientID,
                ClientSecret: yt.ClientSecret,
                Scopes:       []string{youtube.YoutubeUploadScope},
                // AuthURL:      cfg.Installed.AuthURI,
                // TokenURL:     cfg.Installed.TokenURI,
                // RedirectURL:  redirectUri,
                // TokenCache:   oauth.CacheFile(*cacheFile),
                // Get a refresh token so we can use the access token indefinitely
                // AccessType: "offline",
                // If we want a refresh token, we must set this attribute
                // to force an approval prompt or the code won't work.
                // ApprovalPrompt: "force",
    }
    log.Println("AccessToken: ", accessToken)

    return config.Client(oauth2.NoContext, &oauth2.Token{AccessToken: accessToken})
}

func (yt *Youtube) SendVideo(title string, description string, videoFilename string, socialID int) (string, error) {
    cfg := config.LoadConfig("config/config.yaml")
    //TODO create a test db for this
    pgCfg := cfg.Db.Development
    pg := dbc.NewPostgres(&pgCfg)
    db, err := pg.GetConn()
    if err != nil {
        return "", err
    }

    var accessToken string
    err = db.QueryRow(`
        SELECT oauth_token 
        FROM socials WHERE id = $1`, 
        socialID).Scan(&accessToken)

    if pg.IsNoResultsErr(err) {
        log.Println("No results found")
        return "", err
    }

    if err != nil {
        log.Fatalf("Query Error: %v", err)
        return "", err
    }

    client := yt.BuildYoutubeClient(accessToken)
    if err != nil {
        log.Fatalf("Error building OAuth client: %v", err)
        return "", err
    }

    service, err := youtube.New(client)
    if err != nil {
        log.Fatalf("Error creating YouTube client: %v", err)
        return "", err
    }

    upload := &youtube.Video{
        Snippet: &youtube.VideoSnippet{
            Title:       title,
            Description: description,
            CategoryId:  "10",
        },
        Status: &youtube.VideoStatus{PrivacyStatus: "unlisted"},
    }

    call := service.Videos.Insert("snippet,status", upload)
    log.Println(videoFilename)
    videoFile, err := os.Open(videoFilename)
    defer videoFile.Close()
    if err != nil {
        log.Fatalf("Error opening %v: %v", videoFilename, err)
        return "", err
    }

    resp, err := call.Media(videoFile).Do()
    if err != nil {
        log.Fatalf("Error making YouTube API call: %v", err)
        return "", err
    }

    log.Printf("Response: %s\n", resp.Id)
    return "done", err
}