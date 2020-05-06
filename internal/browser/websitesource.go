package browser

import (
	b64 "encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
)

//WebsiteSource represents the source code of a website
type WebsiteSource struct {
	URL    string
	Source string
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
