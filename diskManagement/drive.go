package diskmanagement

import (
	"github.com/ricochet2200/go-disk-usage/du"
)

// Disk is for each mapped entity, such as C: or D: on Windows
type Disk struct {
	//Path holds device path, '%DISKLETTER%://' in Windows, '/Volumes/%NAME%' in MacOS
	Path string

	//FreeSpace holds free space (in MB) for each mounted Disk
	FreeSpace uint64

	DiskUsed *du.DiskUsage
}
