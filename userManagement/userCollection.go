package usermanagement

import (
	"fmt"
	"math"
	"strconv"

	humanize "github.com/dustin/go-humanize"
	cm "github.com/mjmorell/gobak/consolemanagement"
	sm "github.com/mjmorell/gobak/stringmodifiers"
)

// UserCollection is for useful functions over the entire list of users found
type UserCollection struct {
	AllUsers []UserProfile

	TotalSizeNormal uint64

	TotalSizeParanoid uint64
}

//Backup backs the user folder usm. This will break it down per folder internally to the []userRootFolder

func (u UserCollection) GetList() (temp []string) {
	for _, v := range u.AllUsers {
		temp = append(temp, v.Username)
	}
	return
}

func (u UserCollection) GetEasyPrint() (temp []string) {
	for _, v := range u.AllUsers {
		if v.SizeNormal > 2048 {
			temp = append(temp, v.Username+" - "+strconv.Itoa(int(v.SizeNormal/1024))+" GB")
		} else {
			temp = append(temp, v.Username+" - "+strconv.Itoa(int(v.SizeNormal))+" MB")
		}
	}
	return
}

func (u UserCollection) LogPrintoutAllUsers() {
	cm.Header()
	for k := range u.AllUsers {
		fmt.Println()
		fmt.Println(sm.HIMagenta(u.AllUsers[k].Username) + " " + sm.Yellow(humanize.Comma(int64(u.AllUsers[k].SizeNormal))) + " bytes")

		for _, v := range u.AllUsers[k].Folders {
			if v.Mode != 0 {
				fmt.Printf(sm.Red("  Skipped ") + v.Folder + "\\ \n")
				continue
			}
			if v.Size == 0 {
				continue
			}
			fmt.Printf("     %-15s - %sb\n", v.Folder+"\\", sm.Yellow(humanize.Comma(int64(v.Size))))
		}
		if len(u.AllUsers[k].Files) > 0 {
			fmt.Println(sm.HIGreen("  FILES"))
			for _, v := range u.AllUsers[k].Files {
				if v.Mode == 3 {
					continue
				}
				if v.Mode != 0 {
					fmt.Printf(sm.Red("  Skipped ") + v.Filename + "\n")
					continue
				}
				if v.Size == 0 {
					continue
				}
				fmt.Printf("     %-15s - %sb\n", v.Filename, sm.Yellow(humanize.Comma(int64(v.Size))))
			}
		}
	}
}

func (u UserCollection) SetSize() UserCollection {
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
