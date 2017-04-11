package image

import (
    "os"
    "log"
    "image"
    "image/jpeg"
    "image/png"
    helpers "github.com/ammoses89/thrust-workers/helpers"
)

func ConvertToPNG(sourceImg string) (string, error) {
    source, err := os.Open(sourceImg)
    if err != nil {
        return nil, err
    }
    defer source.Close()
    jpegImage, err := jpeg.Decode(file)

    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    defer jpegImage.Close()

    sourceImgBasename := helpers.RemoveFileExt(sourceImg) 
    pngImageFilename := fmt.Sprintf("%s-%s-%s", sourceImgBasename, ".png")
    pngImage, err := os.Create(pngImageFilename)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    err = png.Encode(pngImage, jpegImage)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    return pngImageFilename, nil
}