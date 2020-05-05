package browser

//WebsiteSource represents the source code of a website
type WebsiteSource struct {
	url    string
	source string
}

//GetURL returns the url of the webpage
func (w *WebsiteSource) GetURL() string {
	return w.url
}

//GetSource returns the source of the webpage
func (w *WebsiteSource) GetSource() string {
	return w.source
}
