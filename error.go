package main

import "fmt"

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func IsError(err error) {
	if err != nil {
		fmt.Println("\n\n ******************************************************* \n Error: ",
			err, "\n ******************************************************* \n")
	}
}
