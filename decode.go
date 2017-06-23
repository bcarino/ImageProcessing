package main

import (
	b64 "encoding/base64"
	"strconv"
	"strings"
)

func UrlDecode(code string) string {
	decode, err := b64.StdEncoding.DecodeString(code)
	IsError(err)
	return string(decode)
}

func SplitResizeParameters(paramString string) ResizeParameter {
	var resizeParameter ResizeParameter
	splitParams := strings.Split(paramString, ",")
	for _, r := range splitParams {
		i := 0
		j := len(r) - 1
		for i < len(r) {
			if r[i] <= 57 && r[i] >= 48 {
				break
			} else {
				i++
			}
		}
		for j >= 0 {
			if r[j] <= 57 && r[j] >= 48 {
				break
			} else {
				j--
			}
		}
		if i == len(r) {
			r += "/"
		}
		switch strings.ToUpper(r[0:i]) {
		case "W", "WEIGHT":
			resizeParameter.width.value, _ = strconv.Atoi(r[i : j+1])
			resizeParameter.width.modifier = r[j+1:]
		case "H", "HEIGHT":
			resizeParameter.height.value, _ = strconv.Atoi(r[i : j+1])
			resizeParameter.height.modifier = r[j+1:]
		case "C", "CROP":
			resizeParameter.ResizeOption.crop = true
		case "F", "FORCE":
			resizeParameter.ResizeOption.force = true
		case "PD", "DENSITY":
			resizeParameter.density, _ = strconv.Atoi(r[i : j+1])
			resizeParameter.ResizeOption.density = true
		case "BH", "BASEDONHEIGHT":
			resizeParameter.mode = "BASEDONHEIGHT"
		case "BW", "BASEDONWIDTH":
			resizeParameter.mode = "BASEDONWIDTH"
		case "X", "XA", "XAXIS":
			resizeParameter.xaxis, _ = strconv.Atoi(r[i : j+1])
		case "Y", "YA", "YAXIS":
			resizeParameter.yaxis, _ = strconv.Atoi(r[i : j+1])
		case "VF", "FLIP":
			resizeParameter.ResizeOption.Option.flip = true
		case "HF", "FLOP":
			resizeParameter.ResizeOption.Option.flop = true
		case "G", "GREY":
			resizeParameter.ResizeOption.Option.grey = true
		case "WM", "WATERMARK":
			resizeParameter.ResizeOption.Option.watermark = true
		}
	}
	return resizeParameter
}

func SplitCropParameters(paramString string) CropParameter {
	var cropParameter CropParameter
	splitParams := strings.Split(paramString, ",")
	for _, r := range splitParams {
		i := 0
		j := len(r) - 1
		for i < len(r) {
			if r[i] <= 57 && r[i] >= 48 {
				break
			} else {
				i++
			}
		}
		for j >= 0 {
			if r[j] <= 57 && r[j] >= 48 {
				break
			} else {
				j--
			}
		}
		if i == len(r) {
			r += "/"
		}
		switch strings.ToUpper(r[0:i]) {
		case "W", "WEIGHT":
			cropParameter.width.value, _ = strconv.Atoi(r[i : j+1])
			cropParameter.width.modifier = r[j+1:]
		case "H", "HEIGHT":
			cropParameter.height.value, _ = strconv.Atoi(r[i : j+1])
			cropParameter.height.modifier = r[j+1:]
		case "X", "XA", "XAXIS":
			cropParameter.xaxis, _ = strconv.Atoi(r[i : j+1])
		case "Y", "YA", "YAXIS":
			cropParameter.yaxis, _ = strconv.Atoi(r[i : j+1])
		case "T", "TOP":
			cropParameter.CropOption.trim = true
			cropParameter.trim.top, _ = strconv.Atoi(r[i : j+1])
		case "R", "RIGHT":
			cropParameter.CropOption.trim = true
			cropParameter.trim.right, _ = strconv.Atoi(r[i : j+1])
		case "B", "BOTTOM":
			cropParameter.CropOption.trim = true
			cropParameter.trim.bottom, _ = strconv.Atoi(r[i : j+1])
		case "L", "LEFT":
			cropParameter.CropOption.trim = true
			cropParameter.trim.left, _ = strconv.Atoi(r[i : j+1])
		case "VF", "FLIP":
			cropParameter.CropOption.Option.flip = true
		case "HF", "FLOP":
			cropParameter.CropOption.Option.flop = true
		case "G", "GREY":
			cropParameter.CropOption.Option.grey = true
		case "WM", "WATERMARK":
			cropParameter.CropOption.Option.watermark = true
		}
	}
	return cropParameter
}

func SplitSmartParameters(paramString string) SmartParameter {
	var smartParameter SmartParameter
	splitParams := strings.Split(paramString, ",")
	for _, r := range splitParams {
		i := 0
		j := len(r) - 1
		for i < len(r) {
			if r[i] <= 57 && r[i] >= 48 {
				break
			} else {
				i++
			}
		}
		for j >= 0 {
			if r[j] <= 57 && r[j] >= 48 {
				break
			} else {
				j--
			}
		}
		if i == len(r) {
			r += "/"
		}
		switch strings.ToUpper(r[0:i]) {
		case "W", "WEIGHT":
			smartParameter.width.value, _ = strconv.Atoi(r[i : j+1])
			smartParameter.width.modifier = r[j+1:]
		case "H", "HEIGHT":
			smartParameter.height.value, _ = strconv.Atoi(r[i : j+1])
			smartParameter.height.modifier = r[j+1:]
		case "C", "CASCADE":
			smartParameter.cascade, _ = strconv.Atoi(r[i : j+1])
		}
	}
	return smartParameter
}
