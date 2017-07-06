package main

import bimg "gopkg.in/h2non/bimg.v1"

//flip  กลับด้ายรูปภาพตามแกน Y
func flip() error {
	var err error
	B.newImage, err = B.image.Flip()
	if err != nil {
		return err
	}
	return nil
}

//flop กลับด้านรูปภาพตามแกน X
func flop() error {
	var err error
	B.newImage, err = B.image.Flop()
	if err != nil {
		return err
	}
	return nil
}

//grey ปรับสีภาพเป็นขาวดำ
func grey() error {
	var err error
	B.newImage, err = B.image.Colourspace(bimg.InterpretationBW)
	if err != nil {
		return err
	}
	return nil
}

//wateramrk ใส่ลายน้ำ
func watermark() error {
	var err error
	watermarkBuffer, err := bimg.Read("watermark.png")
	if err != nil {
		return err
	}
	watermarkImage := bimg.WatermarkImage{
		Left:    10,
		Top:     10,
		Buf:     watermarkBuffer,
		Opacity: 1.0,
	}
	options := bimg.Options{
		WatermarkImage: watermarkImage,
	}
	B.newImage, err = B.image.Process(options)
	if err != nil {
		return err
	}
	return nil
}
