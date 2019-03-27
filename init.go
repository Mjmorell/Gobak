package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	p "github.com/mjmorell/gobak/strutils"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Version Setup

	p.Version = Version
	p.GitCommit = GitCommit

	// OS Setup

	p.Header()

	if runtime.GOOS == "windows" {
		backupInfo.OS = "WIN"
	} else if runtime.GOOS == "darwin" {
		backupInfo.OS = "MAC"
	} else {
		p.PanicR("You are running this on a device that is not Windows or Mac. Sorry but this device is not supported.")
	}

	// Network check

	if !connected() {
		if !p.QYesNo("You are not connected to any network. Would you like to continue in Offline Mode?") {
			p.Header()
			fmt.Println("Ending Program Execution")
			fmt.Println("Check Internet Connection")
			fmt.Println()
			os.Exit(0)
		}
		p.OffMode = true
	} else {
		p.OffMode = false
	}

	// Dev / PPRD flags setup for url

	devPtr := flag.Bool("dev", false, "Devmode for program")
	pprdPtr := flag.Bool("pprd", false, "Devmode for program")

	flag.Parse()

	goserveClient.Instance = "libertydev.service-now.com/"
	p.DevMode = 0
	if *pprdPtr {
		p.DevMode = 1
		goserveClient.Instance = "libertypprd.service-now.com/"
	}
	if *devPtr {
		p.DevMode = -1
		goserveClient.Instance = "libertydev.service-now.com/"
	}
}
