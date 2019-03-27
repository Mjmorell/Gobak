package main

import (
	"fmt"

	goserve "github.com/mjmorell/GoServe"
	d "github.com/mjmorell/gobak/driveutils"
	p "github.com/mjmorell/gobak/strutils"
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

	if backupInfo.OS == "WIN" {
		w.MapDrive(goserveClient.Username, goserveClient.Password)
		src, dst := w.DiskSelection()
		source := d.Drive{Path: src}
		destination := d.Drive{Path: dst}
		destination = destination

		users := w.GetUsers(source)
		w.SetupUsers(&users)
		w.SizeUsers(&users)

		temp := users.GetEasyPrint()
		for _, v := range temp {
			fmt.Println(v)
		}
	} else if backupInfo.OS == "MAC" {
		//MacOS Backup / Setup
	}
}
