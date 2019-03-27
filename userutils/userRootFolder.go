package userutils

type UserRootFolder struct {
	//Size holds size in B
	Size uint64

	//Percentage holds the percentage total of /Users/%USERNAME%/* that this user is, may be useful for progress bars. idk
	Percentage float64

	//AbsolutePath is the absolute path from disk to the root folder
	AbsolutePath string

	//Folder is just the folder name
	Folder string
}

func (uRF UserRootFolder) Backup(dst string) {

}

type UserRootFile struct {
	//Size holds size in B
	Size uint64

	//Percentage holds the percentage total of /Users/%USERNAME%/* that this user is, may be useful for progress bars. idk
	Percentage float64

	//AbsolutePath is the absolute path from disk to the root folder
	AbsolutePath string

	//Filename is just the folder name
	Filename string
}