package main

import (
	"os/exec"
	"strings"

	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var build = goyek.Define(goyek.Task{
	Name:  "build",
	Usage: "go build",
	Action: func(a *goyek.A) {
		extBytes, err := exec.CommandContext(a.Context(), "go", "env", "GOEXE").CombinedOutput()
		if err != nil {
			a.Fatal(err)
		}
		ext := strings.TrimSpace(string(extBytes))
		cmd.Exec(a, `go build -ldflags="-s -w" -o dist/app`+ext)
	},
})
