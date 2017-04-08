package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
    "github.com/RichardKnop/uuid"
)

func CreateTranscodeAudioTask(machine *Machine, pg *Postgres) string {
	// TODO add task to worker
	var payload AudioTranscodePayload
	if err := json.Unmarshal(data, &payload); err != nil {
		fmt.Println("Could not parse JSON: %v", err)
	}

	// add UUID

	task := Task{
        Id: fmt.Sprintf("task-%v", uuid.New())
		Name:     "transcode_audio",
		Metadata: payload,
	}
	machine.SendTask(&task)
	return "{\"status\": 200}"
}

func removeFileExt(filename string) {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func TranscodeAudio(task Task) (bool, error) {
	var payload AudioTranscodePayload
	task.DeserializeMetadata(&payload)

    extname := filepath.Ext(payload.SourceUrl)
    filename := fmt.Sprintf("audio_dl_%s-%s", task.Id, extname) 

	// grab file
	DownloadFromGCS(payload.SourceUrl, filename)

	// transcode
	basename := removeFileExt(filename)
	targetFilename := basename + "." + payload.transcode_type
	exec.Command("ffmpeg -i %s %s", filename, targetFilename)

	// upload to gcs
	UploadToGCS(targetFilename, payload.TargetUrl)

	return true, nil
}
