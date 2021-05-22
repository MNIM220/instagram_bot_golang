package bot

import "template/core/usecase"

func StartBot() {
	basicService := usecase.NewBasicService()
	wd, seleniumService := newSeleniumWebDriver()
	defer seleniumService.Stop()
	defer wd.Close()
	err := basicService.Login(wd, "csgobga5", "csgobga1234")
	if err != nil {
		panic(err)
	}
	err = basicService.GoToMainMenu(wd)
	if err != nil {
		panic(err)
	}
}