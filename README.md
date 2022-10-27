# goyek repository template

[![GitHub Release](https://img.shields.io/github/v/release/goyek/template)](https://github.com/goyek/template/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/goyek/template.svg)](https://pkg.go.dev/github.com/goyek/template)
[![go.mod](https://img.shields.io/github/go-mod/go-version/goyek/template)](go.mod)
[![LICENSE](https://img.shields.io/github/license/goyek/template)](LICENSE)
[![Build Status](https://img.shields.io/github/workflow/status/goyek/template/build)](https://github.com/goyek/template/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/goyek/template)](https://goreportcard.com/report/github.com/goyek/template)
[![Codecov](https://codecov.io/gh/goyek/template/branch/main/graph/badge.svg)](https://codecov.io/gh/goyek/template)

⭐ `Star` this repository if you find it valuable and worth maintaining.

## Description

This is a GitHub repository template for a Go application
that uses [`goyek`](https://github.com/goyek/goyek) for build automation.

It also includes:

- continuous integration via [GitHub Actions](https://github.com/features/actions),
- dependency management using [Go Modules](https://github.com/golang/go/wiki/Modules),
- code formatting using [gofumpt](https://github.com/mvdan/gofumpt),
- linting with [golangci-lint](https://github.com/golangci/golangci-lint),
- spell checking using [misspell](https://github.com/client9/misspell),
- unit testing with
  [race detector](https://blog.golang.org/race-detector),
  code covarage [HTML report](https://blog.golang.org/cover)
  and [Codecov report](https://codecov.io/),
- dependencies scanning and updating thanks to [Dependabot](https://dependabot.com),
- security code analysis using [CodeQL Action](https://docs.github.com/en/github/finding-security-vulnerabilities-and-errors-in-your-code/about-code-scanning),
- [Visual Studio Code](https://code.visualstudio.com) configuration with
  [Go](https://code.visualstudio.com/docs/languages/go) support.

## Usage

1. Sign up on [Codecov](https://codecov.io/) and configure
   [Codecov GitHub Application](https://github.com/apps/codecov) for all repositories.
1. Click the `Use this template` button (alt. clone, fork or download this repository).
1. Replace all occurrences of `goyek/template` to `your_org/repo_name` in all files.
1. Update the following files:
   - [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md)
   - [LICENSE](LICENSE)
   - [README.md](README.md)

## Setup

Below you can find sample instructions on how to set up the development environment.
Of course, you can use other tools like [GoLand](https://www.jetbrains.com/go/),
[Vim](https://github.com/fatih/vim-go), [Emacs](https://github.com/dominikh/go-mode.el).
However, take notice that the Visual Studio Go extension is
[officially supported](https://blog.golang.org/vscode-go) by the Go team.

1. Install [Go](https://golang.org/doc/install).
1. Install [Visual Studio Code](https://code.visualstudio.com/).
1. Install [Go extension](https://code.visualstudio.com/docs/languages/go).
1. Clone and open this repository.
1. `F1` -> `Go: Install/Update Tools` -> (select all) -> OK.

## Build

```sh
cd build
go run .
```

Using convenient Bash script:

```sh
./goyek.sh
```

Using convenient PowerShell script:

```pwsh
.\goyek.ps1
```

Using Visual Studio Code:

`F1` → `Tasks: Run Build Task (Ctrl+Shift+B or ⇧⌘B)`

## Maintenance

Notable files:

- [.github/workflows](.github/workflows) - GitHub Actions workflows,
- [.github/dependabot.yml](.github/dependabot.yml) - Dependabot configuration,
- [.vscode](.vscode) - Visual Studio Code configuration files,
- [build](build) - build pipeline used for local development, [CI build](.github/workflows),
  and [.vscode/tasks.json](.vscode/tasks.json),
- [build/tools.go](build/tools.go) - [build tools](https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module),
- [.golangci.yml](.golangci.yml) - golangci-lint configuration.
