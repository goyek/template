// Build is the build pipeline for this repository.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/goyek/goyek/v2"
	"github.com/goyek/goyek/v2/middleware"
)

// Directories used in repository.
const (
	dirRoot  = "."
	dirBuild = "build"
)

var (
	dryRun  = flag.Bool("dry-run", false, "print all tasks that would be run without running them")
	longRun = flag.Duration("long-run", 0, "print when a task takes longer")
	noDeps  = flag.Bool("no-deps", false, "do not process dependencies")
	skip    = flag.String("skip", "", "skip processing the `comma-separated tasks`")
)

func main() {
	goyek.SetDefault(all)

	flag.CommandLine.SetOutput(goyek.Output())
	flag.Usage = usage
	flag.Parse()

	if *dryRun {
		goyek.Use(middleware.DryRun)
	}
	goyek.Use(middleware.ReportStatus)
	if *longRun > 0 {
		goyek.Use(middleware.ReportLongRun(*longRun))
	}

	var opts []goyek.Option
	if *noDeps {
		opts = append(opts, goyek.NoDeps())
	}
	if *skip != "" {
		skippedTasks := strings.Split(*skip, ",")
		opts = append(opts, goyek.Skip(skippedTasks...))
	}

	goyek.SetUsage(usage)
	goyek.Main(os.Args[1:])
}

func usage() {
	fmt.Println("Usage of build: [flags] [--] [tasks]")
	goyek.Print()
	fmt.Println("Flags:")
	flag.PrintDefaults()
}
