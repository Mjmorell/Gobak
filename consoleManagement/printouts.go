package consolemanagement

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	. "gobak/stringformatting"
)

// TODO: Keep cleaning up the starter code, start with CenterString ones

func Header(col ...func(...interface{}) string) {
	Clear()
	fmt.Println("────────────────────────────────────────────────────────────────────────────────────────────────────")
	fmt.Println("    " + HIYellow("GoBak"))
	fmt.Println("────────────────────────────────────────────────────────────────────────────────────────────────────")

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
	if len(str) > 0 {
		PrintlnCenter("──────────────────────────────", 0, 30)
		PrintlnCenter(str[0], 0, 28)
		PrintlnCenter("──────────────────────────────", 0, 30)
	} else {
		PrintlnCenter("──────────────────────────────", 0, 30)
		PrintlnCenter(Cyan("Press")+Magenta(" [Enter] ")+Cyan("to continue."), 0, 26)
		PrintlnCenter("──────────────────────────────", 0, 30)
	}
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

//PrintlnCenter Prints a string in the center
func PrintlnCenter(str string, totalWidth, strWidth int) {
	if totalWidth == 0 {
		totalWidth = 100
	}
	if strWidth == 0 {
		fmt.Printf(fmt.Sprintf("%[1]*s", -totalWidth, fmt.Sprintf("%[1]*s", (totalWidth+len(str))/2, str)))
		fmt.Println()
		return
	}
	for i := 0; i < (totalWidth-strWidth)/2; i++ {
		fmt.Printf(" ")
	}
	fmt.Println(str)

}

func SPrintlnCenter(str string, totalWidth, strWidth int) (temp string) {
	if totalWidth == 0 {
		totalWidth = 100
	}
	if strWidth == 0 {
		return fmt.Sprintf("%[1]*s", -totalWidth, fmt.Sprintf("%[1]*s", (totalWidth+len(str))/2, str))
	}
	for i := 0; i < (totalWidth-strWidth)/2; i++ {
		temp = temp + " "
	}
	temp = temp + str
	return temp
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
