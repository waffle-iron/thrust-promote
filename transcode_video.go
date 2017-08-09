package main

import (
    "encoding/json"
    "fmt"
    "time"
    dbc "github.com/ammoses89/thrust-promote/db"
    imgPkg "github.com/ammoses89/thrust-promote/image"
    helpers "github.com/ammoses89/thrust-promote/helpers"
    config "github.com/ammoses89/thrust-promote/config"
    "io/ioutil"
    "log"
    "net/http"
    "path/filepath"
)

func CreateTranscodeVideoTask(rw http.ResponseWriter, req *http.Request, machine *Machine) string {
    // TODO add task to worker
    var payload VideoTranscodePayload
    res, err := ioutil.ReadAll(req.Body)
    if err := json.Unmarshal(res, &payload); err != nil {
        fmt.Println("Could not parse JSON: %v", err)
    }

    metadata, err := json.Marshal(payload)

    if err != nil {
        fmt.Println("Error ocurred: %v", err)
    }

    task := NewTask("transcode_video", string(metadata))
    machine.SendTask(task)
    return "{\"status\": 200}"
}

func TranscodeVideo(task *Task) (bool, error) {
    var payload VideoTranscodePayload
    err := task.DeserializeMetadata(&payload)
    if err != nil {
        log.Fatalf("Failed to deserialize payload: %v", err)
        return false, nil
    }

    extname := filepath.Ext(payload.SourceUrl)
    filename := fmt.Sprintf("/tmp/audio_dl_%s-%s", task.Id, extname)

    // grab file
    log.Println("Downloading Audio File...")
    _, err = DownloadFromGCS(payload.SourceUrl, filename)
    if err != nil {
        return false, err
    }

    // create file path to download to
    basename := helpers.RemoveFileExt(filename)
    targetFilename := fmt.Sprintf("%s.%s", basename, payload.TranscodeType)

    // transcode
    log.Println("Transcoding Audio File...")
    err = helpers.ConvertAudioCommand(filename, targetFilename)
    if err != nil {
        return false, err
    }


    imageExtname := filepath.Ext(payload.ImageUrl)
    imageFilename := fmt.Sprintf("/tmp/image_dl_%s-%s", task.Id, imageExtname)

    // grab image file
    log.Println("Downloading Image File...")
    _, err = DownloadFromGCS(payload.ImageUrl, imageFilename)

    if err != nil {
        return false, err
    }

    jpegFileTypes := []string{".jpeg", ".jpg"}
    for _, imageFileType := range jpegFileTypes {
        if imageExtname == imageFileType {
            imageFilename, err = imgPkg.ConvertToPNG(imageFilename)
            if err != nil {
                return false, err
            }
            break
        }
    }

    videoTargetFilename := fmt.Sprintf("/tmp/video_render_%s.mp4", task.Id)
    log.Println("Converting Video File...")
    err = helpers.ConvertVideoCommand(targetFilename, imageFilename, videoTargetFilename)
    if err != nil {
        return false, err
    }

    // upload to gcs
    log.Println("Uploading Video File...")
    UploadToGCS(videoTargetFilename, payload.TargetUrl)

    // if it fails to delete, don't worry about it
    files := []string{filename, imageFilename, targetFilename, 
        videoTargetFilename}
    helpers.RemoveFiles(files)

    // add filename to database
    cfg := config.LoadConfig("config/config.yaml")
    //TODO create a test db for this
    pgCfg := cfg.Db.Development
    pg := dbc.NewPostgres(&pgCfg)
    db, err := pg.GetConn()
    if err != nil {
        return false, err
    }

    var trackId int
    err = db.QueryRow("SELECT id FROM tracks WHERE id = $1", payload.TrackID).Scan(&trackId)
    if pg.IsNoResultsErr(err) {
        log.Println("No results found")
        return false, err
    }

    if trackId != 0 {
        _, err := db.Exec(`
            INSERT INTO asset_files(url_path, staged, file_type, track_id, created_at, updated_at) 
            VALUES($1, $2, $3, $4, $5, $6)`, payload.TargetUrl, true, "video", trackId, time.Now(), time.Now())
        if err != nil {
            return false, err
        }
    }
    return true, nil
}