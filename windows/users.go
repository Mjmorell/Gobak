package windows

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"

	humanize "github.com/dustin/go-humanize"
	d "github.com/mjmorell/gobak/driveutils"
	p "github.com/mjmorell/gobak/strutils"
	u "github.com/mjmorell/gobak/userutils"
)

func GetUsers(src d.Drive) (users u.UserList) {
	entries, err := ioutil.ReadDir(src.Path)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		tmpsrc := filepath.Join(src.Path, entry.Name())
		si, _ := os.Stat(tmpsrc)
		if !si.IsDir() {
			continue
		}
		switch entryL := strings.ToLower(entry.Name()); entryL {
		case "public":
			continue
		case "all users":
			continue
		case "default":
		case "admini~1":
			continue
		case "default user":
			continue
		default:
			var tempUser u.UserProfile
			tempUser.Username = entry.Name()
			tempUser.AbsolutePath = tmpsrc

			users.AllUsers = append(users.AllUsers, tempUser)
		}
	}

	return
}

func SetupUsers(users *u.UserList) {
	for k, v := range users.AllUsers {

		files, err := ioutil.ReadDir(v.AbsolutePath)
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			if f.IsDir() && strings.ToLower(f.Name()) != "appdata" {
				var temp u.UserRootFolder
				temp.Folder = f.Name()
				temp.AbsolutePath = v.AbsolutePath + "\\" + f.Name()
				temp.Size = DirSize(temp.AbsolutePath)
				//temp.Percentage = math.Round(float64(temp.Size)/float64(v.SizeParanoid)*100) / 100
				users.AllUsers[k].Folders = append(users.AllUsers[k].Folders, temp)
			} else if strings.ToLower(f.Name()) != "appdata" {
				var temp u.UserRootFile
				temp.Filename = f.Name()
				temp.AbsolutePath = v.AbsolutePath + "\\" + f.Name()
				temp.Size = uint64(f.Size())
				users.AllUsers[k].Files = append(users.AllUsers[k].Files, temp)
			}
		}
	}
}

func SizeUsers(users *u.UserList) {
	Paranoia = true
	for k := range users.AllUsers {
		// sizes
		for kf, eachFolder := range users.AllUsers[k].Folders {
			if regexableFolder(eachFolder.Folder) {
				continue
			}
			users.AllUsers[k].Folders[kf].Mode = 1
			users.AllUsers[k].SizeParanoid = users.AllUsers[k].SizeParanoid + eachFolder.Size
		}
		for kf, eachFile := range users.AllUsers[k].Files {
			if regexableFile(eachFile.Filename) {
				users.AllUsers[k].Files[kf].Mode = 3
				continue
			}
			users.AllUsers[k].Files[kf].Mode = 1
			users.AllUsers[k].SizeParanoid = users.AllUsers[k].SizeParanoid + eachFile.Size
		}
		// percentages
		for kf, eachFolder := range users.AllUsers[k].Folders {
			if regexableFolder(eachFolder.Folder) {
				continue
			}
			if eachFolder.Size == 0 {
				users.AllUsers[k].Folders[kf].PercentageP = 0
				continue
			}
			users.AllUsers[k].Folders[kf].PercentageP = math.Round(float64(eachFolder.Size)/float64(users.AllUsers[k].SizeParanoid)*10000.) / 100.
		}
		for kf, eachFile := range users.AllUsers[k].Files {
			if regexableFile(eachFile.Filename) {
				continue
			}
			if eachFile.Size == 0 {
				users.AllUsers[k].Files[kf].PercentageP = 0
				continue
			}
			users.AllUsers[k].Files[kf].PercentageP = math.Round(float64(eachFile.Size)/float64(users.AllUsers[k].SizeParanoid)*10000.) / 100.
		}

	}

	////
	////

	Paranoia = false
	for k := range users.AllUsers {
		// sizes
		for kf, eachFolder := range users.AllUsers[k].Folders {
			if regexableFolder(eachFolder.Folder) {
				continue
			}
			users.AllUsers[k].Folders[kf].Mode = 0
			users.AllUsers[k].SizeNormal = users.AllUsers[k].SizeNormal + eachFolder.Size
		}
		for kf, eachFile := range users.AllUsers[k].Files {
			if regexableFile(eachFile.Filename) {
				users.AllUsers[k].Files[kf].Mode = 3
				continue
			}
			users.AllUsers[k].Files[kf].Mode = 0
			users.AllUsers[k].SizeNormal = users.AllUsers[k].SizeNormal + eachFile.Size
		}
		// percentages
		for kf, eachFolder := range users.AllUsers[k].Folders {
			if regexableFolder(eachFolder.Folder) {
				continue
			}
			if eachFolder.Size == 0 {
				users.AllUsers[k].Folders[kf].Percentage = 0
				continue
			}
			users.AllUsers[k].Folders[kf].Percentage = math.Round(float64(eachFolder.Size)/float64(users.AllUsers[k].SizeNormal)*10000.) / 100.
		}
		for kf, eachFile := range users.AllUsers[k].Files {
			if regexableFile(eachFile.Filename) {
				continue
			}
			if eachFile.Size == 0 {
				users.AllUsers[k].Files[kf].Percentage = 0
				continue
			}
			users.AllUsers[k].Files[kf].Percentage = math.Round(float64(eachFile.Size)/float64(users.AllUsers[k].SizeNormal)*10000.) / 100.
		}
	}
}

