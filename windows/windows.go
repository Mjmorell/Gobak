package windows

import (
	"fmt"

	humanize "github.com/dustin/go-humanize"
	p "github.com/mjmorell/gobak/strutils"
	u "github.com/mjmorell/gobak/userutils"
)

func ParanoiaRequestWindows(users u.UserList) bool {
	fmt.Printf(p.Green("  OVERALL                        ~TB,~GB,~MB,~KB,  Byte\n"))
	fmt.Printf("  Total Size "+p.Cyan(" NORMAL ")+" =        = %28s bytes\n", p.Yellow(humanize.Comma(int64(users.TotalSizeNormal))))
	fmt.Printf("  Total Size "+p.Cyan("PARANOID")+" =        = %28s bytes\n", p.Yellow(humanize.Comma(int64(users.TotalSizeParanoid))))

	for _, v := range users.AllUsers {
		fmt.Printf(p.Green("\n  %-30s ~TB,~GB,~MB,~KB,  Byte\n"), v.Username)
		fmt.Printf("  Total Size "+p.Cyan(" NORMAL ")+" = %05.2f%% = %28s bytes\n", v.Percentage, p.Yellow(humanize.Comma(int64(v.SizeNormal))))
		fmt.Printf("  Total Size "+p.Cyan("PARANOID")+" = %05.2f%% = %28s bytes\n", v.PercentageP, p.Yellow(humanize.Comma(int64(v.SizeParanoid))))
	}

	fmt.Println()

	return p.QYesNoNOCLEAR("Would you like PARANOIA MODE to be enabled? Note: Not recommended in MOST cases.")
}
