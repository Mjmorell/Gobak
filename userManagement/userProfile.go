package usermanagement

import (
	"github.com/mjmorell/gobak/filemanagement"
	"github.com/mjmorell/gobak/foldermanagement"
)

// UserProfile holds unique information for each found user
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

	//Percentage holds the percentage total of /Users/* that this user is, may be useful for progress bars. idk
	Percentage float64

	//Percentage holds the percentage total of /Users/%USERNAME%/* that this user is, may be useful for progress bars. idk
	PercentageP float64

	//Folders hold the root folders in the user directory
	Folders []foldermanagement.Folder

	//Files is files
	Files []filemanagement.File
}
