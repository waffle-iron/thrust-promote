package main

import (
    "testing"
    "encoding/json"
    "github.com/stretchr/testify/assert"
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
    task := NewTask("transcode_audio", string(metadata))
    status, err := TranscodeAudio(task)
    if assert.NoError(t, err) {
        assert.Equal(t, status, true, "Successful transcode")

    }
}