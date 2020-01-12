package backupmanagement

import "github.com/mjmorell/GoBak/diskmanagement"

var (
	//Version is version...
	Version string
	//GitCommit holds commit
	GitCommit string
)

// BackupInfo holds
type BackupInfo struct {

	//DEPRECATED: 'ID' will simply be changed to a date/time string
	//ID holds identifier used to name the folder when data is backed up. This will generally be a CS#
	ID string

	//DEPRECATED: wtf is 'servicenow'?
	//IDServiceNowInformation will hold the ticket csv if it is a true CS#
	//IDServiceNowInformation map[string]string

	//DEPRECATED: No need for this crap
	//Technician holds the tech username who did the backup. may be useful
	//Technician string

	//OS is either ( -1 : Mac, 0 : Linux, 1 : Windows )
	OS int

	//OSVersion holds versioning information for the OS, may be useful in debugging
	OSVersion string

	//Destination holds the 'drive' for the destination for backup
	Destination diskmanagement.Disk

	//Source holds the 'drive' for the source of backup
	Source diskmanagement.Disk
}
