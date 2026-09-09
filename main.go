/*
go build -ldflags="-s -w" -o goinit.exe .
go install .
*/
package main

import (
	"fmt"
	"os"

	"os/exec"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

/***************************************************************************************************
 * [ Globals / Consts ]
 **************************************************************************************************/
const help = `
╔═════════════════════════════════════════════════════════════╗
║                           GOINIT                            ║
╠═════════════════════════════════════════════════════════════╣
║                                                             ║
║ Usage                                                       ║
║   goinit -h           Print help                            ║
║   goinit [project]    Build scaffold in current directory   ║
║                                                             ║
╚═════════════════════════════════════════════════════════════╝
`

const gitignore = `# Binaries for the current OS and architecture
*.exe
*.dll
*.so
*.dylib

# Binaries for other OS/architectures
bin/
out/
build/

# Go module cache
.go/
pkg/mod/

# Test cache
.test/

# Vendor directory (if not using Go Modules or explicitly excluded)
vendor/

# Environment files
.env
.env.*
.secrets
.secrets.*

# Editor-specific files
.idea/
.vscode/
*.swp
*~
`

const readmeBody = `

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Overview](#overview)

## Overview

> This is the start of something amazing i just know it!
`

const mainFunc = `package main

func main() {
	println("Hello, world!")
}

`

const ptr = "\x1b[38;5;203m\x1b[0m"


/***************************************************************************************************
 * [ Main ]
 **************************************************************************************************/

func main() {
    args := os.Args[1:]
    if (args[0] == "-h" || args[0] == "--help") {
        fmt.Print(help)
        os.Exit(0)
    }

    if len(args) > 1 {
        fmt.Fprintf(os.Stderr, "%s goinit expects 1 arg, project name or help\n", ptr)
        os.Exit(1)
    }

    pName := strings.TrimSpace(args[0])

    fmtName := pName
    if strings.Contains(fmtName, " ") {
        parts := strings.Split(fmtName, " ")
        var new string
        for i, p := range parts {
            if i == 0 {
                new = strings.ToLower(p)
            } else {
                new = new + cases.Title(language.English).String(p)
            }
        }
        fmtName = new
    }

	if err := os.WriteFile(".gitignore", []byte(gitignore), 0755); err != nil {
        panic(err)
	}

    readme := fmt.Sprintf("# %s", pName) + readmeBody
	if err := os.WriteFile("README.md", []byte(readme), 0755); err != nil {
		panic(err)
	}

	if err := os.WriteFile("main.go", []byte(mainFunc), 0755); err != nil {
		panic(err)
	}

	execCmd := exec.Command("go", "mod", "init", fmtName)
	if err := execCmd.Run(); err != nil {
		panic(err)
	}

    if err := os.Mkdir("src", 0755); err != nil {
        panic(err)
    }

	if (os.Getenv("TERM_PROGRAM") != "vscode" || os.Getenv("VSCODE_RESOLVING_ENVIRONMENT") != "") {
        execCmd = exec.Command("code", ".")
        if err := execCmd.Run(); err != nil {
            panic(err)
        }
	}
}
