package main

import (
    "fmt"
)

func CreateTranscodeVideoTask(rw http.ResponseWriter, req *http.Request, machine *Machine, pg *dbc.Postgres) string {
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
    DownloadFromGCS(payload.SourceUrl, filename)

    // create file path to download to
    basename := helpers.RemoveFileExt(filename)
    targetFilename := fmt.Sprintf("%s.%s", basename, payload.TranscodeType)

    // transcode
    err = helpers.ConvertAudioCommand(filename, targetFilename)
    if err != nil {
        return false, err
    }


    imageExtname := filepath.Ext(payload.ImageUrl)
    imageFilename := fmt.Sprintf("/tmp/image_dl_%s-%s", task.Id, imageExtname)

    // grab image file
    DownloadFromGCS(payload.ImageUrl, imageFilename)

    jpegFileTypes := []string{".jpeg", ".jpg"}
    for _, imageFileType := range jpegFileTypes {
        if imageExtname == imageFileType {
            imageFilename, err = imgPkg.ConvertToPNG(imageFilename)
            if err != nil {
                return nil, err
            }
            break
        }
    }

    videoTargetFilename := fmt.Sprintf("/tmp/video_render_%s-.mp4", task.Id)
    err = helpers.ConvertVideoCommand(targetFilename, imageFilename, videoTargetFilename)
    if err != nil {
        return false, err
    }

    // upload to gcs
    UploadToGCS(videoTargetFilename, payload.TargetUrl)

    // if it fails to delete, don't worry about it
    files := []string{filename, imageFilename, targetFilename, 
        videoTargetFilename}
    helpers.RemoveFiles(files)
    return true, nil
}