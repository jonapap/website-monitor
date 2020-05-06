package main

import (
	"fmt"
	"os"

	"github.com/jonapap/website-watcher/internal/browser"
)

func main() {
	websites, _, err := browser.GetAllWebsitesFromFiles()

	switch err.(type) {
	case nil:
	case *os.PathError:
		fmt.Println("The folder savedWebsites doesn't seem to exist. Please run websitesaver.exe beforehand.")
		return
	default:
		panic(err)
	}
	if len(websites) == 0 {
		fmt.Println("The folder savedWebsites exist but there is nothing in it. Please run websitesaver.exe beforehand.")
	}

	b, err := browser.NewBrowser()
	if err != nil {
		panic(err)
	}
	defer b.Close()

	for _, w := range websites {
		if err = b.NavigateTo(w.URL); err != nil {
			panic(err)
		}

		source, err := b.GetSource(w.CSSSelect)
		if err != nil {
			panic(err)
		}

		if w == source {
			fmt.Printf("%s didn't change since last save\n", w.URL)
		} else {
			fmt.Printf("%s was modified!\n", w.URL)
		}
	}
}
