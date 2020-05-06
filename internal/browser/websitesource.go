package browser

import (
	b64 "encoding/base64"
	"encoding/json"
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
	fileName := b64.StdEncoding.EncodeToString([]byte(w.URL))

	sourceJSON, err := json.Marshal(*w)

	if err != nil {
		return err
	}

	if err := os.MkdirAll("savedWebsites", 0644); err != nil {
		return err
	}

	return ioutil.WriteFile("savedWebsites/"+fileName+".json", sourceJSON, 0644)
}

//GetAllWebsitesFromFiles reads the folder savedWebsites and returns all the websites
//in there
func GetAllWebsitesFromFiles() ([]WebsiteSource, error) {
	websites, err := ioutil.ReadDir("savedWebsites/")
	if err != nil {
		return nil, err
	}

	sources := []WebsiteSource{}

	for _, f := range websites {
		if !strings.HasSuffix(f.Name(), ".json") {
			continue //Ignore files that are not of type json
		}

		bytes, err := ioutil.ReadFile("savedWebsites/" + f.Name())
		if err != nil {
			return nil, err
		}
		var dat map[string]interface{}

		if err := json.Unmarshal(bytes, &dat); err != nil {
			return nil, err
		}

		sources = append(sources, WebsiteSource{dat["URL"].(string), dat["Source"].(string), dat["CSSSelect"].(string)})
	}

	return sources, nil
}
