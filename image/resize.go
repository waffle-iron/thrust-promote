package resize

import (
    "github.com/nfnt/resize"
    helpers "github.com/ammoses89/thrust-workers/helpers"
)

func Resize(sourceImg string) {
    source, _ := os.Open(sourceImg)
    defer source.Close()

    sourceImage, _ := png.Decode(source)
    newImage := resize.Resize(1280, 720, sourceImage, resize.Lanczos3)

    sourceImgBasename := helpers.RemoveFileExt(sourceImg) 
    resizedImage, _ := os.Create(fmt.Sprintf("%s-%s-%s", sourceImgBasename, "-resized", ".png"))
    png.Encode(resizedImage, newImage, &png.Options{Quality: png.DefaultQuality})
    defer resizedImage.Close()
}