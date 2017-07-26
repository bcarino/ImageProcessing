package main

import (
	bimg "gopkg.in/h2non/bimg.v1"
)

/*
---------------------------------------------
				คำอธิบาย เพิ่มเติม
---------------------------------------------
width, height ของ Image, BImage, Params
	Params => width: 50p, height: 50p --- 50p หมายถึง 50% ของขนาดรูปต้นฉบับ
	BImage => width: 500, height: 500 --- รูปต้นฉบับมีขนาด 500x500
	Image => width: 250, height: 250 --- เมื่อทำ image processing จะได้รุปที่มีขนาด 250x250
*/

//Image เก็บค่าต่างๆเกี่ยวกับรูปภาพ ที่จะนำไปใช้ทำ image processing
type Image struct {
	width  int //ความกว้างของรูปที่ถูกแปลงค่าเรียบร้อยแล้วเพื่อนำไปทำ image processing
	height int // ความสูงของรูปที่ถูกแปลงค่าเรียบร้อยแล้วเพื่อนำไปทำ image processing
	ratio  float64
	x      int
	y      int
}

//BImage เก็บค่าต่างๆที่เรียกใช้ผ่าน package bimg
type BImage struct {
	buffer   []byte
	image    *bimg.Image
	newImage []byte
	options  bimg.Options
	width    int //ความกว้างของรูปต้นฉบับ
	height   int // ความสูงของรูปต้นฉบับ
	ratio    float64
}

//PathParam เก็บค่าที่รับมาผ่าน path
type PathParam struct {
	mode string
	Params
	url     string
	imgType string
}

//Params เก็บ parameters ที่รับมาผ่าน path
type Params struct {
	width  Measure
	height Measure
	x      Measure
	y      Measure
	Options
}

//Options เก็บ option parameter ที่รับมาผ่าน path
type Options struct {
	force     bool
	crop      bool
	flip      bool //flips the image about the vertical Y axis.
	flop      bool //flops the image about the horizontal X axis.
	grey      bool
	watermark bool
}

//Measure เป็นประเภทของพารามิเตอร์ ความกว้าง และ ความยาว ซึ่งประกอบด้วย ค่า(ตัวเลข) และ ส่วนขยาย(สตริง)
type Measure struct {
	value    int
	modifier string
}
