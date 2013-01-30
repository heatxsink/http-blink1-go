//
// http-blink1.go
//
// By Nick Granado, 2013
//
//

package main

import "github.com/heatxsink/blink1-go/blink1"

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Path[1:]	
	b := blink1.NewBlink()
	color := "off"
	
	if param == "red" {
		b.SetRGB(255, 0, 0)
		color = "red"
	} 
	
	if param == "green" {
		b.SetRGB(0, 255, 0)
		color = "green"
	}
	
	if param == "blue" {
		b.SetRGB(0, 0, 255)
		color = "blue"
	}
	
	if param == "white" {
		b.SetRGB(255, 255, 255)
		color = "white"
	}

	if param == "" {
		b.SetRGB(0, 0, 0)
	}
	
	fmt.Fprintf(w, "Light is %s", color)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}	

