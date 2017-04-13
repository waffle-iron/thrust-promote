package image

import (
    "os"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestConvertImage(t *testing.T) {
    sourceImg := "../data/example.jpg"
    convertedFilename, err := ConvertToPNG(sourceImg)
    assert.NoError(t, err, "Conversion failed")
    _, err = os.Stat(convertedFilename)
    assert.NoError(t, err, "File does not exist")
    assert.False(t, os.IsNotExist(err))
    os.Remove(convertedFilename)
}