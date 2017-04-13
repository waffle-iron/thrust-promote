package image

import (
    "os"
    "fmt"
    "image"
    "image/draw"
    "image/png"
    helpers "github.com/ammoses89/thrust-workers/helpers"
)

func Watermark(sourceImg string, watermarkImg string) (string, error) {

    // we assume that the image has been coverted to PNG
    source, err := os.Open(sourceImg)
    if err != nil {
        return "", err
    }
    defer source.Close()

    sourceImage, err := png.Decode(source)
    if err != nil {
        return "", err
    }

    // Open and decode watermark PNG
    watermark, err := os.Open(watermarkImg)
    if err != nil {
        return "", err
    }
    defer watermark.Close()

    watermarkImage, err := png.Decode(watermark)
    if err != nil {
        return "", err
    }

    // Watermark offset 20 px from bottom and right
    sourceImageBounds := sourceImage.Bounds()
    x := sourceImageBounds.Max.X - 200
    y := sourceImageBounds.Max.Y - 200
    offset := image.Pt(x, y)

    // create new image with watermark
    newImage := image.NewRGBA(sourceImageBounds)
    draw.Draw(newImage, sourceImageBounds, sourceImage, image.ZP, draw.Src)
    draw.Draw(newImage, watermarkImage.Bounds().Add(offset), watermarkImage, image.ZP, draw.Over)

    // save new file
    sourceImgBasename := helpers.RemoveFileExt(sourceImg) 
    watermarkedImageFilename := fmt.Sprintf("%s%s%s", sourceImgBasename, "-watermarked", ".png")
    watermarkedImage, err := os.Create(watermarkedImageFilename)
    if err != nil {
        return "", err
    }
    png.Encode(watermarkedImage, newImage)
    defer watermarkedImage.Close()

    return watermarkedImageFilename, nil
} 