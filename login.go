package main

import (
	"fmt"
	"os"

	p "github.com/mjmorell/gobak/strutils"

	goserve "github.com/mjmorell/GoServe"
)

func agreement() {
	p.Header()
	p.CenterString(p.Green("Please note, ")+p.Red("this program is not 100%% guaranteed."), 0, 49)
	p.CenterString("This program is built as-is, and is not certified by LU.", 0, 56)
	p.CenterString("Any questions can be directed to "+p.Green("mjmorell@liberty.edu"), 0, 53)

	p.CenterString(p.Red("──────────────────────────────"), 0, 30)

	p.CenterString("All sizing estimates are based off of division with "+p.Green("1024")+" as bytes", 0, 65)
	p.CenterString("These sizing estimates are based off the binary system!! Such as:", 0, 65)
	p.CenterString("A file sized at 10 Kilobytes is 10*1024 Bytes. Not 10*1000 Bytes.", 0, 65)
	fmt.Println()
	p.CenterString("Therefore, size estimates post-backup may be different depending", 0, 65)
	p.CenterString("on which Operating System the sizing estimate is made from.", 0, 59)
	p.Wait(p.Cyan("Press") + p.Magenta(" [Enter] ") + p.Cyan("to acknowledge"))
}

func login() {
	goserveClient.Username = p.QGeneric("  What is your username?")
	goserveClient.Password = p.QPassword("  What is your password, " + goserveClient.Username + "?")

	if !testLogin(goserveClient) {
		goserveClient.Password = p.QPassword("  Error! What is your password, " + goserveClient.Username + "?")
		if !testLogin(goserveClient) {
			goserveClient.Password = p.QPassword("  Error! What is your password, " + goserveClient.Username + "?")
			if !testLogin(goserveClient) {
				if !p.QYesNo("  Login Error. Continue in Offline Mode?") {
					p.Header()
					fmt.Println("  Ending Program Execution")
					fmt.Println("  Check Internet Connection")
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
