package image

import (
    "os"
    "image"
    "image/draw"
    "image/png"
    helpers "github.com/ammoses89/thrust-workers/helpers"
)

func Watermark(sourceImg string) (string, error) {

    // we assume that the image has been coverted to PNG
    source, err := os.Open(sourceImg)
    if err != nil {
        return nil, err
    }
    defer source.Close()

    sourceImage, err := png.Decode(source)
    if err != nil {
        return nil, err
    }

    // Open and decode watermark PNG
    watermark, err := os.Open("thrust-watermark.png")
    if err != nil {
        return nil, err
    }
    defer watermark.Close()

    watermarkImage, err := png.Decode(watermark)
    if err != nil {
        return nil, err
    }

    // Watermark offset 20 px from bottom and right
    sourceImageBounds := sourceImage.Bounds()
    x := 20 - sourceImageBounds.Max.X
    y := 20 - sourceImageBounds.Max.Y
    offset := image.Pt(x, y)

    // create new image with watermark
    newImage := image.NewRGBA(sourceImageBounds)
    draw.Draw(newImage, sourceImageBounds, sourceImage, image.Zp. draw.Src)
    draw.Draw(newImage, watermarkImage.Bounds().Add(offset), watermarkImage, image.Zp. draw.Over)

    // save new file
    sourceImgBasename := helpers.RemoveFileExt(sourceImg) 
    watermarkedImageFilename := fmt.Sprintf("%s-%s-%s", sourceImgBasename, "-watermarked", ".png")
    watermarkedImage, err := os.Create(watermarkedImageFilename)
    if err != nil {
        return nil, err
    }
    png.Encode(watermarkedImage, newImage, &png.Options{Quality: png.DefaultQuality})
    defer watermarkedImage.Close()

    return watermarkedImageFilename, nil
} 