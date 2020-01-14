package consolemanagement

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"gobak/backupmanagement"
	. "gobak/stringformatting"
)

func Header(col ...func(...interface{}) string) {
	Clear()
	fmt.Println("────────────────────────────────────────────────────────────────────────────────────────────────────")
	fmt.Println(CenterStringR(HIYellow("GoBak")+` ── v`+fmt.Sprintf("%4s", backupmanagement.Version)+"_"+backupmanagement.GitCommit, 0, 38))
	fmt.Println("────────────────────────────────────────────────────────────────────────────────────────────────────")

}

func PanicR(str string) {
	Clear()
	fmt.Printf("\n")
	fmt.Printf(Red(str))
	fmt.Printf("\n\n")
	os.Exit(9)
}

//Wait is a function used to pause the program
func Wait(str ...string) {
	Flush()
	fmt.Println()
	if len(str) > 0 {
		CenterString("──────────────────────────────", 0, 30)
		CenterString(str[0], 0, 28)
		CenterString("──────────────────────────────", 0, 30)
	} else {
		CenterString("──────────────────────────────", 0, 30)
		CenterString(Cyan("Press")+Magenta(" [Enter] ")+Cyan("to continue."), 0, 26)
		CenterString("──────────────────────────────", 0, 30)
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

func CenterString(s string, w, sw int) {
	if w == 0 {
		w = 100
	}
	if sw == 0 {
		fmt.Printf(fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s)))
		fmt.Printf("\n")
		return
	}
	for i := 0; i < (w-sw)/2; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf(s + "\n")

}

func CenterStringR(s string, w, sw int) (temp string) {
	if w == 0 {
		w = 100
	}
	if sw == 0 {
		return fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s))
	}
	for i := 0; i < (w-sw)/2; i++ {
		temp = temp + " "
	}
	temp = temp + s
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
