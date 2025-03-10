package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var vuln = goyek.Define(goyek.Task{
	Name:  "vuln",
	Usage: "govulncheck",
	Action: func(a *goyek.A) {
		cmd.Exec(a, "go tool govulncheck ./...")
	},
})
