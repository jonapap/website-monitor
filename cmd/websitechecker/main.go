package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jonapap/website-watcher/internal/browser"
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
		}
	}
}
