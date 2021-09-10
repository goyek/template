package main

import (
	"io"
	"strings"

	"github.com/goyek/goyek"
)

func main() {
	flow().Main()
}

func flow() *goyek.Flow {
	flow := &goyek.Flow{}

	// parameters
	ci := flow.RegisterBoolParam(goyek.BoolParam{
		Name:  "ci",
		Usage: "Whether CI is calling the build script",
	})

	// tasks
	clean := flow.Register(taskClean())
	fmt := flow.Register(taskFmt())
	test := flow.Register(taskTest())
	lint := flow.Register(taskLint())
	modTidy := flow.Register(taskModTidy())
	diff := flow.Register(taskDiff(ci))

	// pipelines
	all := flow.Register(taskAll(goyek.Deps{
		clean,
		fmt,
		test,
		lint,
		modTidy,
		diff,
	}))
	flow.DefaultTask = all

	return flow
}

func taskClean() goyek.Task {
	return goyek.Task{
		Name:  "clean",
		Usage: "remove git ignored files",
		Action: func(tf *goyek.TF) {
			if err := tf.Cmd("git", "clean", "-fX").Run(); err != nil {
				tf.Fatal(err)
			}
		},
	}
}

func taskFmt() goyek.Task {
	return goyek.Task{
		Name:  "fmt",
		Usage: "go fmt",
		Action: func(tf *goyek.TF) {
			ForGoModules(tf, func(tf *goyek.TF) {
				if err := tf.Cmd("go", "fmt", "./...").Run(); err != nil {
					tf.Fatal(err)
				}
			})
		},
	}
}

func taskTest() goyek.Task {
	return goyek.Task{
		Name:  "test",
		Usage: "go test with race detector and code covarage",
		Action: func(tf *goyek.TF) {
			ForGoModules(tf, func(tf *goyek.TF) {
				if err := tf.Cmd("go", "test", "-race", "./...").Run(); err != nil {
					tf.Fatal(err)

				}
			})

		},
	}
}

func taskLint() goyek.Task {
	return goyek.Task{
		Name:  "golangci-lint",
		Usage: "golangci-lint",
		Action: func(tf *goyek.TF) {
			ForGoModules(tf, func(tf *goyek.TF) {
				if err := tf.Cmd("go", "vet", "./...").Run(); err != nil {
					tf.Fatal(err)
				}
			})
		},
	}
}

func taskModTidy() goyek.Task {
	return goyek.Task{
		Name:  "mod-tidy",
		Usage: "go mod tidy",
		Action: func(tf *goyek.TF) {
			ForGoModules(tf, func(tf *goyek.TF) {
				if err := tf.Cmd("go", "mod", "tidy").Run(); err != nil {
					tf.Error(err)
				}
			})
		},
	}
}

func taskDiff(ci goyek.RegisteredBoolParam) goyek.Task {
	return goyek.Task{
		Name:   "diff",
		Usage:  "git diff",
		Params: goyek.Params{ci},
		Action: func(tf *goyek.TF) {
			if !ci.Get(tf) {
				tf.Skip("ci param is not set, skipping")
			}

			if err := tf.Cmd("git", "diff", "--exit-code").Run(); err != nil {
				tf.Error(err)
			}

			cmd := tf.Cmd("git", "status", "--porcelain")
			sb := &strings.Builder{}
			cmd.Stdout = io.MultiWriter(tf.Output(), sb)
			if err := cmd.Run(); err != nil {
				tf.Error(err)
			}
			if sb.Len() > 0 {
				tf.Error("git status --porcelain returned output")
			}
		},
	}
}

func taskAll(deps goyek.Deps) goyek.Task {
	return goyek.Task{
		Name:  "all",
		Usage: "build pipeline",
		Deps:  deps,
	}
}
