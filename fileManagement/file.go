package filemanagement

// File holds information on _each_ file. Might probably not be used, sorry memory
type File struct {
	//Size holds size in B
	Size uint

	Mode int

	//AbsolutePath is the absolute path from disk to the root folder
	AbsolutePath string

	//Filename is just the folder name
	Filename string
}
