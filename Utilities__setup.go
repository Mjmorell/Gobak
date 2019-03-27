package main

import (
	"net/http"
	"regexp"
	"strings"

	p "github.com/mjmorell/gobak/strutils"

	goserve "github.com/mjmorell/GoServe"
)

func connected() (ok bool) {
	_, err := http.Get("http://google.com")
	if err != nil {
		return false
	}
	return true
}

func identifier() {
	for true {
		backupInfo.ID = strings.ToUpper(p.QGeneric("What is the Ticket or Identification you would like to use?"))
		if ok, _ := regexp.MatchString("(^CS\\w+)", backupInfo.ID); ok {
			if ok, _ := regexp.MatchString("(^CS[0-9]{7}$)", backupInfo.ID); ok {
				test, len := goserveClient.PULL("sc_req_item", goserve.Filter("number")+goserve.IS(backupInfo.ID))
				if len == 1 {
					if !p.QYesNo("To Confirm, is " + p.Red(backupInfo.ID) + p.Cyan(" correct?")) {
						continue
					}
					backupInfo.IDServiceNowInformation = test[0]
					return
				}
			}
			if p.QYesNo("You typed " + p.Red(backupInfo.ID) + p.Cyan(", and this is not a valid CS#. Is this correct?")) {
				return
			}
		} else {
			if !p.QYesNo("To Confirm, is " + p.Red(backupInfo.ID) + p.Cyan(" correct?")) {
				continue
			}
			return
		}
	}
}
