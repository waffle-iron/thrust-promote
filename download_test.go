package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
    urlPath := "test/unstaged/audio/test.flac"
    filename := "test_data/test.flac"

    newFilename := DownloadFromGCS(urlPath, filename)
    assert.Equal(t, filename, newFilename, "Filename should be passed back")
    // _, err := os.Stat(newFilename)
    // if assert.NotError(t, err) {
    //     if err = os.Remove(newFilename); err != nil {
    //         fmt.Println(err.Error())
    //         os.Exit(0)
    //     }
    // }

}