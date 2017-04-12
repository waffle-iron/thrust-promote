package main

import (
	"encoding/json"
	"fmt"
	dbc "github.com/ammoses89/thrust-workers/db"
    helpers "github.com/ammoses89/thrust-workers/helpers"
	"io/ioutil"
	"log"
	"net/http"
	"bytes"
	"os/exec"
	"path/filepath"
)

func CreateTranscodeAudioTask(rw http.ResponseWriter, req *http.Request, 
							  machine *Machine, pg *dbc.Postgres) string {
	// TODO add task to worker
	var payload AudioTranscodePayload
	res, err := ioutil.ReadAll(req.Body)
	if err := json.Unmarshal(res, &payload); err != nil {
		fmt.Println("Could not parse JSON: %v", err)
	}

	metadata, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("Error ocurred: %v", err)
	}

	task := NewTask("transcode_audio", string(metadata))
	machine.SendTask(task)
	return "{\"status\": 200}"
}

func TranscodeAudio(task *Task) (bool, error) {
	var payload AudioTranscodePayload
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

	// upload to gcs
	UploadToGCS(targetFilename, payload.TargetUrl)

	// if it fails to delete, don't worry about it
	helpers.Remove([]string{filename})
	return true, nil
}
