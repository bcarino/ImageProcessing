package main

import (
	bimg "gopkg.in/h2non/bimg.v1"
)

//ResizeForce resize รูปภาพ ไม่คงอัตราส่วนเดิมของรูปต้นฉบับ
func ResizeForce() error {
	var err error
	if P.force {
		SetWidthHeightByOriginal()
	} else {
		SetFitWidthHeight()
	}
	B.options = bimg.Options{
		Width:  I.width,
		Height: I.height,
	}
	B.newImage, err = B.image.Process(B.options)
	if err != nil {
		return err
	}
	return nil
}

//ResizeCrop resize รูปภาพ โดยคงอัตราส่วนเดิมของรูปต้นฉบับ และ crop ส่วนที่เกิน
func ResizeCrop() error {
	var err error
	if P.crop {
		SetWidthHeightByRatio()
		B.options = bimg.Options{
			Width:  I.width,
			Height: I.height,
		}
		B.newImage, err = B.image.Process(B.options)
		if err != nil {
			return err
		}
		SetOriginalImage()
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
	}
	return nil
}
