package windows

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strings"

	d "github.com/mjmorell/gobak/driveutils"
	p "github.com/mjmorell/gobak/consolemanagement"
	u "github.com/mjmorell/gobak/userutils"
)

func CopyFile(src, dst string) (err error) {
	//copyLog.Println(src)
	if regexableFile(src) {
		//copyLog.Println("---------- SKIPPED " + src + "\n")
		return
	}
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()
	_, err = io.Copy(out, in)
	if err != nil {
		return
	}
	err = out.Sync()
	if err != nil {
		return
	}
	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}
	return
}

func CopyUser(backupInfo d.BackupInformation, user u.UserProfile) (err error) {
	for _, parentFolder := range user.Folders {
		if regexableFolder(parentFolder.Folder) {
			continue
		}
		if p.DevMode != 0 {
			if parentFolder.Folder == "go" {
				continue
			}
		}
		if parentFolder.Size == 0 {
			fmt.Printf(" %40s ", parentFolder.Folder)
			fmt.Printf(p.Yellow("EMPTY FOLDER\n"))
			continue
		}

		backedUp := 0.0
		fmt.Printf(" %40s [                         ] 0%%    ", parentFolder.Folder)
		src := filepath.Clean(parentFolder.AbsolutePath)
		dst := filepath.Clean(backupInfo.Destination.Path + "\\" + user.Username + "\\" + parentFolder.Folder)
		si, err := os.Stat(src)
		if err != nil {
			//fmt.Printf("  ERROR! " + err.Error() + "\n - Has occured. Please verify this error!")
			continue
		}
		if !si.IsDir() {
			continue
		}
		_, err = os.Stat(dst)
		if err != nil && !os.IsNotExist(err) {
			continue
		}
		if err == nil {
			continue
			//return fmt.Errorf("destination already exists")
		}
		err = os.MkdirAll(dst, si.Mode())
		if err != nil {
			continue
		}
		entries, err := ioutil.ReadDir(src)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			srcPath := filepath.Join(src, entry.Name())
			dstPath := filepath.Join(dst, entry.Name())
			if entry.IsDir() {
				backedUp += float64(DirSize(srcPath))
				err = CopyDirClean(srcPath, dstPath)
				if err != nil {
					continue
				}
				percentage := math.Round((backedUp/float64(parentFolder.Size))*1000.0) / 10.0
				bar := int(math.Round(percentage / 4))
				fmt.Printf("\r %40s [", parentFolder.Folder)
				for i := 0; i < 25; i++ {
					if bar > i {
						fmt.Printf("=")
					} else {
						fmt.Printf(" ")
					}
				}
				fmt.Printf("] %3.1f%%    ", percentage)
				ext := strings.Split(entry.Name(), ".")
				if len(ext) == 1 {
					if len(entry.Name()) > 20 {
						fmt.Printf("- %-20s", entry.Name()[:20])
					} else {
						fmt.Printf("- %-20s", entry.Name())
					}
				} else {
					if len(ext[len(ext)-1]) > 20 {
						fmt.Printf("- %-20s", ext[len(ext)-1][:20])
					} else {
						fmt.Printf("- %-20s", ext[len(ext)-1])
					}
				}

			} else {
				if entry.Mode()&os.ModeSymlink != 0 {
					continue
				}
				err = CopyFile(srcPath, dstPath)
				if err != nil {
					continue
				}
				backedUp += float64(entry.Size())
				percentage := math.Round((backedUp/float64(parentFolder.Size))*1000.0) / 10.0
				bar := int(math.Round(percentage / 4))
				fmt.Printf("\r %40s [", parentFolder.Folder)
				for i := 0; i < 25; i++ {
					if bar > i {
						fmt.Printf("=")
					} else {
						fmt.Printf(" ")
					}
				}
				fmt.Printf("] %3.1f%%    ", percentage)
				ext := strings.Split(entry.Name(), ".")
				if len(ext) == 1 {
					if len(entry.Name()) > 20 {
						fmt.Printf("- %-20s", entry.Name()[:20])
					} else {
						fmt.Printf("- %-20s", entry.Name())
					}
				} else {
					if len(ext[len(ext)-1]) > 20 {
						fmt.Printf("- %-20s", ext[len(ext)-1][:20])
					} else {
						fmt.Printf("- %-20s", ext[len(ext)-1])
					}
				}
			}

		}
		percentage := math.Round((backedUp/float64(parentFolder.Size))*1000.0) / 10.0
		bar := int(math.Round(percentage / 4))
		fmt.Printf("\r %40s [", parentFolder.Folder)
		for i := 0; i < 25; i++ {
			if bar > i {
				fmt.Printf("=")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("] ")

		if percentage < 90. {
			fmt.Printf(p.Red("%-8s", fmt.Sprintf("%3.1f%%    ", percentage)))
		} else if percentage < 100. {
			fmt.Printf(p.Yellow("%-8s", fmt.Sprintf("%3.1f%%    ", percentage)))
		} else {
			fmt.Printf(p.Green("100%%    "))
		}

		fmt.Printf("                        ")
		fmt.Println()
	}
	return nil
}

func CopyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)
	si, err := os.Stat(src)
	if err != nil {
		fmt.Printf("  ERROR! " + err.Error() + "\n - Has occured. Please verify this error!")
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}
	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if err == nil {
		return fmt.Errorf("destination already exists")
	}
	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			fmt.Printf("   Copying ~\\" + entry.Name() + "\n")
			err = CopyDirClean(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}
			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}
	return
}

func CopyDirClean(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)
	si, err := os.Stat(src)
	if err != nil {
		fmt.Printf("  ERROR! " + err.Error() + "\n - Has occured. Please verify this error!")
		return err
	}
	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			err = CopyDirClean(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}
			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}
	return
}
