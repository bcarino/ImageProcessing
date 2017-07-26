package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	bimg "gopkg.in/h2non/bimg.v1"
)

// P เป็น PathParam ที่เก็บค่า mode, parameter,image url,image type ที่รับมาผ่าน path
var P PathParam

// B เป็น BImage ที่เก็บค่าสำหรับนำไปใช้กับ package bimg
var B BImage

// I เป็น Image ที่เก็บค่าต่างๆของรูปภาพ เพื่อใช้ทำ image processing
var I Image

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello2!\n")
}

//ImageProcessing เป็นการประมวลผลภาพ
func ImageProcessing(w http.ResponseWriter, r *http.Request) {

	//กำหนดค่าเริ่มต้น
	Initialize()

	//รับค่าผ่าน path และกำหนดค่าให้กับ PathParam
	vars := mux.Vars(r)
	err := Validation(vars["mode"], vars["params"], vars["url"])
	if err != nil {
		Error404()
	} else {
		err := Process()
		if err != nil {
			Error500()
		}
	}

	//Write image to http.ResponseWriter
	b := B.newImage
	switch bimg.DetermineImageTypeName(b) {
	case "jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	case "gif":
		w.Header().Set("Content-Type", "image/gif")
	case "svg":
		w.Header().Set("Content-Type", "image/svg+xml")
	case "png":
		w.Header().Set("Content-Type", "image/png")
	case "webp":
		w.Header().Set("Content-Type", "image/webp")
	default:
		w.Header().Set("Content-Type", "image/jpeg")
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	w.Write(b)
}

func Process() error {

	var err error
	//เอา ความกว้าง ความยาว ของรูปภาพต้นฉบับ
	var size bimg.ImageSize
	size, err = bimg.Size(B.buffer)
	if err != nil {
		return err
	}
	B.width, B.height = size.Width, size.Height
	B.ratio = float64(B.width) / float64(B.height)

	//เอา type ของภาพ
	P.imgType = bimg.DetermineImageTypeName(B.buffer)

	//สร้าง Image ใหม่ จากรูปที่อ่านมา
	B.image = bimg.NewImage(B.buffer)

	//ตรวจสอบ mode การทำ image processing ว่าจะ resize หรือ crop
	switch strings.ToLower(P.mode) {
	case "r", "resize":
		err = ResizeForce()
		if err != nil {
			return err
		}
		err = ResizeCrop()
		if err != nil {
			return err
		}
	case "c", "crop":
		err = Crop()
		if err != nil {
			return err
		}
	case "s", "smart":
		err = SmartCrop()
		if err != nil {
			return err
		}
	}
	if P.flip {
		err = flip()
		if err != nil {
			return err
		}
	}
	if P.flop {
		err = flop()
		if err != nil {
			return err
		}
	}
	if P.watermark {
		err = watermark()
		if err != nil {
			return err
		}
	}
	if P.grey {
		err = grey()
		if err != nil {
			return err
		}
	}

	err = setType()
	if err != nil {
		return err
	}

	//เขียนรูปภาพออกมาเป็นไฟล์ output._

	output := "output/output." + P.imgType
	a, _ := filepath.Abs(output)
	newpath := filepath.Join(".", "output/temp")

	if _, err := os.Stat(newpath); os.IsNotExist(err) {
		os.Mkdir(newpath, os.ModePerm)
		bimg.Write(a, B.newImage)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}
