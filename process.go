package main

import (
	b64 "encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	bimg "gopkg.in/h2non/bimg.v1"
)

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
		case "x", "xaxis":
			P.x = Measure{
				value:    value,
				modifier: modifier,
			}
		case "y", "yaxis":
			P.y = Measure{
				value:    value,
				modifier: modifier,
			}
		case "vf", "flip":
			P.flip = true
		case "hf", "flop":
			P.flop = true
		case "g", "greyscale":
			P.grey = true
		case "wm", "watermark":
			P.watermark = true
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

//SetFitWidthHeight กำหนดความกว้าง ความสูง ให้ได้ภาพที่ใหญ่ที่สุดภายใต้ความกว้าง ความสูงนั้นๆ
func SetFitWidthHeight() {
	SetWidthValue()
	SetHeightValue()
	if P.height.value == 0 {
		I.height = B.height * I.width / B.width
		I.ratio = float64(I.width) / float64(I.height)
	} else {
		I.ratio = float64(I.width) / float64(I.height)
		if I.ratio <= B.ratio {
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

//URLDecode แปลง url ที่อยุ่ในรูปฐาน 64
func URLDecode(code string) (string, error) {
	decode, err := b64.StdEncoding.DecodeString(code)
	if err != nil {
		return "", err
	}
	return string(decode), nil
}

//Error404 กำหนดรูป Error code 404
func Error404() {
	B.newImage, _ = bimg.Read("404.png")
}

//Error500 กำหนดรูป Error code 500
func Error500() {
	B.newImage, _ = bimg.Read("500.jpg")
}

func Validation(mode string, params string, url string) error {
	var err error
	P.mode = mode

	err = SetParam(params)
	if err != nil {
		return err
	}

	P.url, err = URLDecode(url)
	if err != nil {
		return err
	}

	response, e := http.Get(P.url)
	if e != nil {
		return e
	}

	defer response.Body.Close()

	B.buffer, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return nil
}
