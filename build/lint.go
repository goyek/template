package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var lint = goyek.Define(goyek.Task{
	Name:  "lint",
	Usage: "golangci-lint run --fix",
	Action: func(a *goyek.A) {
		cmd.Exec(a, "go tool golangci-lint run --fix")
	},
})
