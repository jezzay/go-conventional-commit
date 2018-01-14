package main

import (
	"fmt"

	"github.com/jezzay/go-conventional-commit/commit"
)

func main() {

	// Example commit message; commit.Parse will return the different parts of the commit message

	commitMsg :=
		`feat: add new feature
Description of the new feature

BREAKING CHANGE: A new breaking change
Details on the breaking change

Closes #42`
	message := commit.Parse(commitMsg)

	fmt.Printf("Parsed commit = \n%+v\n", message)
}
