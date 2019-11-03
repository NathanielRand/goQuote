package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// http.Get call to quote API
	response, err := http.Get("http://api.theysaidso.com/qod")
	if err != nil {
		log.Fatalf("Error while calling quote API via http.GET %v\n", err)
		return
	}
	// Defer close of http.Get call
	defer response.Body.Close()
	// Grab the entire response body.
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error while reading response body %v\n", err)
		return
	}
	// 	Print to Terminal
	fmt.Println(string(contents))
}
