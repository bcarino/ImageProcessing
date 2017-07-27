package main

import (
	"image"
	"math"

	"github.com/muesli/smartcrop"

	"bytes"

	bimg "gopkg.in/h2non/bimg.v1"
)

//Crop crop รูปภาพ
func Crop() error {
	var err error
	SetWidthHeightByOriginal()
	NoUpScale()
	NoOutOfBounds()
	B.options = bimg.Options{
		AreaWidth:  I.width,
		AreaHeight: I.height,
		Top:        I.y,
		Left:       I.x,
	}
	B.newImage, err = B.image.Process(B.options)
	if err != nil {
		return err
	}
	return nil
}

//SmartCrop crop รูปภาพโหมดฉลาด
func SmartCrop() error {
	var err error
	// fi, _ := os.Open(P.url)

	// defer fi.Close()

	// img, _, _ := image.Decode(fi)

	img, _, err := image.Decode(bytes.NewReader(B.buffer))
	if err != nil {
		return err
	}

	cascade := "haarcascade_frontalface_default"

	settings := smartcrop.CropSettings{
		FaceDetection:                    true,
		FaceDetectionHaarCascadeFilepath: "/go/src/godocker/haarcascades/" + cascade + ".xml",
	}
	analyzer := smartcrop.NewAnalyzerWithCropSettings(settings)
	topCrop, err := analyzer.FindBestCrop(img, 100, 100)
	if err != nil {
		return err
	}
	// topCropWidth, topCropHeight, topCropX, topCropY := topCrop.Width, topCrop.Height, topCrop.X, topCrop.Y
	// fmt.Printf("%+v", topCrop)

	S := Image{
		width:  topCrop.Width,
		height: topCrop.Height,
		ratio:  float64(topCrop.Width) / float64(topCrop.Height),
		x:      topCrop.X,
		y:      topCrop.Y,
	}

	I.x, I.y = S.x, S.y

	SetWidthHeightByOriginal()
	tempWidth, tempHeight := I.width, I.height

	if I.ratio >= S.ratio {
		I.height = S.height
		I.width = int(float64(I.height) * I.ratio)
		d := math.Abs(float64(I.width-S.width)) / 2
		I.x = I.x - int(d)
		if I.width > B.width {
			I.height = I.height * B.width / I.width
			I.width = B.width
			d := math.Abs(float64(I.height-S.height)) / 2
			I.y = I.y + int(d)
			I.x = 0
		} else {
			totalWidth := I.width + I.x
			if totalWidth > B.width {
				d := math.Abs(float64(totalWidth-B.width)) / 2
				I.x = I.x - int(d)
			}
		}
	} else {
		I.width = S.height
		I.height = int(float64(I.width) / I.ratio)
		d := math.Abs(float64(I.height-S.height)) / 2
		I.y = I.y - int(d)
		if I.height > B.height {
			I.width = I.width * B.height / I.height
			I.height = B.height
			d := math.Abs(float64(I.width-S.width)) / 2
			I.x = I.x + int(d)
			I.y = 0
		} else {
			totalHeight := I.height + I.y
			if totalHeight > B.height {
				d := math.Abs(float64(totalHeight-B.height)) / 2
				I.y = I.y - int(d)
			}
		}
	}

	B.options = bimg.Options{
		AreaWidth:  I.width,
		AreaHeight: I.height,
		Top:        I.y,
		Left:       I.x,
	}

	B.newImage, err = B.image.Process(B.options)
	if err != nil {
		return err
	}

	B.options = bimg.Options{
		Width:  tempWidth,
		Height: tempHeight,
	}

	B.newImage, err = B.image.Process(B.options)
	if err != nil {
		return err
	}

	return nil
}
