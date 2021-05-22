package ports

import "github.com/tebeka/selenium"

type HelloService interface {
	SayHello()
}

type BasicsService interface {
	Login(wd selenium.WebDriver, username string, password string) error
	GoToMainMenu(wd selenium.WebDriver) error
	FollowByLink(wd selenium.WebDriver, link string) error
	UnFollowByLink(wd selenium.WebDriver, link string) error
}
