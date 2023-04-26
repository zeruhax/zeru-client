package list_ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"time"
	"zclient/pkg/helper"
	"zclient/pkg/lib"
	tabs_ui "zclient/pkg/model/tabs-ui"
)

var (
	routes = ""
)

func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(item); ok {
			title = i.Title()
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				switch title {
				case "Login":
					routes = "login"
					return tea.Quit
				case "Register":
					routes = "register"
					return tea.Quit
				case "View Profile":
					routes = "view_profile"
					return tea.Quit
				case "Change Username":
					routes = "change_username"
					return tea.Quit
				case "Change Password":
					routes = "change_password"
					return tea.Quit
				case "Add White List Ip":
					routes = "add_white_list_ip"
					return tea.Quit
				case "Add Black List Ip":
					routes = "add_black_list_ip"
					return tea.Quit
				case "Logout":
					routes = "logout"
					return tea.Quit
				}
			}
		}

		return nil
	}

	help := []key.Binding{keys.choose}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}

type delegateKeyMap struct {
	choose key.Binding
}

func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.choose,
	}
}

func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.choose,
		},
	}
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose"),
		),
	}
}

func Routes() {
	switch routes {
	case "login":
		time.Sleep(2 * time.Second)
		lib.CallClear()
		isLogin, newToken, NewRefreshToken := helper.Login("Login Form")
		if isLogin == true {
			helper.CreateSession(newToken, NewRefreshToken)
			lib.CallClear()
			DashboardMenu()
		}
	case "register":
		time.Sleep(2 * time.Second)
		lib.CallClear()
		helper.Register("Register Form")
	case "view_profile":
		time.Sleep(2 * time.Second)
		lib.CallClear()
		data := helper.ViewProfile()
		tabs := []string{
			"Username",
			"Email",
			"Apikey",
		}
		tabContent := []string{
			data.Data.Username,
			data.Data.Email,
			data.Data.Apikey,
		}
		tabs_ui.MenuTab(tabs, tabContent)
	case "change_username":
		time.Sleep(2 * time.Second)
		lib.CallClear()
		response := helper.EditUsername()
		if response == true {
			lib.CallClear()
			DashboardMenu()
		}
	case "change_password":
		time.Sleep(2 * time.Second)
		lib.CallClear()
		response := helper.EditPassword()
		if response == true {
			lib.CallClear()
			DashboardMenu()
		}
	case "add_white_list_ip":
		time.Sleep(2 * time.Second)
		lib.CallClear()
		response := helper.AddWhitelistIp()
		if response == true {
			lib.CallClear()
			DashboardMenu()
		}
	case "add_black_list_ip":
		time.Sleep(2 * time.Second)
		lib.CallClear()
		response := helper.AddBlacklistIp()
		if response == true {
			lib.CallClear()
			DashboardMenu()
		}
	case "logout":
		time.Sleep(2 * time.Second)
		helper.Logout()
	}
}
