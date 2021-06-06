package bot

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"os"
)

func newSeleniumWebDriver() (selenium.WebDriver, *selenium.Service) {
	const (
		// These paths will be different on your system.
		seleniumPath     = "selenium-server-standalone-3.141.59.jar"
		chromeDriverPath = "chromedriver"
		port             = 8080
	)
	opts := []selenium.ServiceOption{
		//selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.ChromeDriver(chromeDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),              // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}

	caps.AddChrome(chrome.Capabilities{
		MobileEmulation:  &chrome.MobileEmulation{
			DeviceName:    "iPhone X",
		},
	})

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	return wd, service
}
