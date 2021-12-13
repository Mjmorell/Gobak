package main

import (
	"fmt"
	cm "gobak/consolemanagement"
)

// /************************************************
// 	MIT License
// 	Details viewable in the Github Directory
// 	Copyright (c) 2021 Michael Morell
// ************************************************/

func main() {

	fmt.Printf("")
	cm.Agreement()

	cm.Header()
}

// 	// --------------
// 	// WINDOWS BACKUP
// 	// --------------

// 	switch backupInfo.OS {
// 	case "WIN":
// 		src, dst := w.DiskSelection()
// 		backupInfo.Source = d.Drive{Path: src}
// 		backupInfo.Destination = d.Drive{Path: dst}
// 		if p.DevMode != 0 {
// 			backupInfo.Destination.Path += "\\_TEST-FOLDERS\\"
// 		}
// 		backupInfo.Destination.Path = backupInfo.Destination.Path + backupInfo.ID
// 		p.Flush()

// 		var users u.UserList
// 		users = w.GetUsers(backupInfo.Source)
// 		w.SetupUsers(&users)
// 		w.SizeUsers(&users)
// 		p.Flush()

// 		//users.LogPrintoutAllUsers()
// 		users = users.SetSize()

// 		w.ParanoiaRequestWindows(users)

// 		usersToBackup := w.WhatUsers(users)

// 		//for _, v := range usersToBackup.AllUsers {
// 		//	fmt.Println(v)
// 		//}

// 		if _, err := os.Stat(backupInfo.Destination.Path); !os.IsNotExist(err) {
// 			for i := 1; i < 25; i++ {
// 				if _, err := os.Stat(backupInfo.Destination.Path + "-(" + strconv.Itoa(i) + ")"); os.IsNotExist(err) {
// 					backupInfo.Destination.Path += "-(" + strconv.Itoa(i) + ")"
// 					os.Mkdir(backupInfo.Destination.Path, 0666)
// 					break
// 				}
// 			}
// 		} else {
// 			os.Mkdir(backupInfo.Destination.Path, 0666)
// 		}

// 		for _, v := range usersToBackup.AllUsers {
// 			fmt.Println(p.Green("Backing up " + v.Username))
// 			w.CopyUser(backupInfo, v)
// 		}

// 	// --------------
// 	//  MacOS BACKUP
// 	// --------------

// 	case "MAC":
// 		//MacOS Backup / Setup
// 	}
// }
