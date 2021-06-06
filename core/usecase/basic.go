package usecase

import (
	"errors"
	"fmt"
	"github.com/tebeka/selenium"
	"template/core/ports"
	"time"
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

	acceptAllButton, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[2]/div/div/button[1]")
	if err == nil {
		t, err := acceptAllButton.Text()
		if err != nil {
			return err
		}
		if t == "Accept All" {
			err := acceptAllButton.Click()
			if err != nil {
				return err
			}
		}
	}
	err = wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		for i := 0; i < 20; i++ {
			_, err := wd.FindElement(selenium.ByName, "username")
			if err != nil {
				time.Sleep(1000 * time.Millisecond)
				continue
			}
			return true, nil
		}
		return false, errors.New("paged didnt respond in time")
	})
	if err != nil {
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

	err = waitIfISeeInstagramIcon(wd)
	if err != nil {
		return err
	}

	return nil
}

func (bh *basicHandler) GoToMainMenu(wd selenium.WebDriver) error {
	if err := wd.Get("https://www.instagram.com"); err != nil {
		return err
	}

	err := waitIfISeeInstagramIcon(wd)
	if err != nil {
		return err
	}
	return nil
}

func (bh *basicHandler) FollowByLink(wd selenium.WebDriver, link string) error {
	err := wd.Get(link)
	if err != nil {
		return err
	}
	err = waitIfISeeInstagramIcon(wd)
	if err != nil {
		return err
	}
	followElement, err := wd.FindElement(selenium.ByXPATH, "//*[@id=\"react-root\"]/section/main/div/header/section/div[1]/div[1]/div/div/div/span/span[1]/button")
	if err != nil {
		ter, ok := err.(*selenium.Error)
		if !ok {
			return err
		}
		if ter.Err == "no such element" {
			return errors.New("followed before")
		} else {
			return err
		}
	}

	if err := followElement.Click(); err != nil {
		return err
	}
	return nil
}

func (bh *basicHandler) UnFollowByLink(wd selenium.WebDriver, link string) error {
	err := wd.Get(link)
	if err != nil {
		return err
	}
	err = waitIfISeeInstagramIcon(wd)
	if err != nil {
		return err
	}
	unfollowElement, err := wd.FindElement(selenium.ByXPATH, "//*[@id=\"react-root\"]/section/main/div/header/section/div[1]/div[1]/div/div[2]/div/span/span[1]/button")
	if err != nil {
		ter, ok := err.(*selenium.Error)
		if !ok {
			return err
		}
		if ter.Err == "no such element" {
			return errors.New("not followed yet")
		} else {
			return err
		}
	}
	if err := unfollowElement.Click(); err != nil {
		return err
	}

	time.Sleep(500 * time.Millisecond)
	unfollowButtonElement, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[5]/div/div/div/div[3]/button[1]")
	if err != nil {
		return err
	}
	if err := unfollowButtonElement.Click(); err != nil {
		return err
	}
	return nil
}

func waitIfISeeInstagramIcon(wd selenium.WebDriver) error {
	err := wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		for i := 0; i < 20; i++ {
			_, err := wd.FindElement(selenium.ByXPATH, "//*[@id=\"react-root\"]/section/nav/div[2]/div/div/div[1]/a/div/div/img")
			if err != nil {
				e2, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[5]/div/div/div/div[3]/button[2]")
				if err != nil {
					time.Sleep(1000 * time.Millisecond)
					continue
				}
				tt, _ := e2.Text()
				if tt == "Not Now" {
					e2.Click()
				}
				time.Sleep(1000 * time.Millisecond)
				continue
			}
			return true, nil
		}
		return false, errors.New("paged didnt respond in time")
	})
	if err != nil {
		return err
	}
	return nil
}

func (bh *basicHandler) LikeByPostLink(wd selenium.WebDriver, link string) error {
	err := wd.Get(link)
	if err != nil {
		return err
	}
	err = waitIfISeeInstagramIcon(wd)
	if err != nil {
		return err
	}
	likeElement, err := wd.FindElement(selenium.ByXPATH, "//*[@id=\"react-root\"]/section/main/div/div[1]/article/div[3]/section[1]/span[1]/button/div")
	if err != nil {
		return err
	}

	a, _ := wd.ExecuteScript("var mylabel = document.querySelector(\"[aria-label=Like]\");return mylabel.ariaLabel;", nil)
	fmt.Println(a)
	//ariaLabelElement, err := wd.FindElement(selenium.ByCSSSelector,"" )
	//label, err  := ariaLabelElement.TagName()
	//ko, err  := ariaLabelElement.GetAttribute("")
	//fmt.Println(label)
	//fmt.Println(ko)
	if err == nil {
		err = likeElement.Click()
		if err != nil {
			ter, ok := err.(*selenium.Error)
			if !ok {
				return err
			}
			if ter.Err == "no such element" {
				return errors.New("post does n't load yet")
			} else {
				return err
			}
		}
	} else {
		return errors.New("ku element")
	}
	return nil
}

func (bh *basicHandler) UnlikeByPostLink(wd selenium.WebDriver, link string) error {
	err := wd.Get(link)
	if err != nil {
		return err
	}
	err = waitIfISeeInstagramIcon(wd)
	if err != nil {
		return err
	}
	likeElement, err := wd.FindElement(selenium.ByXPATH, "//*[@id=\"react-root\"]/section/main/div/div[1]/article/div[3]/section[1]/span[1]/button/div/span/svg")
	if err != nil {
		return err
	}
	isSelected, err := likeElement.IsSelected()
	if err != nil {
		return err
	}
	fmt.Println(isSelected)
	if isSelected {
		err = likeElement.Click()
		if err != nil {
			ter, ok := err.(*selenium.Error)
			if !ok {
				return err
			}
			if ter.Err == "no such element" {
				return errors.New("post does n't load yet")
			} else {
				return err
			}
		}
	} else {
		return errors.New("this post is disliked ")
	}
	return nil
}

func (bh *basicHandler) CommentByPostLink(wd selenium.WebDriver, link string, comment string) error {
	err := wd.Get(link)
	if err != nil {
		return err
	}
	err = waitIfISeeInstagramIcon(wd)
	if err != nil {
		return err
	}
	ariaLabelElement, err := wd.FindElement(selenium.ByCSSSelector,"[aria-label=Add a comment...]" )
	fmt.Println(ariaLabelElement)
	return nil
}

func (bh *basicHandler) LikeCommentByPostLink(wd selenium.WebDriver, link string, matchedComment []string, totalNum int) error {
	panic("implement me")
}

func (bh *basicHandler) SharePostByPostLink(wd selenium.WebDriver, link string, shareTo []string) error {
	panic("implement me")
}

func (bh *basicHandler) ReplyCommentByPostLink(wd selenium.WebDriver, matchedComment []string, totalNum int) error {
	panic("implement me")
}

func (bh *basicHandler) BookMarkByPostLink(wd selenium.WebDriver, link string) error {
	panic("implement me")
}

func (bh *basicHandler) SeeStories(wd selenium.WebDriver, totalNum int, perStory int) error {
	panic("implement me")
}

func (bh *basicHandler) SendDirectByUserName(wd selenium.WebDriver, username string, message string) error {
	panic("implement me")
}

func (bh *basicHandler) SendDirectByAccountLink(wd selenium.WebDriver, link string, message string) error {
	panic("implement me")
}

func (bh *basicHandler) SeenDirectMessages(wd selenium.WebDriver, totalNum int) error {
	panic("implement me")
}
