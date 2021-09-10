package main

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/goyek/goyek"
)

// ForGoModules is a helper that executes given function
// in each directory containing go.mod file.
func ForGoModules(tf *goyek.TF, fn func(tf *goyek.TF)) {
	curDir := WorkDir(tf)
	_ = filepath.WalkDir(curDir, func(path string, dir fs.DirEntry, err error) error {
		if dir.Name() != "go.mod" {
			return nil
		}

		goModDir := filepath.Dir(path)
		tf.Log("Go Module:", goModDir)
		if err := os.Chdir(goModDir); err != nil {
			tf.Fatal(err)
		}

		fn(tf) // execute function in file containing go.mod

		return nil
	})

	defer ChDir(tf, curDir)
}

// WorkDir returns current working directory.
func WorkDir(tf *goyek.TF) string {
	curDir, err := os.Getwd()
	if err != nil {
		tf.Fatal(err)
	}
	return curDir
}

// ChDir changes the working directory.
func ChDir(tf *goyek.TF, path string) {
	if err := os.Chdir(path); err != nil {
		tf.Fatal(err)
	}
}
