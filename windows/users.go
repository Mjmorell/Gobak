package windows

import (
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"

	d "github.com/mjmorell/gobak/driveutils"
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
	for _, v := range users.AllUsers {

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
				v.Folders = append(v.Folders, temp)
			} else {
				var temp u.UserRootFile
				temp.Filename = f.Name()
				temp.AbsolutePath = v.AbsolutePath + "\\" + f.Name()
				temp.Size = uint64(f.Size())
				v.Files = append(v.Files, temp)
			}
		}
	}
}

func SizeUsers(users *u.UserList) {
	Paranoia = false
	ExParanoia = false
	for _, v := range users.AllUsers {
		// sizes
		for _, eachFolder := range v.Folders {
			if regexableFolder(eachFolder.Folder) {
				continue
			}
			v.SizeNormal = v.SizeNormal + eachFolder.Size
		}
		for _, eachFile := range v.Files {
			if regexableFile(eachFile.Filename) {
				continue
			}
			v.SizeNormal = v.SizeNormal + eachFile.Size
		}
		// percentages
		for _, eachFolder := range v.Folders {
			if regexableFolder(eachFolder.Folder) {
				continue
			}
			eachFolder.Percentage = math.Round(float64(v.SizeNormal)/float64(eachFolder.Size)*100) / 100.
		}
		for _, eachFile := range v.Files {
			if regexableFile(eachFile.Filename) {
				continue
			}
			eachFile.Percentage = math.Round(float64(v.SizeNormal)/float64(eachFile.Size)*100) / 100.
		}

	}

	////
	////

	Paranoia = true
	ExParanoia = false
	for _, v := range users.AllUsers {
		// sizes
		for _, eachFolder := range v.Folders {
			if regexableFolder(eachFolder.Folder) {
				continue
			}
			v.SizeParanoid = v.SizeParanoid + eachFolder.Size
		}
		for _, eachFile := range v.Files {
			if regexableFile(eachFile.Filename) {
				continue
			}
			v.SizeParanoid = v.SizeParanoid + eachFile.Size
		}
		// percentages
		for _, eachFolder := range v.Folders {
			if regexableFolder(eachFolder.Folder) {
				continue
			}
			eachFolder.Percentage = math.Round(float64(v.SizeParanoid)/float64(eachFolder.Size)*100) / 100.
		}
		for _, eachFile := range v.Files {
			if regexableFile(eachFile.Filename) {
				continue
			}
			eachFile.Percentage = math.Round(float64(v.SizeParanoid)/float64(eachFile.Size)*100) / 100.
		}

	}

	////
	////

	Paranoia = true
	ExParanoia = true
	for _, v := range users.AllUsers {
		// sizes
		for _, eachFolder := range v.Folders {
			if regexableFolder(eachFolder.Folder) {
				continue
			}
			v.SizeExParanoid = v.SizeExParanoid + eachFolder.Size
		}
		for _, eachFile := range v.Files {
			if regexableFile(eachFile.Filename) {
				continue
			}
			v.SizeExParanoid = v.SizeExParanoid + eachFile.Size
		}
		// percentages
		for _, eachFolder := range v.Folders {
			if regexableFolder(eachFolder.Folder) {
				continue
			}
			eachFolder.Percentage = math.Round(float64(v.SizeExParanoid)/float64(eachFolder.Size)*100) / 100.
		}
		for _, eachFile := range v.Files {
			if regexableFile(eachFile.Filename) {
				continue
			}
			eachFile.Percentage = math.Round(float64(v.SizeExParanoid)/float64(eachFile.Size)*100) / 100.
		}
	}

	Paranoia = false
	ExParanoia = false
}

func getSizeSingle(user u.UserProfile) string {
	// temp := float64(DirSize(user.AbsolutePath + "\\"))
	/*
		if temp > 1024*1024*1024 {
			return STR(temp/1024.0/1024.0/1024.0) + "_GB"
		} else if temp > 1024 {
			return STR(temp/1024.0/1024.0) + "_MB"
		} else {
			return STR(temp/1024.0) + "_KB"
		}
		return ""
	*/
	return ""
}
