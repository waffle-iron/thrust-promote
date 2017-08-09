package image

import (
    "os"
    "fmt"
    "image/png"
    "github.com/nfnt/resize"
    helpers "github.com/ammoses89/thrust-promote/helpers"
)

func Resize(sourceImg string) (string, error) {
    source, err := os.Open(sourceImg)
    if err != nil {
        return "", err
    }
    defer source.Close()

    sourceImage, err := png.Decode(source)
    if err != nil {
        return "", err
    }
    newImage := resize.Resize(1280, 720, sourceImage, resize.Lanczos3)

    sourceImgBasename := helpers.RemoveFileExt(sourceImg) 
    resizedImageFilename := fmt.Sprintf("%s%s%s", sourceImgBasename, "-resized", ".png")
    resizedImage, err := os.Create(resizedImageFilename)
    if err != nil {
        return "", err
    }
    png.Encode(resizedImage, newImage)
    defer resizedImage.Close()

    return resizedImageFilename, nil
}