//
// http-blink1.go
//
// By Nick Granado, 2013
//
//

package main

import "github.com/heatxsink/blink1-go"

import (
	"fmt"
	"net/http"
)

var (
	blink_usb = blink1.NewBlink()
)

func handler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Path[1:]	
	color := "off"
	
	if param == "red" {
		blink_usb.SetRGB(255, 0, 0)
		color = "red"
	} 
	
	if param == "green" {
		blink_usb.SetRGB(0, 255, 0)
		color = "green"
	}
	
	if param == "blue" {
		blink_usb.SetRGB(0, 0, 255)
		color = "blue"
	}
	
	if param == "white" {
		blink_usb.SetRGB(255, 255, 255)
		color = "white"
	}

	if param == "" {
		blink_usb.SetRGB(0, 0, 0)
	}
	
	fmt.Fprintf(w, "The blink(1) is %s", color)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}	

