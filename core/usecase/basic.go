package usecase

import (
	"github.com/tebeka/selenium"
	"template/core/ports"
)

type basicHandler struct {
}

func NewBasicService() ports.BasicsService {
	return &basicHandler{}
}

func (bh *basicHandler) Login(wd selenium.WebDriver, username string, password string) error {
	if err := wd.Get("https://www.instagram.com/accounts/login/"); err != nil {
		return err
	}
	userNameLoginElement, err := wd.FindElement(selenium.ByName, "username")
	if err != nil {
		return err
	}
	err = userNameLoginElement.SendKeys(username)
	if err != nil {
		return err
	}
	passwordLoginElement, err := wd.FindElement(selenium.ByName, "password")
	if err != nil {
		return err
	}
	err = passwordLoginElement.SendKeys(password)
	if err != nil {
		return err
	}
	loginButtonElement, err := wd.FindElement(selenium.ByXPATH, "//*[@id=\"loginForm\"]/div/div[3]/button/div")
	if err != nil {
		return err
	}
	err = loginButtonElement.Click()
	if err != nil {
		return err
	}
	return nil
}
