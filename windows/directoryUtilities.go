package windows

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	d "github.com/mjmorell/gobak/driveutils"
	p "github.com/mjmorell/gobak/strutils"

	"github.com/ricochet2200/go-disk-usage/du"
)

func getDrives(drivestoinclude ...string) (r d.DriveList) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		_, err := os.Open(string(drive) + ":\\")
		if err != nil {
			continue
		}

		r.AllDrives = append(r.AllDrives, getInformation(string(drive)+":\\"))

	}
	return
}

func isHDBackups(src string) bool {
	entries, _ := ioutil.ReadDir(src)
	for _, entry := range entries {
		if entry.Name() == "_backupFlag" {
			return true
		}
	}
	return false
}

func getInformation(src string) (drive d.Drive) {

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

func MapDrive(username, password string) ([]byte, error) {
	temp := "/user:" + username
	p.Header()
	exec.Command("net", "use", "M:", "/delete").CombinedOutput()
	if p.QYesNo("Is this a Liberty Machine?") {
		return exec.Command("net", "use", "M:", `\\fs3\hdbackups`, temp, password, "/P:NO").CombinedOutput()
	} else {
		return exec.Command("net", "use", "M:", `\\fs3.liberty.edu\hdbackups`, temp, password, "/P:NO").CombinedOutput()
	}
}

func DiskSelection() (src, dst string) {
	driveList := getDrives()

	src = p.QCustom("What is the Backup Source?", driveList.GetList(), driveList.GetEasyPrint())
	dst = p.QCustom("What is the Backup Destination?", driveList.GetList(), driveList.GetEasyPrint())

	return
}

func DirSize(src string) uint64 {
	var size uint64
	si, err := os.Stat(src)
	if err != nil {
		return 0
	}
	if !si.IsDir() {
		return 0
	}
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return 0
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		if entry.IsDir() {
			size += DirSize(srcPath)
		} else if regexableFile(entry.Name()) {
			continue
		} else {
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}
			size += uint64(entry.Size())
		}
	}

	return size
}