package consolemanagement

import (
	"fmt"

	c "gobak/stringformatting"
)

//Agreement is an agreement for craps sake
func Agreement() {
	Header()
	PrintlnCenter("Please note, "+c.Red("this program is NOT guaranteed."), 0, 44)
	PrintlnCenter("This program is built as-is and is not certified by anyone", 0, 58)
	PrintlnCenter("All questions can be directed to the github repo", 0, 48)

	PrintlnCenter(c.Red("──────────────────────────────"), 0, 30)

	PrintlnCenter("All sizing estimates are based off of division with 2^x as bytes", 0, 65)
	PrintlnCenter("These sizing estimates are based off the binary system!! Such as:", 0, 65)
	PrintlnCenter("A file sized at a Kilobyte is 1024 (2^10) Bytes. Not 1000 (10^3) Bytes.", 0, 71)
	fmt.Println()
	PrintlnCenter("Therefore, size estimates post-backup may be different depending", 0, 65)
	PrintlnCenter("on which Operating System the sizing estimate is made from.", 0, 59)
	Wait(c.Cyan("Press") + c.Magenta(" [Enter] ") + c.Cyan("to acknowledge"))
}
