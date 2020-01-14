package windows

import (
	"fmt"

	humanize "github.com/dustin/go-humanize"
	p "github.com/mjmorell/gobak/consolemanagement"
	u "github.com/mjmorell/gobak/userutils"
)

func ParanoiaRequestWindows(users u.UserList) {
	p.Header()
	fmt.Printf(p.Green("  USERNAME     ") + p.Cyan("MODE") + "   | %% of total source | " + p.Yellow("      Size in Bytes") + "\n")
	fmt.Println("────────────────────────────────────────────────────────────────────────────────────────────────────")
	fmt.Printf(p.Green("  OVERALL                                   ~TB,~GB,~MB,~KB,  B       | Estimated\n"))
	fmt.Printf(p.Cyan("              NORMAL ")+" |                   | %28s bytes | %s\n", p.Yellow(humanize.Comma(int64(users.TotalSizeNormal))), printSize(users.TotalSizeNormal))
	fmt.Printf(p.Cyan("             PARANOID")+" |                   | %28s bytes | %s\n", p.Yellow(humanize.Comma(int64(users.TotalSizeParanoid))), printSize(users.TotalSizeParanoid))
	fmt.Println("────────────────────────────────────────────────────────────────────────────────────────────────────")
	for _, v := range users.AllUsers {
		fmt.Printf(p.Green("  %-41s ~TB,~GB,~MB,~KB,  B       | Estimated\n"), v.Username)
		fmt.Printf(p.Cyan("              NORMAL ")+" |       %05.2f%%      | %28s bytes | %s\n", v.Percentage, p.Yellow(humanize.Comma(int64(v.SizeNormal))), printSize(v.SizeNormal))
		fmt.Printf(p.Cyan("             PARANOID")+" |       %05.2f%%      | %28s bytes | %s\n", v.PercentageP, p.Yellow(humanize.Comma(int64(v.SizeParanoid))), printSize(v.SizeParanoid))
	}

	fmt.Println("────────────────────────────────────────────────────────────────────────────────────────────────────")

	if p.QYesNoNOCLEAR("Would you like to use normal backup? Select (NO) for Paranoia Note: Not recommended in MOST cases.") {
		Paranoia = false
	} else {
		Paranoia = true
	}
}
