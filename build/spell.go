package main

import (
	"strings"

	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var spell = goyek.Define(goyek.Task{
	Name:  "spell",
	Usage: "misspell",
	Action: func(a *goyek.A) {
		mdFiles := find(a, ".md")
		if len(mdFiles) == 0 {
			a.Skip("no .md files")
		}
		cmd.Exec(a, "go tool misspell -error -locale=US -w "+strings.Join(mdFiles, " "))
	},
})
