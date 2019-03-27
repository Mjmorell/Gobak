package windows

import (
	"fmt"
	"regexp"
	"strings"
)

func regexableFolder(test string) bool {
	test = strings.ToLower(test)
	if matched, _ := regexp.MatchString("appdata", test); matched {

		return true
	} else if matched, _ := regexp.MatchString("dropbox.*", test); matched {
		if ExParanoia {
			return false
		}

		return true
	} else if matched, _ := regexp.MatchString("onedrive", test); matched {
		if ExParanoia {
			return false
		}

		return true
	} else if matched, _ := regexp.MatchString("\\..*", test); matched {
		if Paranoia || ExParanoia {
			return false
		}

		return true

	}
	return false
}

func regexableFile(test string, regex ...interface{}) bool {
	test = strings.ToLower(test)
	if matched, _ := regexp.MatchString(`.*\.dat`, test); matched {
		if len(regex) > 0 {
			fmt.Printf("   - Skipped %s\n", test)
		}
		return true
	} else if matched, _ := regexp.MatchString(`.*\.lnk`, test); matched {
		if Paranoia || ExParanoia {
			return false
		}
		if len(regex) > 0 {
			fmt.Printf("   - Skipped %s\n", test)
		}
		return true
	} else if matched, _ := regexp.MatchString("ntuser.*", test); matched {
		if len(regex) > 0 {
			fmt.Printf("   - Skipped %s\n", test)
		}
		return true
	} else if matched, _ := regexp.MatchString("desktop.ini", test); matched {
		if len(regex) > 0 {
			fmt.Printf("   - Skipped %s\n", test)
		}
		return true
	}
	return false
}
