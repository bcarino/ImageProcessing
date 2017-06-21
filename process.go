package main

import (
	"strings"

	bimg "gopkg.in/h2non/bimg.v1"
)

func ResizeProcess(url string, r ResizeParameter, imageType string) {
	buffer, err := bimg.Read(url)
	IsError(err)
	image := bimg.NewImage(buffer)
	var newImage []byte
	size, err := bimg.Size(buffer)
	originWidth, originHeight := size.Width, size.Height
	resizeWidth, resizeHeight := originWidth, originHeight
	IsError(err)
	xaxis, yaxis := r.xaxis, r.yaxis

	if r.ResizeOption.crop {
		if originHeight/r.height.value <= originWidth/r.width.value {
			r.mode = "BASEDONHEIGHT"
		}
	}

	if r.ResizeOption.force {
		switch strings.ToUpper(r.width.modifier) {
		case "P", "PERCENT":
			resizeWidth = originWidth * r.width.value / 100
		default:
			resizeWidth = r.width.value
		}

		switch strings.ToUpper(r.height.modifier) {
		case "P", "PERCENT":
			resizeHeight = originHeight * r.height.value / 100
		default:
			resizeHeight = r.height.value
		}

	} else {
		if r.mode == "BASEDONHEIGHT" {
			switch strings.ToUpper(r.width.modifier) {
			case "P", "PERCENT":
				resizeWidth = originWidth * r.height.value / 100
			default:
				resizeWidth = originWidth * r.height.value / originHeight
			}

			switch strings.ToUpper(r.height.modifier) {
			case "P", "PERCENT":
				resizeHeight = originHeight * r.height.value / 100
			default:
				resizeHeight = r.height.value
			}
		} else {
			switch strings.ToUpper(r.width.modifier) {
			case "P", "PERCENT":
				resizeWidth = originWidth * r.width.value / 100
			default:
				resizeWidth = r.width.value
			}

			switch strings.ToUpper(r.height.modifier) {
			case "P", "PERCENT":
				resizeHeight = originHeight * r.width.value / 100
			default:
				resizeHeight = originHeight * r.width.value / originWidth
			}
		}

	}

	if r.ResizeOption.density {
		resizeWidth *= r.density
		resizeHeight *= r.density
	}

	options := bimg.Options{
		Width:  resizeWidth,
		Height: resizeHeight,
	}
	newImage, err = image.Process(options)
	IsError(err)

	cropWidth, cropHeight := resizeWidth, resizeHeight

	if strings.ToUpper(r.height.modifier) == "P" || strings.ToUpper(r.height.modifier) == "PERCENT" {
		if r.mode == "BASEDONHEIGHT" {
			cropWidth = originWidth * r.width.value / 100
		} else {
			cropHeight = originHeight * r.width.value / 100
		}
	}

	if cropWidth+r.xaxis > resizeWidth {
		xaxis = resizeWidth - cropWidth
	}
	if cropHeight+r.yaxis > resizeHeight {
		yaxis = resizeHeight - cropHeight
	}

	if r.ResizeOption.crop {
		options = bimg.Options{
			AreaWidth:  cropWidth,
			AreaHeight: cropHeight,
			Top:        yaxis,
			Left:       xaxis,
		}
		newImage, err = image.Process(options)
		IsError(err)
	}

	if r.ResizeOption.Option.flip {
		newImage, err = image.Flip()
		IsError(err)
	}
	if r.ResizeOption.Option.flop {
		newImage, err = image.Flop()
		IsError(err)
	}

	if r.ResizeOption.Option.watermark {
		watermarkBuffer, err := bimg.Read("watermark.png")
		IsError(err)
		watermarkImage := bimg.WatermarkImage{
			Left:    10,
			Top:     10,
			Buf:     watermarkBuffer,
			Opacity: 1.0,
		}
		options = bimg.Options{
			WatermarkImage: watermarkImage,
		}
		newImage, err = image.Process(options)
		IsError(err)
	}

	if r.ResizeOption.Option.grey {
		newImage, err = image.Colourspace(bimg.InterpretationBW)
		IsError(err)
	}

	var convertType bimg.ImageType
	switch imageType {
	case "jpg", "jpeg":
		convertType = bimg.JPEG
	case "webp":
		convertType = bimg.WEBP
	case "png":
		convertType = bimg.PNG
	case "tif", "tiff":
		convertType = bimg.TIFF
	case "gif":
		convertType = bimg.GIF
	case "pdf":
		convertType = bimg.PDF
	case "svg":
		convertType = bimg.SVG
	case "magick":
		convertType = bimg.MAGICK
	}

	options = bimg.Options{
		Type: convertType,
	}
	newImage, err = image.Process(options)
	IsError(err)

	output := "output." + imageType
	bimg.Write(output, newImage)
}

