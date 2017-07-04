package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	b64 "encoding/base64"

	bimg "gopkg.in/h2non/bimg.v1"

	"github.com/gorilla/mux"
)

// P เป็น PathParam ที่เก็บค่า mode, parameter,image url,image type ที่รับมาผ่าน path
var P PathParam

// B เป็น BImage ที่เก็บค่าสำหรับนำไปใช้กับ package bimg
var B BImage

// I เป็น Image ที่เก็บค่าต่างๆของรูปภาพ เพื่อใช้ทำ image processing
var I Image

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

//ImageProcessing เป็นการประมวลผลภาพ
func ImageProcessing(w http.ResponseWriter, r *http.Request) {
	//เวลาเริ่มต้น
	start := time.Now()
	Initialize()
	var err error
	isError := func(e error) {
		if err != nil {
			fmt.Fprintf(w, "error: %v", err)
		}
	}

	//รับค่าผ่าน path และกำหนดค่าให้กับ PathParam
	vars := mux.Vars(r)
	P.mode = vars["mode"]
	err = SetParam(vars["params"])
	isError(err)
	P.url = vars["url"]
	P.imgType = vars["type"]

	//อ่านรูปภาพ โดยต้อง decode url จากฐาน64 ให้เป็นสตริง
	B.buffer, err = bimg.Read(URLDecode(P.url))
	isError(err)

	//เอา ความกว้าง ความยาว ของรูปภาพต้นฉบับ
	var size bimg.ImageSize
	size, err = bimg.Size(B.buffer)
	isError(err)
	B.width, B.height = size.Width, size.Height
	B.ratio = float64(B.width) / float64(B.height)

	//สร้าง Image ใหม่ จากรูปที่อ่านมา
	B.image = bimg.NewImage(B.buffer)

	//ตรวจสอบ mode การทำ image processing ว่าจะ resize หรือ crop
	switch strings.ToLower(P.mode) {
	case "r", "R":
		err = ResizeForce()
		isError(err)
		err = ResizeCrop()
		isError(err)
	case "c", "C":
		err = Crop()
		isError(err)
	case "s", "S":
		SmartCrop()
	}

	//เขียนรูปภาพออกมาเป็นไฟล์ output._
	output := "output." + P.imgType
	bimg.Write(output, B.newImage)

	//แสดงเวลาสิ้นสุด
	fmt.Fprintf(w, "process successfully in: %v", time.Since(start))

	//แสดงค่าทั้งหมด
	//fmt.Fprintf(w, "\nwidth: %+v %v\nheight: %v %v\nx: %v\ny: %v\nforce: %v\ncrop: %v\noriginal width: %v\noriginal height: %v\noriginal ratio: %v\nimage width: %v\nimage height: %v\nimage ratio: %v\n", P.width.value, P.width.modifier, P.height.value, P.height.modifier, P.x.value, P.y.value, P.force, P.crop, B.width, B.height, B.ratio, I.width, I.height, I.ratio)
	fmt.Fprintf(w, "\n%+v\n{width:%v height:%v ratio:%v}\n%v", P, B.width, B.height, B.ratio, I)
}

//URLDecode แปลง url ที่อยุ่ในรูปฐาน 64 ให้เป็นสตริง
func URLDecode(code string) string {
	decode, _ := b64.StdEncoding.DecodeString(code)
	return string(decode)
}
