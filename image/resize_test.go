package image

import (
    "os"
    "image"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestResizeImage(t *testing.T) {
    sourceImg := "../data/dropbox-example.png"
    resizedFilename, err := Resize(sourceImg)
    assert.NoError(t, err, "Resize failed")
    _, err = os.Stat(resizedFilename)
    assert.NoError(t, err, "File does not exist")
    assert.False(t, os.IsNotExist(err))
    resizedFile, err := os.Open(resizedFilename)
    assert.NoError(t, err, "Could not open resized filename")
    resizedImage, _, err := image.DecodeConfig(resizedFile)
    assert.NoError(t, err, "Image could not be decoded")
    assert.Equal(t, resizedImage.Width, 1280, "Width was not adjusted")
    assert.Equal(t, resizedImage.Height, 720, "Height was not adjusted")
    os.Remove(resizedFilename)
}