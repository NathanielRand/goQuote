package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Quote API - TheySaidSo.com
	response, err := http.Get("http://api.theysaidso.com/qod")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 	Print to Terminal
	fmt.Println(string(contents))
}
