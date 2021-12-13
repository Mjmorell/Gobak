package consolemanagement

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	. "gobak/stringformatting"
)

func Header(col ...func(...interface{}) string) {
	Clear()
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("    " + HIYellow("GoBak"))
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

func PanicR(str string) {
	Clear()
	fmt.Println()
	fmt.Println(Red(str))
	fmt.Println()
	os.Exit(9)
}

//Wait is a function used to pause the program
func Wait(str ...string) {
	Flush()
	fmt.Println()
	border := ""
	printer := Cyan("Press") + Magenta(" [Enter] ") + Cyan("to continue.")
	if len(str) > 0 {
		printer = str[0]
	}
	for i := 0; i < Length(printer)+2; i++ {
		border += "━"
	}

	PrintlnCentered("┏" + border + "┓")
	PrintlnCentered("┃ " + printer + " ┃")
	PrintlnCentered("┗" + border + "┛")

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

//Clear clears the screen
func Clear() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	}
}

func Flush() {
	reader := bufio.NewReader(os.Stdin)
	var i int
	for i = 0; i < reader.Buffered(); i++ {
		reader.ReadByte()
	}
}

func PrintlnCentered(str string) {
	for i := 0; i < (100-Length(str))/2; i++ {
		fmt.Printf(" ")
	}
	fmt.Println(str)
}

var (
	loadingCycle = 0
)

func LoadingGif() {
	switch loadingCycle % 6 {
	case 0:
		fmt.Printf("\rLoading                                                                                            ")

	case 1:
		fmt.Printf("\rLoading █                                                                                          ")

	case 2:
		fmt.Printf("\rLoading █ █                                                                                        ")

	case 3:
		fmt.Printf("\rLoading █ █ █                                                                                      ")

	case 4:
		fmt.Printf("\rLoading █ █                                                                                        ")

	case 5:
		fmt.Printf("\rLoading █                                                                                          ")
	}
	loadingCycle = loadingCycle + 1
}
