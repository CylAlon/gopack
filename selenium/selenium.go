package selenium

import (
	"github.com/tebeka/selenium"
	"fmt"
)


const (
	Defpath   = "C:/Program Files/Google/Chrome/Application/chromedriver.exe"
	Defport   = 9090
	Defbrower = "chrome"
)

type selenium_s struct {
	Server *selenium.Service
	Win    selenium.WebDriver
}

func init() {
}

func (s *selenium_s) NewBrowser(path string, port int, name string) (*selenium.Service, selenium.WebDriver) { //(*selenium.Service,selenium.WebDriver)
	ops := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(path, port, ops...)
	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v\n", err)
	}
	caps := selenium.Capabilities{
		"browserName": name,
	}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		fmt.Printf("Failed to call browser: %v\n", err)
	}
	s.Server = service
	s.Win = wd
	return service, wd
}

func New(path string, port int, name string) *selenium_s {
	val := selenium_s{}
	val.NewBrowser(path, port, name)
	return &val
}
func NewDef() *selenium_s {
	val := New(Defpath, Defport, Defbrower)
	return val
}

func NewSel(path string, port int, name string) (*selenium.Service, selenium.WebDriver) {
	val := selenium_s{}
	s, w := val.NewBrowser(path, port, name)
	return s, w
}
func NewDefSell() (*selenium.Service, selenium.WebDriver) {
	val := selenium_s{}
	s, w := val.NewBrowser(Defpath, Defport, Defbrower)
	return s, w
}
