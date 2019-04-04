package strutils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

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

func Header(col ...func(...interface{}) string) {
	Clear()
	fmt.Println("────────────────────────────────────────────────────────────────────────────────────────────────────")
	fmt.Println(CenterStringR(HIYellow("ShonStop")+` ── KB0016604 ── v`+Version+"_"+GitCommit, 0, 39))
	if OffMode {
		fmt.Println(CenterStringR(HICyan("-- OFFLINE MODE --"), 0, 18))
	}
	if DevMode == 1 {
		fmt.Println(CenterStringR(HIRed("-- PPRD MODE --"), 0, 15))
	}
	if DevMode == -1 {
		fmt.Println(CenterStringR(HIRed("-- DEVELOPER MODE --"), 0, 20))
	}
	fmt.Println("────────────────────────────────────────────────────────────────────────────────────────────────────")

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
