package diskmanagement

import (
	"strconv"
)

// DiskCollection used for maintaining simple functions accross all disks
type DiskCollection struct {
	DiskList []Disk
}

// GetList returns an array of each disk
func (d DiskCollection) GetList() (temp []string) {
	for _, v := range d.DiskList {
		temp = append(temp, v.Path)
	}
	return
}

// GetEasyPrint returns a pre-formatted array for easy listing
func (d DiskCollection) GetEasyPrint() (temp []string) {
	for _, v := range d.DiskList {
		if v.FreeSpace > 1024*1024 {
			temp = append(temp, v.Path+" - "+strconv.Itoa(int(v.FreeSpace/1024))+" TB Free")
		} else if v.FreeSpace > 1024 {
			temp = append(temp, v.Path+" - "+strconv.Itoa(int(v.FreeSpace))+" GB Free")
		} else {
			temp = append(temp, v.Path+" - "+strconv.Itoa(int(v.FreeSpace))+" MB Free")
		}
	}
	return
}
