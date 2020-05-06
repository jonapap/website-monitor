package main

import (
	"fmt"
	"os"

	"github.com/jonapap/website-watcher/internal/browser"
)

func main() {
	websites, fileNames, err := browser.GetAllWebsitesFromFiles()

	switch err.(type) {
	case nil:
	case *os.PathError:
		fmt.Println("The folder savedWebsites doesn't seem to exist.")
		return
	default:
		panic(err)
	}
	if len(websites) == 0 {
		fmt.Println("The folder savedWebsites exist but there is nothing in it.")
		return
	}

	for i, w := range websites {
		fmt.Printf("%d:\n\tFile Name: %s\n\tURL: %s\n\tSelector: %s\n", i, fileNames[i], w.URL, w.CSSSelect)
	}

	fmt.Println("\n\nPlease select the file number to delete: ")
	var num int
	_, err = fmt.Scanf("%d", &num)
	if err != nil || num < 0 || num >= len(websites) {
		fmt.Printf("Make sure you enter an interger between %d and %d (exclusive)\n", 0, len(websites))
		return
	}

	err = os.Remove("savedWebsites/" + fileNames[num])
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully removed " + websites[num].URL)
}
