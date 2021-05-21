package bot

import "template/core/usecase"

func StartBot() {
	basicService := usecase.NewBasicService()
	wd, seleniumService := newSeleniumWebDriver()
	defer seleniumService.Stop()
	defer wd.Close()
	basicService.Login(wd, "csgobga5", "csgobga1234")
}