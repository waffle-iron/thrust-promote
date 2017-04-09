package image

import (
    "image"
    "image/draw"
    "image/png"
    helpers "github.com/ammoses89/thrust-workers/helpers"
)

func Watermark(sourceImg string) (string, error) {

    // we assume that the image has been coverted to PNG
    source, _ := os.Open(sourceImg)
    defer source.Close()

    sourceImage, _ := png.Decode(source)

    // Open and decode watermark PNG
    watermark, _ := os.Open("thrust-watermark.png")
    defer watermark.Close()

    watermarkImage, _ := png.Decode(watermark)

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
    watermarkedImage, _ := os.Create(fmt.Sprintf("%s-%s-%s", sourceImgBasename, "-watermarked", ".png"))
    png.Encode(watermarkedImage, newImage, &png.Options{Quality: png.DefaultQuality})
    defer watermarkedImage.Close()

} 