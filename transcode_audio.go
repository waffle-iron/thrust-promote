package main

import (
	"encoding/json"
	"fmt"
	"github.com/RichardKnop/uuid"
	dbc "github.com/ammoses89/thrust-workers/db"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CreateTranscodeAudioTask(rw http.ResponseWriter, req *http.Request, machine *Machine, pg *dbc.Postgres) string {
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

	// add UUID
	task := Task{
		Id:       fmt.Sprintf("task-%v", uuid.New()),
		Status:   "Queued",
		Name:     "transcode_audio",
		Metadata: string(metadata)}
	machine.SendTask(&task)
	return "{\"status\": 200}"
}

func removeFileExt(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func TranscodeAudio(task Task) (bool, error) {
	var payload AudioTranscodePayload
	task.DeserializeMetadata(&payload)

	extname := filepath.Ext(payload.SourceUrl)
	filename := fmt.Sprintf("/tmp/audio_dl_%s-%s", task.Id, extname)

	// grab file
	DownloadFromGCS(payload.SourceUrl, filename)

	// create file path to download to
	basename := removeFileExt(filename)
	targetFilename := fmt.Sprintf("%s.%s", basename, payload.TranscodeType)

	// transcode
	cmd := "ffmpeg"
	args := []string{"-i", filename, targetFilename}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		log.Fatalf("Command failed: %v", err)
		return false, err
	}

	// upload to gcs
	UploadToGCS(targetFilename, payload.TargetUrl)

	// if it fails to delete, don't worry about it
	os.Remove(filename) 
	return true, nil
}
