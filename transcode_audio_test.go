package main

import (
    "fmt"
    "testing"
    "encoding/json"
    "github.com/stretchr/testify/assert"
    "github.com/RichardKnop/uuid"
)

func TestTranscodeAudio(t *testing.T) {
    sourceUrlPath := "test/unstaged/audio/test.flac"
    targetUrlPath := "test/staged/audio/test.mp3"

    payload := AudioTranscodePayload{
        SourceUrl: sourceUrlPath,
        TargetUrl: targetUrlPath,
        TranscodeType: "mp3",
        TrackID: 1,
    }

    metadata, err := json.Marshal(payload)

    task := Task{
        Id:       fmt.Sprintf("task-%v", uuid.New()),
        Status:   "Queued",
        Name:     "transcode_audio",
        Metadata: string(metadata)}

    status, err := TranscodeAudio(task)
    if assert.NoError(t, err) {
        assert.Equal(t, status, true, "Successful transcode")

    }
}