package consolemanagement

import (
	"fmt"
	c "gobak/stringformatting"
)

//Agreement is an agreement for craps sake
func Agreement() {
	Header()
	fmt.Println()

	PrintlnCentered("Please note, " + c.Red("this program is NOT guaranteed."))
	PrintlnCentered("This program is built as-is and is not certified by anyone")
	PrintlnCentered("All questions can be directed to the github repo")
	fmt.Println()
	PrintlnCentered("──────────────────────────────")

	fmt.Println()
	PrintlnCentered(c.HIRed("SIZING NOTICE:"))
	fmt.Println()
	PrintlnCentered("All sizing estimates are based off of division with 2^x as bytes")
	PrintlnCentered("These sizing estimates are based off the binary system!! Such as:")
	PrintlnCentered("A file sized at a Kilobyte is 1024 (2^10) Bytes. Not 1000 (10^3) Bytes.")
	PrintlnCentered("Therefore, size estimates post-backup may be different depending")
	PrintlnCentered("on which Operating System the sizing estimate is made from.")
	Wait(c.Cyan("Press") + c.Magenta(" [Enter] ") + c.Cyan("to acknowledge"))
}
