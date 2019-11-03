package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func main() {
	info()
	commands()

	err := app.Run(os.Args)
  	if err != nil {
    	log.Fatal(err)
  	}
}

func info() {
	app.Name = "Simple quote CLI"
	app.Usage = "A simple CLI for getting a quote."
	app.Author = "Nathaniel Rand" 
	app.Version = "2.0.0"
}

func commands() {
	app.Commands = []cli.Command{
	  	{
			Name:    "qod",
			Aliases: []string{"qod"},
			Usage:   "Quote of the day",
			Action: func(c *cli.Context) { 
				response, err := http.Get("http://api.theysaidso.com/qod")
				if err != nil {
					log.Fatalf("Error while calling QOD quote API via http.GET %v\n", err)
					return
				}		
				// Defer close of http.Get call
				defer response.Body.Close()		
				// Grab the entire response body.
				// Warning! ioutil.ReadAll is not memory efficient, therefore generally not recommended. 
				contents, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatalf("Error while reading response body %v\n", err)
					return
				}
				// 	Print the entire response body to terminal.
				fmt.Println(string(contents))
			},
		},
	  	{
			Name:    "inspire",
			Aliases: []string{"ins"},
			Usage:   "Inspirational quote",
			Action: func(c *cli.Context) { 
				response, err := http.Get("http://quotes.rest/qod.json?category=inspire")
				if err != nil {
					log.Fatalf("Error while calling QOD quote API via http.GET %v\n", err)
					return
				}		
				// Defer close of http.Get call
				defer response.Body.Close()		
				// Grab the entire response body.
				// Warning! ioutil.ReadAll is not memory efficient, therefore generally not recommended. 
				contents, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatalf("Error while reading response body %v\n", err)
					return
				}
				// 	Print the entire response body to terminal.
				fmt.Println(string(contents))
			},
		},
		{
			Name:    "life",
			Aliases: []string{"life"},
			Usage:   "Life quote",
			Action: func(c *cli.Context) { 
				response, err := http.Get("http://quotes.rest/qod.json?category=life")
				if err != nil {
					log.Fatalf("Error while calling QOD quote API via http.GET %v\n", err)
					return
				}		
				// Defer close of http.Get call
				defer response.Body.Close()		
				// Grab the entire response body.
				// Warning! ioutil.ReadAll is not memory efficient, therefore generally not recommended. 
				contents, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatalf("Error while reading response body %v\n", err)
					return
				}
				// 	Print the entire response body to terminal.
				fmt.Println(string(contents))
			},
		},
		{
			Name:    "Funny",
			Aliases: []string{"funny"},
			Usage:   "Funny quote",
			Action: func(c *cli.Context) { 
				response, err := http.Get("http://quotes.rest/qod.json?category=funny")
				if err != nil {
					log.Fatalf("Error while calling QOD quote API via http.GET %v\n", err)
					return
				}		
				// Defer close of http.Get call
				defer response.Body.Close()		
				// Grab the entire response body.
				// Warning! ioutil.ReadAll is not memory efficient, therefore generally not recommended. 
				contents, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatalf("Error while reading response body %v\n", err)
					return
				}
				// 	Print the entire response body to terminal.
				fmt.Println(string(contents))
			},
		},
		{
			Name:    "management",
			Aliases: []string{"manage"},
			Usage:   "Management quote",
			Action: func(c *cli.Context) { 
				response, err := http.Get("http://quotes.rest/qod.json?category=management")
				if err != nil {
					log.Fatalf("Error while calling QOD quote API via http.GET %v\n", err)
					return
				}		
				// Defer close of http.Get call
				defer response.Body.Close()		
				// Grab the entire response body.
				// Warning! ioutil.ReadAll is not memory efficient, therefore generally not recommended. 
				contents, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatalf("Error while reading response body %v\n", err)
					return
				}
				// 	Print the entire response body to terminal.
				fmt.Println(string(contents))
			},
		},
		{
			Name:    "love",
			Aliases: []string{"love"},
			Usage:   "Love quote",
			Action: func(c *cli.Context) { 
				response, err := http.Get("http://quotes.rest/qod.json?category=love")
				if err != nil {
					log.Fatalf("Error while calling QOD quote API via http.GET %v\n", err)
					return
				}		
				// Defer close of http.Get call
				defer response.Body.Close()		
				// Grab the entire response body.
				// Warning! ioutil.ReadAll is not memory efficient, therefore generally not recommended. 
				contents, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatalf("Error while reading response body %v\n", err)
					return
				}
				// 	Print the entire response body to terminal.
				fmt.Println(string(contents))
			},
		},
	}
}

