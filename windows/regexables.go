package windows

import (
	"regexp"
	"strings"
)

func regexableFolder(test string) bool {
	test = strings.ToLower(test)
	if matched, _ := regexp.MatchString("appdata", test); matched {

		return true
	} else if matched, _ := regexp.MatchString("dropbox.*", test); matched {
		if Paranoia {
			return false
		}

		return true
	} else if matched, _ := regexp.MatchString("onedrive", test); matched {
		if Paranoia {
			return false
		}

		return true
	} else if matched, _ := regexp.MatchString("\\..*", test); matched {
		if Paranoia {
			return false
		}

		return true

	}
	return false
}

func regexableFile(test string) bool {
	test = strings.ToLower(test)
	if matched, _ := regexp.MatchString(`.*\.dat`, strings.ToLower(test)); matched {
		return true

	} else if matched, _ := regexp.MatchString(`.*\.lnk`, test); matched {
		return true

	} else if matched, _ := regexp.MatchString("ntuser.*", strings.ToLower(test)); matched {
		return true

	} else if matched, _ := regexp.MatchString("desktop.ini", test); matched {
		return true

	}
	return false
}
