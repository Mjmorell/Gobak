package main

import (
	"fmt"
	"os"

	p "github.com/mjmorell/gobak/strutils"

	goserve "github.com/mjmorell/GoServe"
)

func login() {
	goserveClient.Username = p.QGeneric("What is your username?")
	goserveClient.Password = p.QPassword("What is your password, " + goserveClient.Username + "?")

	if !testLogin(goserveClient) {
		goserveClient.Password = p.QPassword("Error! What is your password, " + goserveClient.Username + "?")
		if !testLogin(goserveClient) {
			goserveClient.Password = p.QPassword("Error! What is your password, " + goserveClient.Username + "?")
			if !testLogin(goserveClient) {
				if !p.QYesNo("Login Error. Continue in Offline Mode?") {
					p.Header()
					fmt.Println("Ending Program Execution")
					fmt.Println("Check Internet Connection")
					fmt.Println()
					os.Exit(0)
				}
				p.OffMode = true
			}
		}
	}

}

func testLogin(Client goserve.Client) bool {
	var csticket []goserve.User
	csticket, err := Client.FilterUsers("sys_user", "email="+Client.Username+"@liberty.edu")
	if err != nil {
		return false
	} else if len(csticket) != 0 {
		return true
	} else {
		return false
	}

}
