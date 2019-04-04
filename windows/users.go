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
