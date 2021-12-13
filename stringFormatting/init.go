package stringformatting

import (
	"github.com/fatih/color"
)

var (
	//Red and all others are colors
	Red func(...interface{}) string // 27 91 51 49 109
	//Green and all others are colors
	Green func(...interface{}) string // 27 91 51 50 109
	//Yellow and all others are colors
	Yellow func(...interface{}) string // 27 91 51 51 109
	//Blue and all others are colors
	Blue func(...interface{}) string // 27 91 51 52 109
	//Magenta and all others are colors
	Magenta func(...interface{}) string // 27 91 51 53 109
	//Cyan and all others are colors
	Cyan func(...interface{}) string // 27 91 51 54 109
	//HIRed  and all others are colors
	HIRed func(...interface{}) string // 27 91 57 49 109
	//HIGreen and all others are colors
	HIGreen func(...interface{}) string // 27 91 57 50 109
	//HIYellow and all others are colors
	HIYellow func(...interface{}) string // 27 91 57 51 109
	//HIBlue and all others are colors
	HIBlue func(...interface{}) string // 27 91 57 52 109
	//HIMagenta and all others are colors
	HIMagenta func(...interface{}) string // 27 91 57 53 109
	//HICyan and all others are colors
	HICyan func(...interface{}) string // 27 91 57 54 109
	//Bold and all others are colors
	Bold func(...interface{}) string // 27 91 49 109
	//Italic and all others are colors
	Italic func(...interface{}) string // 27 91 51 109
	//Underline and all others are colors
	Underline func(...interface{}) string // 27 91 52 109
)

func init() {
	Red = color.New(color.FgRed).SprintFunc()
	Green = color.New(color.FgGreen).SprintFunc()
	Yellow = color.New(color.FgYellow).SprintFunc()
	Blue = color.New(color.FgBlue).SprintFunc()
	Magenta = color.New(color.FgMagenta).SprintFunc()
	Cyan = color.New(color.FgCyan).SprintFunc()

	HIRed = color.New(color.FgHiRed).SprintFunc()
	HIGreen = color.New(color.FgHiGreen).SprintFunc()
	HIYellow = color.New(color.FgHiYellow).SprintFunc()
	HIBlue = color.New(color.FgHiBlue).SprintFunc()
	HIMagenta = color.New(color.FgHiMagenta).SprintFunc()
	HICyan = color.New(color.FgHiCyan).SprintFunc()

	Bold = color.New(color.Bold).SprintFunc()
	Italic = color.New(color.Italic).SprintFunc()
	Underline = color.New(color.Underline).SprintFunc()
}

//Length provides the 'visible' length of a string, ignoring special characters and double-length characters
func Length(str string) int {
	ignoreNum := 0
	newStr := []rune(str)
	// fmt.Println(len(newStr))
	for x, y := range newStr {
		// fmt.Println(x, " - ", y)
		if y == 27 {
			for i := 1; i < 5; i++ {
				// fmt.Println("    ", i, ": ", str[x+i])
				if newStr[x+i] == 109 {
					ignoreNum += i + 1
					break
				}
			}
		}
	}
	return len(newStr) - ignoreNum
}
