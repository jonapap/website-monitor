package main

import (
	"flag"
	"fmt"

	"github.com/jonapap/website-watcher/internal/browser"
)

func main() {
	websiteURL := flag.String("website", "", "URL of website to load. Include the protocol (ex: http://)")
	flag.Parse()

	if *websiteURL == "" {
		fmt.Println("Specify the website to load using -website")
		return
	}

	b, err := browser.NewBrowser()
	if err != nil {
		panic(err)
	}
	defer b.Close()

	fmt.Println("Loading " + *websiteURL)

	err = b.NavigateTo(*websiteURL)
	if err != nil {
		fmt.Println("Error while navigating to page. Make sure this is a valid URL!")
		return
	}

	source, err := b.GetSource()
	if err != nil {
		panic(err)
	}

	if err := source.WriteToFile(); err != nil {
		panic(err)
	}
}