func WhatUsers(users u.UserList) u.UserList {
	p.Header()

	userStrings := []string{}
	userInformation := []string{}
	if Paranoia {
		//fmt.Printf(p.Green("    USERNAME ") + "           | %%total | " + p.Yellow("               Size       ") + "| Estimated" + "\n")
		//fmt.Println("────────────────────────────────────────────────────────────────────────────────────────────────────")
		for _, v := range users.AllUsers {
			fmt.Printf("--\n")
			userStrings = append(userStrings, fmt.Sprintf("%-28s ", v.Username))
			userInformation = append(userInformation, fmt.Sprintf("| %05.2f%% | %28s bytes | %s", v.PercentageP, p.Yellow(humanize.Comma(int64(v.SizeParanoid))), printSize(v.SizeParanoid)))
		}
	} else {
		//fmt.Printf(p.Green("    USERNAME ") + "           | %%total | " + p.Yellow("               Size       ") + "| Estimated" + "\n")
		//fmt.Println("────────────────────────────────────────────────────────────────────────────────────────────────────")
		for _, v := range users.AllUsers {
			fmt.Printf("--\n")
			userStrings = append(userStrings, fmt.Sprintf("%-28s ", v.Username))
			userInformation = append(userInformation, fmt.Sprintf("| %05.2f%% | %28s bytes | %s", v.Percentage, p.Yellow(humanize.Comma(int64(v.SizeNormal))), printSize(v.SizeNormal)))
		}
	}

	p.Flush()

	choices := p.QChoiceMultiple("What users would you like to back up?", userStrings, userInformation)

	var usersToBackup u.UserList
	for k := range choices {
		choices[k] = strings.Trim(choices[k], " ")
		for _, v := range users.AllUsers {
			if v.Username == choices[k] {
				usersToBackup.AllUsers = append(usersToBackup.AllUsers, v)
			}
		}
	}
	return usersToBackup
}

func printSize(size uint64) (str string) {
	if size > 2*1024 { //kb
		if size > 2*1024*1024 { //mb
			if size > 2*1024*1024*1024 { //gb
				if size > 2*1024*1024*1024*1024 { //tb
					return fmt.Sprintf("%.2f TB", float64(size)/(1024.0*1024.0*1024.0*1024.0))
				}
				return fmt.Sprintf("%.2f GB", float64(size)/(1024.0*1024.0*1024.0))
			}
			return fmt.Sprintf("%.2f MB", float64(size)/(1024.0*1024.0))
		}
		return fmt.Sprintf("%.2f KB", float64(size)/(1024.0))
	}
	return fmt.Sprintf("%d Bytes", size)
}
