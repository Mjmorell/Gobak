package main

import (
	"math/rand"
	"time"

	"gobak/backupmanagement"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	backupmanagement.Version = Version
	backupmanagement.GitCommit = GitCommit
}
