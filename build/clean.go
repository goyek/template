package main

import (
	"os"

	"github.com/goyek/goyek/v2"
)

var _ = goyek.Define(goyek.Task{
	Name:  "clean",
	Usage: "remove remove files created during build pipeline",
	Action: func(tf *goyek.TF) {
		remove(tf, "dist")
		remove(tf, "coverage.out")
		remove(tf, "coverage.html")
	},
})

func remove(tf *goyek.TF, path string) {
	if _, err := os.Stat(path); err != nil {
		return
	}
	tf.Log("Remove: " + path)
	if err := os.RemoveAll(path); err != nil {
		tf.Error(err)
	}
}