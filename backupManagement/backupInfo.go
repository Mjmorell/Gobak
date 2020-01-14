package backupmanagement

import (
	"fmt"

	cm "gobak/consolemanagement"
	"gobak/diskmanagement"
	. "gobak/stringformatting"
)

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

func agreement() {
	cm.Header()
	cm.CenterString(Green("Please note, ")+Red("this program is not 100%% guaranteed."), 0, 49)
	cm.CenterString("This program is built as-is, and is not certified by Anyone.", 0, 56)
	cm.CenterString("Any questions can be directed to the github repo "+Green("GoBak"), 0, 53)

	cm.CenterString(Red("──────────────────────────────"), 0, 30)

	cm.CenterString("All sizing estimates are based off of division with "+Green("1024")+" as bytes", 0, 65)
	cm.CenterString("These sizing estimates are based off the binary system!! Such as:", 0, 65)
	cm.CenterString("A file sized at 10 Kilobytes is 10*1024 Bytes. Not 10*1000 Bytes.", 0, 65)
	fmt.Println()
	cm.CenterString("Therefore, size estimates post-backup may be different depending", 0, 65)
	cm.CenterString("on which Operating System the sizing estimate is made from.", 0, 59)
	cm.Wait(Cyan("Press") + Magenta(" [Enter] ") + Cyan("to acknowledge"))
}
