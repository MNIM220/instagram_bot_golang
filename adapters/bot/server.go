package bot

import (
	"fmt"
	"template/core/usecase"
)

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
	err = basicService.LikeByPostLink(wd,"https://www.instagram.com/p/CPq8gVHHbML/?utm_source=ig_web_copy_link")
	if err != nil {
		panic(err)
	}
	err = basicService.FollowByLink(wd, "https://www.instagram.com/neginzare.life/")
	if err != nil {
		panic(err)
	}

	err = basicService.GoToMainMenu(wd)
	if err != nil {
		panic(err)
	}

	err = basicService.UnFollowByLink(wd, "https://www.instagram.com/neginzare.life/")
	if err != nil {
		panic(err)
	}

	err = basicService.GoToMainMenu(wd)
	if err != nil {
		panic(err)
	}
	fmt.Println("Everything is OK")
}