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
			Aliases: []string{"qd"},
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
					log.Fatalf("Error while calling Inspire quote API via http.GET %v\n", err)
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
			Aliases: []string{"lf"},
			Usage:   "Life quote",
			Action: func(c *cli.Context) { 
				response, err := http.Get("http://quotes.rest/qod.json?category=life")
				if err != nil {
					log.Fatalf("Error while calling Life quote API via http.GET %v\n", err)
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
			Name:    "funny",
			Aliases: []string{"fny"},
			Usage:   "Funny quote",
			Action: func(c *cli.Context) { 
				response, err := http.Get("http://quotes.rest/qod.json?category=funny")
				if err != nil {
					log.Fatalf("Error while calling Funny quote API via http.GET %v\n", err)
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
			Aliases: []string{"mng"},
			Usage:   "Management quote",
			Action: func(c *cli.Context) { 
				response, err := http.Get("http://quotes.rest/qod.json?category=management")
				if err != nil {
					log.Fatalf("Error while calling Management quote API via http.GET %v\n", err)
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
			Aliases: []string{"lv"},
			Usage:   "Love quote",
			Action: func(c *cli.Context) { 
				response, err := http.Get("http://quotes.rest/qod.json?category=love")
				if err != nil {
					log.Fatalf("Error while calling Love quote API via http.GET %v\n", err)
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
			Name:    "sports",
			Aliases: []string{"spt"},
			Usage:   "Sport quote",
			Action: func(c *cli.Context) { 
				response, err := http.Get("http://quotes.rest/qod.json?category=sports")
				if err != nil {
					log.Fatalf("Error while calling Sports quote API via http.GET %v\n", err)
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
			Name:    "art",
			Aliases: []string{"at"},
			Usage:   "Art quote",
			Action: func(c *cli.Context) { 
				response, err := http.Get("http://quotes.rest/qod.json?category=art")
				if err != nil {
					log.Fatalf("Error while calling Art quote API via http.GET %v\n", err)
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
			Name:    "students",
			Aliases: []string{"sdn", "std", "stdn"},
			Usage:   "Students quote",
			Action: func(c *cli.Context) { 
				response, err := http.Get("http://quotes.rest/qod.json?category=students")
				if err != nil {
					log.Fatalf("Error while calling Students quote API via http.GET %v\n", err)
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

