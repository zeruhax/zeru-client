package helper

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	figure2 "github.com/common-nighthawk/go-figure"
	"github.com/erikgeiser/promptkit/textinput"
	"regexp"
	"strings"
	"time"
	"zclient/pkg/core/util"
	"zclient/pkg/lib"
)

var (
	alert = lipgloss.NewStyle().Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#D21312"))
	success = lipgloss.NewStyle().Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#1B5E20"))
)

func Login(banner string) (bool, string, string) {
	var (
		token        string
		refreshToken string
	)
	figure := figure2.NewFigure(banner, "", true)
	figure.Print()
	input := textinput.New("Enter email: ")
	input.Placeholder = "Your email cannot be empty"
	input.Validate = func(email string) error {
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !re.MatchString(email) {
			return fmt.Errorf("invalid email address")
		}
		return nil
	}
	email, err := input.RunPrompt()

	if err != nil {
		fmt.Println(alert.Render(err.Error()))
	}

	input = textinput.New("Enter password: ")
	input.Placeholder = "Your password cannot be empty"
	input.Validate = func(password string) error {
		if len(password) < 6 {
			return fmt.Errorf("password must be at least 8 characters long")
		}
		return nil
	}
	input.Hidden = true
	password, err := input.RunPrompt()

	if err != nil {
		fmt.Println(alert.Render("Please try again"))
		time.Sleep(2 * time.Second)
		lib.CallClear()
		Login(banner)
	}

	var auth = util.Auth{
		Email:    email,
		Password: password,
	}

	data, err := auth.Login()

	if err != nil {
		fmt.Println(alert.Render("Please try again"))
		time.Sleep(2 * time.Second)
		lib.CallClear()
		Login(banner)
	}

	if data.Error != "" {
		fmt.Println(alert.Render(data.Error + " Please try again"))
		time.Sleep(2 * time.Second)
		lib.CallClear()
		Login(banner)
	}

	token = data.Token
	refreshToken = data.RefreshToken

	return true, token, refreshToken
}

func Register(banner string) {
	figure := figure2.NewFigure(banner, "", true)
	figure.Print()
	input := textinput.New("Enter email: ")
	input.Placeholder = "Your email cannot be empty"
	input.Validate = func(email string) error {
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !re.MatchString(email) {
			return fmt.Errorf("invalid email address")
		}
		return nil
	}
	email, err := input.RunPrompt()
	if err != nil {
		fmt.Println(alert.Render(err.Error()))
	}
	input = textinput.New("Enter Username: ")
	input.Placeholder = "Your username cannot be empty"
	username, err := input.RunPrompt()
	if err != nil {
		fmt.Println(alert.Render(err.Error()))
	}
	input = textinput.New("Enter password: ")
	input.Placeholder = "Your password cannot be empty"
	input.Validate = func(password string) error {
		if len(password) < 6 {
			return fmt.Errorf("password must be at least 6 characters long")
		}
		return nil
	}
	input.Hidden = true
	password, err := input.RunPrompt()
	if err != nil {
		fmt.Println(alert.Render("Please try again"))
		time.Sleep(2 * time.Second)
		lib.CallClear()
		Register(banner)
	}
	var auth = util.Auth{
		Username: username,
		Email:    email,
		Password: password,
	}
	data, err := auth.Register()

	if err != nil {
		fmt.Println(alert.Render("Please try again"))
		time.Sleep(2 * time.Second)
		lib.CallClear()
		Register(banner)
	}

	if data.Error != "" {
		fmt.Println(alert.Render(data.Error + " Please try again"))
		time.Sleep(2 * time.Second)
		lib.CallClear()
		Register(banner)
	}

	if data.Message != "" {
		fmt.Println(success.Render(data.Message + "Redirecting to login form"))
		time.Sleep(2 * time.Second)
		lib.CallClear()
		Login("Login Form")
	}

}

func Tools() {
	fmt.Println("1. Get All User")
}

func EditUsername() bool {
	input := textinput.New("Enter new username: ")
	input.Placeholder = "Your username cannot be empty"
	username, err := input.RunPrompt()
	if err != nil {
		fmt.Println(alert.Render(err.Error()))
	}
	data := util.EditProfile{Username: username}
	dataToken := GetToken()
	token := dataToken.Token
	refreshToken := dataToken.RefreshToken
	response, err := data.UpdateProfile(token, refreshToken)
	if err != nil {
		fmt.Println(alert.Render(err.Error()))
	}
	fmt.Println(success.Render(response.Message))
	return true
}

func EditPassword() bool {
	input := textinput.New("Enter new password: ")
	input.Placeholder = "Your password cannot be empty"
	input.Validate = func(password string) error {
		if len(password) < 6 {
			return fmt.Errorf("password must be at least 8 characters long")
		}
		return nil
	}
	password, err := input.RunPrompt()
	if err != nil {
		fmt.Println(alert.Render(err.Error()))
	}

	input = textinput.New("Enter confirm password: ")
	input.Placeholder = "Your password cannot be empty"
	input.Validate = func(confirmpassword string) error {
		if len(confirmpassword) < 6 {
			return fmt.Errorf("password must be at least 8 characters long")
		} else if confirmpassword != password {
			return fmt.Errorf("password not match")
		}
		return nil
	}
	confirmpassword, err := input.RunPrompt()
	if err != nil {
		fmt.Println(alert.Render(err.Error()))
	}
	data := util.EditProfile{Password: password, ConfirmPassword: confirmpassword}
	dataToken := GetToken()
	token := dataToken.Token
	refreshToken := dataToken.RefreshToken
	response, err := data.UpdateProfile(token, refreshToken)
	if err != nil {
		fmt.Println(alert.Render(err.Error()))
	}
	fmt.Println(success.Render(response.Message))
	return true
}

func AddWhitelistIp() bool {
	input := textinput.New("Add whitelist ip: ")
	input.Placeholder = "You can add mass ip with seperate (,) Example: 127.0.0.1,0.0.0.0"
	ip, err := input.RunPrompt()
	if err != nil {
		fmt.Println(alert.Render(err.Error()))
	}
	ips := strings.Split(ip, ",")
	data := util.EditProfile{WhitelistIp: ips}
	dataToken := GetToken()
	token := dataToken.Token
	refreshToken := dataToken.RefreshToken
	response, err := data.AddIp(token, refreshToken)
	if err != nil {
		fmt.Println(alert.Render(err.Error()))
	}
	fmt.Println(success.Render(response.Message))
	return true
}

func AddBlacklistIp() bool {
	input := textinput.New("Add blacklist ip: ")
	input.Placeholder = "You can add mass ip with seperate (,) Example: 127.0.0.1,0.0.0.0"
	ip, err := input.RunPrompt()
	if err != nil {
		fmt.Println(alert.Render(err.Error()))
	}
	ips := strings.Split(ip, ",")
	data := util.EditProfile{BlacklistIp: ips}
	dataToken := GetToken()
	token := dataToken.Token
	refreshToken := dataToken.RefreshToken
	response, err := data.AddIp(token, refreshToken)
	if err != nil {
		fmt.Println(alert.Render(err.Error()))
	}
	fmt.Println(success.Render(response.Message))
	return true
}

func ViewProfile() *util.ResponseProfile {
	dataToken := GetToken()
	token := dataToken.Token
	refreshToken := dataToken.RefreshToken
	var store = new(util.ResponseProfile)
	data, err := store.GetProfile(token, refreshToken)
	if err != nil {
		fmt.Println(alert.Render(err.Error()))
	}
	return data
}
