package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var lint = goyek.Define(goyek.Task{
	Name:  "lint",
	Usage: "golangci-lint run --fix",
	Action: func(a *goyek.A) {
		if !cmd.Exec(a, "go install github.com/golangci/golangci-lint/cmd/golangci-lint", cmd.Dir(dirBuild)) {
			return
		}
		cmd.Exec(a, "golangci-lint run --fix", cmd.Dir(dirRoot))
		cmd.Exec(a, "golangci-lint run --fix", cmd.Dir(dirBuild))
	},
})
