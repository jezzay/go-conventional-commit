package main

import (
	"fmt"

	"github.com/jezzay/go-conventional-commit/parse"
)

func main() {
	message := parse.Parse("feat: add new feature")
	fmt.Printf("Parsed commit message = %v \n", message)

	commitWithBody :=
		`feat: add new feature

Commit message body`

	// TODO: look at how conventional commit parser parses commit messages for inspiration.

	message = parse.Parse(commitWithBody)
	fmt.Printf("Parsed commit message with body = \n%v\n", message)

	message = parse.Parse("feat add new feature")
	fmt.Printf("Invalid commit message = %v \n", message)
}


