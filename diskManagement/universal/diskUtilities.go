package diskmanagement

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func DirectorySize(src string) uint64 {
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
			size += DirectorySize(srcPath)
			//} else if regexableFile(entry.Name()) {
			//	continue
		} else {
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}
			size += uint64(entry.Size())
		}
	}

	return size
}
