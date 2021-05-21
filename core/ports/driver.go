package ports

import "github.com/tebeka/selenium"

type HelloService interface {
	SayHello()
}

type BasicsService interface {
	Login(wd selenium.WebDriver, username string, password string) error
}