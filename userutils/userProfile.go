package userutils

import "strconv"

type UserList struct {
	AllUsers []UserProfile
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

	//SizeExParanoid holds the total size (IN BYTES) in paranoid backup mode of the user folder
	SizeExParanoid uint64

	//Percentage holds the percentage total of /Users/* that this user is, may be useful for progress bars. idk
	Percentage float64

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
