package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jonapap/website-monitor/internal/browser"
	"github.com/jonapap/website-monitor/internal/mail"
)

func main() {
	websites, _, err := browser.GetAllWebsitesFromFiles()

	var e *os.PathError
	if errors.As(err, &e) {
		fmt.Println("The folder savedWebsites doesn't seem to exist. Please run websitesaver.exe beforehand.")
		return
	} else if err != nil {
		panic(err)
	}

	if len(websites) == 0 {
		fmt.Println("The folder savedWebsites exist but there is nothing in it. Please run websitesaver.exe beforehand.")
		return
	}

	b, err := browser.NewBrowser()
	if err != nil {
		panic(err)
	}
	defer b.Close()

	modified := false //Tracks if at least one website has changed
	notification := mail.Message{}
	notification.SetSubject("Website Watcher : A watched website has changed")
	notification.AddLineToBody("Here is the list of websites that have changed since their last save:\n")
	for _, w := range websites {
		if err = b.NavigateTo(w.URL); err != nil {
			fmt.Printf("Error navigating to %s", w.URL)
			continue
		}

		source, err := b.GetSource(w.CSSSelect)
		if err != nil {
			fmt.Printf("Error getting the source of %s", w.URL)
			continue
		}

		if w == source {
			fmt.Printf("%s didn't change since last save\n", w.URL)
		} else {
			fmt.Printf("%s was modified!\n", w.URL)
			notification.AddLineToBody(fmt.Sprintf("URL: %s Selector: %s", w.URL, w.CSSSelect))
			modified = true
		}
	}

	if modified {
		err := notification.Send()

		var e mail.ConfigFileDidNotExistError
		if errors.As(err, &e) {
			fmt.Println(err)
		} else if err != nil {
			panic(err)
		}
	}
}
