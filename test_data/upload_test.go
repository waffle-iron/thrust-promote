package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestUpload(t *testing.T) {
    urlPath := "test/staged/audio/test.flac"
    filename := "test_data/test.mp3"

    success := UploadToGCS(urlPath, filename)
    assert.Equal(t, 1, success, "File was successfully uploaded")
}