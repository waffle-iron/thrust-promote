package image

import (
    "os"
    "fmt"
    "log"
    "image/jpeg"
    "image/png"
    helpers "github.com/ammoses89/thrust-workers/helpers"
)

func ConvertToPNG(sourceImg string) (string, error) {
    source, err := os.Open(sourceImg)
    if err != nil {
        return "", err
    }
    defer source.Close()
    jpegImage, err := jpeg.Decode(source)

    if err != nil {
        log.Fatal(err)
        return "", err
    }

    sourceImgBasename := helpers.RemoveFileExt(sourceImg) 
    pngImageFilename := fmt.Sprintf("%s%s", sourceImgBasename, ".png")
    pngImage, err := os.Create(pngImageFilename)
    if err != nil {
        fmt.Println(err)
        return "", err
    }

    err = png.Encode(pngImage, jpegImage)
    if err != nil {
        fmt.Println(err)
        return "", err
    }

    return pngImageFilename, nil
}