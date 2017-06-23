package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func Resize(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	vars := mux.Vars(r)
	paramString := vars["params"]
	paramUrl := vars["url"]
	paramType := vars["type"]
	url := UrlDecode(paramUrl)
	resizeParameter := SplitResizeParameters(paramString)
	ResizeProcess(url, resizeParameter, strings.ToLower(paramType))

	fmt.Fprintf(w, "process successfully in: %v", time.Since(start))
}

func Crop(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	vars := mux.Vars(r)
	paramString := vars["params"]
	paramUrl := vars["url"]
	paramType := vars["type"]
	url := UrlDecode(paramUrl)
	cropParameter := SplitCropParameters(paramString)
	CropProcess(url, cropParameter, strings.ToLower(paramType))

	fmt.Fprintf(w, "process successfully in: %v", time.Since(start))
}

func SmartCrop(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	vars := mux.Vars(r)
	paramString := vars["params"]
	paramUrl := vars["url"]
	paramType := vars["type"]
	url := UrlDecode(paramUrl)
	smartParameter := SplitSmartParameters(paramString)
	SmartProcess(url, smartParameter, strings.ToLower(paramType))

	// sm("haarcascade_eye_tree_eyeglasses")
	// sm("haarcascade_eye")
	// sm("haarcascade_frontalface_alt_tree")
	// sm("haarcascade_frontalface_alt")
	// sm("haarcascade_frontalface_alt2")
	// sm("haarcascade_frontalface_default")
	// sm("haarcascade_fullbody")
	// sm("haarcascade_lefteye_2splits")
	// sm("haarcascade_lowerbody")
	// sm("haarcascade_profileface")
	// sm("haarcascade_righteye_2splits")
	// sm("haarcascade_smile")
	// sm("haarcascade_upperbody")

	fmt.Fprintf(w, "process successfully in: %v", time.Since(start))
}
