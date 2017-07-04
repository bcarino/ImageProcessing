package main

import (
	"errors"
	"strconv"
	"strings"

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
	force bool
	crop  bool
}

//Measure เป็นประเภทของพารามิเตอร์ ความกว้าง และ ความยาว ซึ่งประกอบด้วย ค่า(ตัวเลข) และ ส่วนขยาย(สตริง)
type Measure struct {
	value    int
	modifier string
}

//Initialize เป็นการกำหนดค่าเริ่มต้น
func Initialize() {
	var p PathParam
	var b BImage
	var i Image
	P, B, I = p, b, i
}

//SetParam แยก parameter ด้วย ',' แล้วเช็คว่า parameter แต่ละตัวกำหนดค่าให้กับ Params และ Options
func SetParam(s string) error {
	param := strings.Split(s, ",")
	for _, p := range param {
		name, sValue := SplitNameAndValue(p)
		value, modifier := SplitValueAndModifier(sValue)
		switch strings.ToLower(name) {
		case "w", "width":
			P.width = Measure{
				value:    value,
				modifier: modifier,
			}
		case "h", "height":
			P.height = Measure{
				value:    value,
				modifier: modifier,
			}
		case "f", "force":
			P.force = true
		case "c", "crop":
			P.crop = true
		case "x", "xa", "xaxis":
			P.x = Measure{
				value:    value,
				modifier: modifier,
			}
		case "y", "ya", "yaxis":
			P.y = Measure{
				value:    value,
				modifier: modifier,
			}
		default:
			return errors.New("'" + name + "'" + " is not parameter\n")
		}
	}
	return nil
}

//SplitNameAndValue คือการแยก h100p => parametername: "h" , value: "100p"
//SplitNameAndValue แยก ชื่อพารามิเตอร์(สตริง) ออกจาก ค่า(ขึ้นต้นด้วยตัวเลข อาจมีส่วนขยายที่เป็นสตริงตามหลัง)
func SplitNameAndValue(s string) (string, string) {
	i, name := 0, ""
	for i = 0; i < len(s); i++ {
		if s[i] < 48 || s[i] > 57 {
			name += string(s[i])
		} else {
			break
		}
	}
	return strings.ToLower(name), strings.ToLower(s[i:])
}

//SplitValueAndModifier คือการแยก 100p => value: 100 , modifier: "p"
//SplitValueAndModifier เอา ค่า ที่ได้จากการ SplitNameAndValue มาแยก ซึ่งจะได้ ค่า(ตัวเลข) และ ส่วนขยาย(สตริง)
func SplitValueAndModifier(s string) (int, string) {
	i, sValue := 0, "0"
	for i = 0; i < len(s); i++ {
		if s[i] >= 48 && s[i] <= 57 {
			sValue += string(s[i])
		} else {
			break
		}
	}
	value, _ := strconv.Atoi(sValue)
	return value, strings.ToLower(s[i:])
}

//SetWidthValue กำหนดค่า width ของรูปที่จะนำไปทำ image processing
func SetWidthValue() {
	switch P.width.modifier {
	case "p", "percent":
		I.width = B.width * P.width.value / 100
	default:
		I.width = P.width.value
	}
}

//SetHeightValue กำหนดค่า height ของรูปที่จะนำไปทำ image processing
func SetHeightValue() {
	switch P.height.modifier {
	case "p", "percent":
		I.height = B.height * P.height.value / 100
	default:
		I.height = P.height.value
	}
}

//SetWidthHeightByOriginal กำหนดความกว้าง ความสูง โดยอิงจากต้นฉบับ
func SetWidthHeightByOriginal() {
	SetWidthValue()
	SetHeightValue()
	if I.width == 0 {
		I.width = B.width
	} else if I.height == 0 {
		I.height = B.height
	}
	I.ratio = float64(I.width) / float64(I.height)
}

//SetWidthHeightByRatio กำหนดความกว้าง ความสูง โดยอิงจากอัตราส่วน
func SetWidthHeightByRatio() {
	SetWidthValue()
	SetHeightValue()
	if P.height.value == 0 {
		I.height = B.height * I.width / B.width
		I.ratio = float64(I.width) / float64(I.height)
	} else {
		I.ratio = float64(I.width) / float64(I.height)
		if I.ratio >= B.ratio {
			I.height = B.height * I.width / B.width
		} else {
			I.width = B.width * I.height / B.height
		}
	}

}

//NoOutOfBounds ปรับขนาดรูปให้ไม่ให้เกินขอบรูปภาพ ในกรณีเลื่อนตำแหน่งแกน x y
func NoOutOfBounds() {
	SetXValue()
	SetYValue()
	totalWidth := I.x + I.width
	totalHeight := I.y + I.height
	remainWidth := B.width - I.x
	remainHeight := B.height - I.y
	if remainWidth <= 0 {
		I.x = 0
	} else if totalWidth > B.width {
		I.width = remainWidth
	}
	if remainHeight <= 0 {
		I.y = 0
	} else if totalHeight > B.height {
		I.height = remainHeight
	}
}

//NoUpScale ปรับขนาดรูปที่จะทำ image processing ไม่ให้ใหญ่เกินต้นฉบับ
func NoUpScale() {
	if I.width > B.width {
		I.width = B.width
	}
	if I.height > B.height {
		I.height = B.height
	}
}

//SetOriginalImage กำหนดขนาดรูปต้นฉบับใหม่
func SetOriginalImage() {
	B.width = I.width
	B.height = I.height
	SetWidthValue()
	SetHeightValue()
}

//SetXValue กำหนดค่า X ที่จำไปทำ image processing
func SetXValue() {
	I.x = P.x.value
}

//SetYValue กำหนดค่า y ที่จำไปทำ image processing
func SetYValue() {
	I.y = P.y.value
}
