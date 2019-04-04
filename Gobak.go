package main

import (
	"fmt"

	goserve "github.com/mjmorell/GoServe"
	d "github.com/mjmorell/gobak/driveutils"
	p "github.com/mjmorell/gobak/strutils"
	u "github.com/mjmorell/gobak/userutils"
	w "github.com/mjmorell/gobak/windows"
)

var (
	backupInfo    d.BackupInformation
	goserveClient goserve.Client
)

func main() {
	login()
	identifier()

	p.Header()

	// --------------
	// WINDOWS BACKUP
	// --------------

	switch backupInfo.OS {
	case "WIN":
		w.MapDrive(goserveClient.Username, goserveClient.Password)
		src, dst := w.DiskSelection()
		source := d.Drive{Path: src}
		destination := d.Drive{Path: dst}
		destination = destination
		p.Flush()

		var users u.UserList
		users = w.GetUsers(source)
		w.SetupUsers(&users)
		w.SizeUsers(&users)
		p.Flush()

		//users.LogPrintoutAllUsers()
		users = users.SetSize()

		w.Paranoia = w.ParanoiaRequestWindows(users)

		if w.Paranoia {
			fmt.Println("Yes!")
		} else {
			fmt.Println("NO!")
		}
	// --------------
	//  MacOS BACKUP
	// --------------

	case "MAC":
		//MacOS Backup / Setup
	}
}