func CropProcess(url string, c CropParameter, imageType string) {
	buffer, err := bimg.Read(url)
	IsError(err)
	image := bimg.NewImage(buffer)
	var newImage []byte
	size, err := bimg.Size(buffer)
	originWidth, originHeight := size.Width, size.Height
	cropWidth, cropHeight := originWidth, originHeight
	IsError(err)
	xaxis, yaxis := c.xaxis, c.yaxis
	top, right, bottom, left := c.trim.top, c.trim.right, c.trim.bottom, c.trim.left

	switch strings.ToUpper(c.width.modifier) {
	case "P", "PERCENT":
		cropWidth = originWidth * c.width.value / 100
	default:
		cropWidth = c.width.value
	}

	switch strings.ToUpper(c.height.modifier) {
	case "P", "PERCENT":
		cropHeight = originHeight * c.height.value / 100
	default:
		cropHeight = c.height.value
	}

	if cropWidth+c.xaxis > originWidth {
		xaxis = originWidth - cropWidth
	}

	if cropHeight+c.yaxis > originHeight {
		yaxis = originHeight - cropHeight
	}

	if c.CropOption.trim {
		if c.trim.left >= originWidth {
			xaxis = 0
		} else {
			xaxis = left
		}

		if c.trim.top >= originHeight {
			yaxis = 0
		} else {
			yaxis = top
		}

		if c.trim.right >= originWidth {
			right = 0
		}

		if c.trim.bottom >= originHeight {
			bottom = 0
		}

		cropWidth = originWidth - right - xaxis
		cropHeight = originHeight - bottom - yaxis
	}

	options := bimg.Options{
		AreaWidth:  cropWidth,
		AreaHeight: cropHeight,
		Top:        yaxis,
		Left:       xaxis,
	}

	newImage, err = image.Process(options)

	if c.CropOption.Option.flip {
		newImage, err = image.Flip()
		IsError(err)
	}
	if c.CropOption.Option.flop {
		newImage, err = image.Flop()
		IsError(err)
	}

	if c.CropOption.Option.watermark {
		watermarkBuffer, err := bimg.Read("watermark.png")
		IsError(err)
		watermarkImage := bimg.WatermarkImage{
			Left:    10,
			Top:     10,
			Buf:     watermarkBuffer,
			Opacity: 1.0,
		}
		options = bimg.Options{
			WatermarkImage: watermarkImage,
		}
		newImage, err = image.Process(options)
	}

	if c.CropOption.Option.grey {
		newImage, err = image.Colourspace(bimg.InterpretationBW)
		IsError(err)
	}

	var convertType bimg.ImageType
	switch imageType {
	case "jpg", "jpeg":
		convertType = bimg.JPEG
	case "webp":
		convertType = bimg.WEBP
	case "png":
		convertType = bimg.PNG
	case "tif", "tiff":
		convertType = bimg.TIFF
	case "gif":
		convertType = bimg.GIF
	case "pdf":
		convertType = bimg.PDF
	case "svg":
		convertType = bimg.SVG
	case "magick":
		convertType = bimg.MAGICK
	}

	options = bimg.Options{
		Type: convertType,
	}
	newImage, err = image.Process(options)
	IsError(err)

	output := "output." + imageType
	bimg.Write(output, newImage)
}
