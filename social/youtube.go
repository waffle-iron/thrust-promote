package social

import (
    "net/http"
    "io/ioutil"
    "code.google.com/p/google-api-go-client/youtube/v3"
    "code.google.com/p/goauth2/oauth"
)

// https://github.com/youtube/api-samples/blob/master/go/upload_video.go
// https://developers.google.com/youtube/v3/code_samples/go
type YouTube struct {
    ClientID string
    ClientSecret string
}

func MakeYoutube(clientID string, clientSecret string) *Youtube {
    return &YouTube{ClientID: clientID, ClientSecret: clientSecret}
}

func (yt *YouTube) BuildYoutubeClient(accessToken string) *http.Client {
    config := &oauth.Config{
                ClientId:     yt.ClientID,
                ClientSecret: yt.ClientSecret,
                Scope:        youtube.YoutubeUploadScope,
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

    transport := &oauth.Transport{Config: config}
    transport.Token = accessToken
    return transport.Client()
}

func (yt *YouTube) SendVideo(title string, description string, videoFilename string) (string, error) {
    client, err := yt.BuildYoutubeClient()
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
            CategoryId:  "music", // is a number, we'll figure it out
        },
        Status: &youtube.VideoStatus{PrivacyStatus: "unlisted"},
    }

    call := service.Videos.Insert("snippet,status", upload)
    file, err := os.Open(videoFilename)
    defer file.Close()
    if err != nil {
        log.Fatalf("Error opening %v: %v", *filename, err)
        return "", err
    }

    resp, err := call.Media(file).Do()
    if err != nil {
        log.Fatalf("Error making YouTube API call: %v", err)
        return "", err
    }

    body, err := ioutil.ReadAll(resp.Body)
    log.Printf("Response: %s\n", body)
    return string(body), err
}