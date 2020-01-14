package diskmanagement

import (
	"io/ioutil"
	"os"

	p "github.com/mjmorell/gobak/stringformatting"

	"github.com/ricochet2200/go-disk-usage/du"
)

func getDrives(drivestoinclude ...string) (r DiskCollection) {
	p.LoadingGif()
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		_, err := os.Open(string(drive) + ":\\")
		if err != nil {
			continue
		}

		r.AllDrives = append(r.AllDrives, getInformation(string(drive)+":\\"))

	}
	return
}

// DEPRECATED
// func isHDBackups(src string) bool {
// 	p.LoadingGif()
// 	entries, _ := ioutil.ReadDir(src)
// 	for _, entry := range entries {
// 		if entry.Name() == "_backupFlag" {
// 			return true
// 		}
// 	}
// 	return false
// }

func getInformation(src string) (drive d.Drive) {
	p.LoadingGif()
	drive.DiskUsed = du.NewDiskUsage(src)
	drive.Path = src
	entries, _ := ioutil.ReadDir(src)

	if isHDBackups(src) {
		drive.Path = src + "backups\\"
	} else {
		for _, entry := range entries {
			//fmt.Printf(entry.Name())
			//Wait()
			if entry.Name() == "Users" {
				drive.Path = src + "Users\\"
				break
			}
		}
	}
	drive.FreeSpace = drive.DiskUsed.Free() / 1024 / 1024
	return
}

// DEPRECATED: What's this 'Liberty'?
// func MapDrive(username, password string) ([]byte, error) {
// 	p.LoadingGif()
// 	temp := "/user:" + username
// 	exec.Command("net", "use", "M:", "/delete").CombinedOutput()
// 	/*if p.QYesNo("Is this a Liberty Machine?") {
// 		return exec.Command("net", "use", "M:", `\\fs3\hdbackups`, temp, password, "/P:NO").CombinedOutput()
// 	} else {*/
// 	return exec.Command("net", "use", "M:", `\\fs3.liberty.edu\hdbackups`, temp, password, "/P:NO").CombinedOutput()
// }

// DEPRECATED: Why did I do this here?
// func DiskSelection() (src, dst string) {
// 	p.LoadingGif()
// 	driveList := getDrives()

// 	src = p.QCustom("What is the Backup Source?", driveList.GetList(), driveList.GetEasyPrint())
// 	dst = p.QCustom("What is the Backup Destination?", driveList.GetList(), driveList.GetEasyPrint())

// 	return
// }
