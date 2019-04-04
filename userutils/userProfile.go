package userutils

import (
	"fmt"
	"math"
	"strconv"

	humanize "github.com/dustin/go-humanize"
	p "github.com/mjmorell/gobak/strutils"
)

type UserList struct {
	AllUsers []UserProfile

	TotalSizeNormal uint64

	TotalSizeParanoid uint64
}

type UserProfile struct {
	//Username holds the /Users/%USERNAME% user per directory
	Username string

	//RelativePath holds '/Users/%USERNAME%' or equivalent
	RelativePath string

	//AbsolutePath holds '%DEVICE%/Users/%USERNAME%' or equivalent
	AbsolutePath string

	//SizeNormal holds the total size (IN BYTES) in normal backup mode of the user folder
	SizeNormal uint64

	//SizeParanoid holds the total size (IN BYTES) in paranoid backup mode of the user folder
	SizeParanoid uint64

	//Percentage holds the percentage total of /Users/* that this user is, may be useful for progress bars. idk
	Percentage float64

	//Percentage holds the percentage total of /Users/%USERNAME%/* that this user is, may be useful for progress bars. idk
	PercentageP float64

	//Folders hold the root folders in the user directory
	Folders []UserRootFolder

	//Files is files
	Files []UserRootFile
}

//Backup backs the user folder up. This will break it down per folder internally to the []userRootFolder
func (u UserProfile) Backup(dst string) {
	for _, v := range u.Folders {
		v.Backup(dst)
	}
}

func (u UserList) GetList() (temp []string) {
	for _, v := range u.AllUsers {
		temp = append(temp, v.Username)
	}
	return
}

func (u UserList) GetEasyPrint() (temp []string) {
	for _, v := range u.AllUsers {
		if v.SizeNormal > 2048 {
			temp = append(temp, v.Username+" - "+strconv.Itoa(int(v.SizeNormal/1024))+" GB")
		} else {
			temp = append(temp, v.Username+" - "+strconv.Itoa(int(v.SizeNormal))+" MB")
		}
	}
	return
}

func (u UserList) LogPrintoutAllUsers() {
	p.Header()
	for k := range u.AllUsers {
		fmt.Println()
		fmt.Println(p.HIMagenta(u.AllUsers[k].Username) + " " + p.Yellow(humanize.Comma(int64(u.AllUsers[k].SizeNormal))) + " bytes")

		for _, v := range u.AllUsers[k].Folders {
			if v.Mode != 0 {
				fmt.Printf(p.Red("  Skipped ") + v.Folder + "\\ \n")
				continue
			}
			if v.Size == 0 {
				continue
			}
			fmt.Printf("    %05.2f%% = %-15s - %sb\n", v.Percentage, v.Folder+"\\", p.Yellow(humanize.Comma(int64(v.Size))))
		}
		if len(u.AllUsers[k].Files) > 0 {
			fmt.Println(p.HIGreen("  FILES"))
			for _, v := range u.AllUsers[k].Files {
				if v.Mode == 3 {
					continue
				}
				if v.Mode != 0 {
					fmt.Printf(p.Red("  Skipped ") + v.Filename + "\n")
					continue
				}
				if v.Size == 0 {
					continue
				}
				fmt.Printf("    %05.2f%% = %-15s - %sb\n", v.Percentage, v.Filename, p.Yellow(humanize.Comma(int64(v.Size))))
			}
		}
	}
}

func (u UserList) SetSize() UserList {
	for _, v := range u.AllUsers {
		u.TotalSizeNormal = u.TotalSizeNormal + v.SizeNormal
		u.TotalSizeParanoid = u.TotalSizeParanoid + v.SizeParanoid
	}

	for k, v := range u.AllUsers {
		u.AllUsers[k].Percentage = math.Round((float64(v.SizeNormal)/float64(u.TotalSizeNormal))*10000) / 100.
		u.AllUsers[k].PercentageP = math.Round((float64(v.SizeParanoid)/float64(u.TotalSizeParanoid))*10000) / 100.
	}

	return u
}


