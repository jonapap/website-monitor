package browser

import (
	"fmt"

	"github.com/tebeka/selenium"
)

//Browser simulates a simple web browser using Selenium
type Browser struct {
	service *selenium.Service
	selenium.WebDriver
}

//NavigateTo navigates to the specified url
func (b *Browser) NavigateTo(url string) error {
	return b.Get(url)
}

//GetSource returns the source code of the current webpage
//If xpath is empty, the whole page is returned. If it is specified, the source
//code of the element pointed at by the xpath will be returned.
func (b *Browser) GetSource(cssSelect string) (WebsiteSource, error) {
	if cssSelect == "" {
		sr, err := b.PageSource()
		if err != nil {
			return WebsiteSource{}, err
		}
		url, err := b.CurrentURL()
		return WebsiteSource{url, sr, cssSelect}, nil
	}

	elem, err := b.FindElement(selenium.ByCSSSelector, cssSelect)
	if err != nil {
		panic(err)
	}

	src, err := b.ExecuteScript("return arguments[0].outerHTML;", []interface{}{elem})
	if err != nil {
		panic(err)
	}

	url, err := b.CurrentURL()
	return WebsiteSource{url, src.(string), cssSelect}, nil
}

//Close will clean up and close the Browser. Must be called when the program us done using the Browser.
func (b *Browser) Close() {
	b.Quit()
	b.service.Stop()
}

//NewBrowser returns a new Browser object. In the background, it initializes a new Selenium service
//and uses Firefox as the browser.
func NewBrowser() (*Browser, error) {
	const (
		seleniumPath    = "libs/selenium-server-standalone-3.141.59.jar"
		geckoDriverPath = "libs/geckodriver.exe"
		port            = 8081
	)
	opts := []selenium.ServiceOption{
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		//selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	//selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		return nil, err
	}

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		return nil, err
	}

	return &Browser{service, wd}, nil
}
