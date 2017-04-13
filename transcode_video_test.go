package main

import (
    "testing"
    "encoding/json"
    "github.com/stretchr/testify/assert"
)

func TestTranscodeVideo(t *testing.T) {
    sourceUrlPath := "test/unstaged/audio/test.flac"
    sourceImageUrlPath := "test/unstaged/image/test.jpg"
    targetUrlPath := "test/staged/video/test.mp4"

    payload := VideoTranscodePayload{
        SourceUrl: sourceUrlPath,
        TargetUrl: targetUrlPath,
        ImageUrl: sourceImageUrlPath,
        TranscodeType: "wav",
        TrackID: 1,
    }

    metadata, err := json.Marshal(payload)
    task := NewTask("transcode_video", string(metadata))
    status, err := TranscodeVideo(task)
    if assert.NoError(t, err) {
        assert.Equal(t, status, true, "Successful transcode")
    }
}