package main

import (
	"github.com/joho/godotenv"
	"log"
	"zclient/pkg/helper"
	list_ui "zclient/pkg/model/list-ui"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	isLogin, _ := helper.CheckSession()
	if isLogin == true {
		list_ui.DashboardMenu()
	} else {
		list_ui.ScreenMenu()
	}
}
