package browser

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//WebsiteSource represents the source code of a website
type WebsiteSource struct {
	URL       string
	Source    string
	CSSSelect string
}

//WriteToFile writes this WebsiteSource to a file under the folder savedWebsites
func (w *WebsiteSource) WriteToFile() error {
	sum := md5.Sum([]byte(w.URL + w.CSSSelect))
	fileName := hex.EncodeToString(sum[:]) //Try to generate a file name of fixed length

	sourceJSON, err := json.Marshal(*w)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}

	if err := os.MkdirAll("savedWebsites", 0644); err != nil {
		return fmt.Errorf("Error writing to file: %w", err)
	}

	return ioutil.WriteFile("savedWebsites/"+fileName+".json", sourceJSON, 0644)
}

//GetAllWebsitesFromFiles reads the folder savedWebsites and returns all the websites
//in there. First argument returned is the websites themselves, and the second is the name
//of the file themselves.
func GetAllWebsitesFromFiles() ([]WebsiteSource, []string, error) {
	websites, err := ioutil.ReadDir("savedWebsites/")
	if err != nil {
		return nil, nil, fmt.Errorf("Error reading directory savedWebsites: %w", err)
	}

	sources := []WebsiteSource{}
	fileNames := []string{}

	for _, f := range websites {
		if !strings.HasSuffix(f.Name(), ".json") {
			continue //Ignore files that are not of type json
		}

		bytes, err := ioutil.ReadFile("savedWebsites/" + f.Name())
		if err != nil {
			return nil, nil, fmt.Errorf("Error reading file %s: %w", f.Name(), err)
		}
		var dat map[string]interface{}

		if err := json.Unmarshal(bytes, &dat); err != nil {
			return nil, nil, fmt.Errorf("Error reading file %s: %v", f.Name(), err)
		}

		sources = append(sources, WebsiteSource{dat["URL"].(string), dat["Source"].(string), dat["CSSSelect"].(string)})
		fileNames = append(fileNames, f.Name())
	}
	return sources, fileNames, nil
}
