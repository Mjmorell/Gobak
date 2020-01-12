package foldermanagement

// Folder holds unique information on each folder
type Folder struct {
	//Size holds size in B
	Size uint

	Mode int

	//AbsolutePath is the absolute path from disk to the root folder
	AbsolutePath string

	//Folder is just the folder name
	Folder string
}
