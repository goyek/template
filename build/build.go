package main

import (
	"os/exec"
	"strings"

	"github.com/goyek/goyek/v2"
)

var build = goyek.Define(goyek.Task{
	Name:  "build",
	Usage: "go build",
	Action: func(tf *goyek.TF) {
		extBytes, err := exec.CommandContext(tf.Context(), "go", "env", "GOEXE").CombinedOutput()
		if err != nil {
			tf.Fatal(err)
		}
		ext := strings.TrimSpace(string(extBytes))
		Exec(tf, dirRoot, `go build -ldflags="-s -w" -o dist/app`+ext)
	},
})
