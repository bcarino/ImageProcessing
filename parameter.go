package main

type ResizeParameter struct {
	width  Measure
	height Measure
	ResizeOption
	density int
	mode    string
	xaxis   int
	yaxis   int
}

type Measure struct {
	value    int
	modifier string
}

type Option struct {
	flip      bool
	flop      bool
	grey      bool
	watermark bool
}

type ResizeOption struct {
	Option
	crop    bool
	force   bool
	density bool
}

type CropParameter struct {
	width  Measure
	height Measure
	xaxis  int
	yaxis  int
	trim   Edge
	CropOption
}

type CropOption struct {
	Option
	trim bool
}

type Edge struct {
	top    int
	right  int
	bottom int
	left   int
}

type SmartParameter struct {
	width   Measure
	height  Measure
	cascade int
}
