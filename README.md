# goyek Repository Template

[![build](https://github.com/goyek/template/actions/workflows/build.yml/badge.svg)](https://github.com/goyek/template/actions/workflows/build.yml)

This is a simple Go GitHub repository template that uses [`goyek`](https://github.com/goyek/goyek) for build automation.

## Usage

1. Click the `Use this template` button (alt. clone, fork or download this repository).
1. Replace all occurences of `goyek/template` to `your_org/repo_name` in all files.
1. Update the following files:
   - [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md)
   - [LICENSE](LICENSE)
   - [README.md](README.md)

## Build

### Bash

```sh
./goyek.sh -h
./goyek.sh
```

### PowerShell

```pwsh
.\goyek.ps1 -h
.\goyek.ps1
```

### Visual Studio Code

`F1` → `Tasks: Run Build Task (Ctrl+Shift+B or ⇧⌘B)`
