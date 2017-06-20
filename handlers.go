package main

import (
	"fmt"
	"net/http"
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
	resizeParameter := SplitResizeParameters(paramString)
	ResizeProcess("input.jpg", resizeParameter)

	fmt.Fprintf(w, "process successfully in: %v", time.Since(start))
}

func Crop(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	vars := mux.Vars(r)
	paramString := vars["params"]
	cropParameter := SplitCropParameters(paramString)
	CropProcess("input.jpg", cropParameter)

	fmt.Fprintf(w, "process successfully in: %v", time.Since(start))
}
