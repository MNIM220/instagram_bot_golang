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
	LikeByPostLink(wd selenium.WebDriver, link string) error
	UnlikeByPostLink(wd selenium.WebDriver, link string) error
	CommentByPostLink(wd selenium.WebDriver, link string, comment string) error
	LikeCommentByPostLink(wd selenium.WebDriver, link string, matchedComment []string, totalNum int) error
	SharePostByPostLink(wd selenium.WebDriver, link string, shareTo []string) error
	ReplyCommentByPostLink(wd selenium.WebDriver, matchedComment []string, totalNum int) error
	BookMarkByPostLink(wd selenium.WebDriver, link string) error
	SeeStories(wd selenium.WebDriver, totalNum int, perStory int) error
	SendDirectByUserName(wd selenium.WebDriver, username string, message string) error
	SendDirectByAccountLink(wd selenium.WebDriver, link string, message string) error
	//TODO: its a xpath to search into "//*[@id="react-root"]/section/div/div[2]/div/div/div[1]/div[2]/div/div/div/div/div[1]/a/div/div[3]"
	//																													   *            *
	SeenDirectMessages(wd selenium.WebDriver, totalNum int) error
	// TODO: figure out how to manage posting media and story
}

