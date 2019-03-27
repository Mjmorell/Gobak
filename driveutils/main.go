package driveutils

import (
	"strconv"

	"github.com/ricochet2200/go-disk-usage/du"
)

type DriveList struct {
	AllDrives []Drive
}

type Drive struct {
	//Path holds device path, '%DRIVELETTER%://' in Windows, '/Volumes/%NAME%' in MacOS
	Path string

	//FreeSpace holds free space (in MB) for each mounted drive
	FreeSpace uint64

	DiskUsed *du.DiskUsage
}

type BackupInformation struct {
	//ID holds identifier used to name the folder when data is backed up. This will generally be a CS#
	ID string

	//IDServiceNowInformation will hold the ticket csv if it is a true CS#
	IDServiceNowInformation map[string]string

	//Technician holds the tech username who did the backup. may be useful
	Technician string

	//OS is either 'MAC' or 'WIN'
	OS string

	//OSVersion holds versioning information for the OS, may be useful in debugging
	OSVersion string

	//Destination holds the 'drive' for the destination for backup
	Destination Drive

	//Source holds the 'drive' for the source of backup
	Source Drive
}

func (d DriveList) GetList() (temp []string) {
	for _, v := range d.AllDrives {
		temp = append(temp, v.Path)
	}
	return
}

func (d DriveList) GetEasyPrint() (temp []string) {
	for _, v := range d.AllDrives {
		if v.FreeSpace > 2048 {
			temp = append(temp, v.Path+" - "+strconv.Itoa(int(v.FreeSpace/1024))+" GB Free")
		} else {
			temp = append(temp, v.Path+" - "+strconv.Itoa(int(v.FreeSpace))+" MB Free")
		}
	}
	return
}
