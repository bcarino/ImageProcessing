package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	sm()
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
