package strutils

/************************************************
	MIT License
	Details viewable in the Github Directory
	Copyright (c) 2018 Michael Morell
*************************************************
	Created by Michael Morell
	Released 04/04/2018
	Github: https://github.com/Mjmorell/WinbakGo
************************************************/

import (
	"fmt"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func QGeneric(question ...string) string {
	Header()
	var answer string
	for _, each := range question {
		fmt.Printf("    %s\n", Cyan(each))
	}
	fmt.Printf("  > ")
	fmt.Scan(&answer)
	Header()
	return answer
}

func QPassword(question string) string {
	var answer string
	for true {
		Header()
		fmt.Printf("    %s\n", Cyan(question))
		fmt.Printf("  > ")
		answer = password()
		if answer != "" {
			break
		}
	}
	Header()
	return answer
}

func password() string {
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	return string(bytePassword)
}

func QChoiceString(question string, choices ...string) string {
	var answer int
	for true {
		Header()
		fmt.Printf("    %s\n", Cyan(question))
		fmt.Println("    ─────")
		for i, value := range choices {
			fmt.Printf("(%v) %s\n", i+1, value)
		}
		fmt.Println("    ─────")
		fmt.Printf("  > ")
		fmt.Scan(&answer)
		if answer > 0 && answer < len(choices)+1 {
			Header()
			return choices[answer-1]
		}
	}
	panic("HOW DID THIS HAPPEN?")
}

func QYesNo(question string) bool {
	var answer string
	for true {
		Header()
		fmt.Printf("    %s\n", Cyan(question))
		fmt.Println("    ─────")
		fmt.Println("(1) Yes")
		fmt.Println("(2) No")
		fmt.Println("    ─────")
		fmt.Printf("  > ")

		fmt.Scan(&answer)
		answer = strings.ToLower(answer)
		if answer == "y" || answer == "yes" || answer == "1" {
			Header()
			return true
		} else if answer == "n" || answer == "no" || answer == "2" {
			Header()
			return false
		}
	}
	panic("HOW DID THIS HAPPEN?")
}

func QYesNoNOCLEAR(question string) bool {
	var answer string
	for true {
		fmt.Printf("    %s\n", Cyan(question))
		fmt.Println("    ─────")
		fmt.Println("(1) Yes")
		fmt.Println("(2) No")
		fmt.Println("    ─────")
		fmt.Printf("  > ")

		fmt.Scan(&answer)
		answer = strings.ToLower(answer)
		if answer == "y" || answer == "yes" || answer == "1" {
			Header()
			return true
		} else if answer == "n" || answer == "no" || answer == "2" {
			Header()
			return false
		}
	}
	panic("HOW DID THIS HAPPEN?")
}

func QArray(question string, choices []string) (answer int) {
	for true {
		Header()
		fmt.Printf("    %s\n", Cyan(question))
		fmt.Println("    ─────")
		for i, value := range choices {
			fmt.Printf("(%v) %s\n", i+1, value)
		}
		fmt.Println("    ─────")

		fmt.Printf("  > ")
		fmt.Scan(&answer)
		if answer > 0 && answer < len(choices)+1 {
			Header()
			return answer - 1
		}
	}
	panic("HOW DID THIS HAPPEN?")
}

func QCustom(question string, choicesAnswer, choices []string) (answer string) {
	var temp int
	Header()
	fmt.Printf("    %s\n", Cyan(question))
	fmt.Println("    ─────")
	for i, value := range choices {
		fmt.Printf("(%v) %s\n", i+1, value)
	}
	fmt.Println("    ─────")
	fmt.Println("(0) Custom Answer")
	fmt.Println("    ─────")

	fmt.Printf("  > ")
	fmt.Scan(&temp)
	if temp > 0 && temp < len(choices)+1 {
		Header()
		return choicesAnswer[temp-1]
	} else if temp == 0 {
		return QGeneric(question)
	}
	return
}

func QChoiceMultiple(question string, choicesPre []string, information ...[]string) (answers []string) {
	var answer int
	choices := make([]string, len(choicesPre))
	copy(choices, choicesPre)
	for true {
		Header()
		if len(choices) == 0 {
			return answers
		}
		for i, value := range choices {
			fmt.Printf("%4s %s", fmt.Sprintf("(%v)", i+1), value)
			if len(information) > 0 {
				fmt.Print(" - " + information[0][i])
			}
			fmt.Printf("\n")
		}
		fmt.Println("    ─────")
		fmt.Println(" (0) Submit")
		fmt.Println("    ─────")

		fmt.Printf("    %s\n", Cyan(question))
		fmt.Printf("  > ")
		fmt.Scan(&answer)
		if answer == 0 {
			if len(answers) > 0 {
				return answers
			} else {
				continue
			}
		}
		if answer > 0 && answer < len(choices)+1 {
			answers = append(answers, choices[answer-1])
			choices = append(choices[:answer-1], choices[answer:]...)
			if len(information) > 0 {
				information[0] = append(information[0][:answer-1], information[0][answer:]...)
			}
		}
	}
	panic("HOW DID THIS HAPPEN?")
	//return -1, ""
}
