package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jonapap/website-monitor/internal/browser"
)

func main() {
	websiteURL := flag.String("website", "", "URL of website to load. Include the protocol (ex: http://)")
	cssFlag := flag.String("selector", "", "Optional. CSS selector of a specific element to save.")
	flag.Parse()

	if *websiteURL == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Which website do you want to save? Please include the protocol (ex: http://):")
		*websiteURL, _ = reader.ReadString('\n')
		*websiteURL = strings.Replace(*websiteURL, "\n", "", -1)

		fmt.Println("Please enter a CSS selector if you want to save part of the webpage. If you want to save the full page, enter nothing:")
		*cssFlag, _ = reader.ReadString('\n')
		*cssFlag = strings.Replace(*cssFlag, "\n", "", -1)
		*cssFlag = strings.Replace(*cssFlag, "\r", "", -1)
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

	source, err := b.GetSource(*cssFlag)
	if err != nil {
		panic(err)
	}

	if err := source.WriteToFile(); err != nil {
		panic(err)
	}
}
