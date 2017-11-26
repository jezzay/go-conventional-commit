package main

import (
	"fmt"

	"gitlab.com/jezzay/go-conventional-commit/commit"
)

func main() {
	//message := commit.Parse(`feat: add new feature`)

	//fmt.Printf("Parsed commit message = %v \n", message)

	commitWithBody :=
		`feat: add new feature

Commit message body`
//
//	// TODO: look at how conventional commit parser parses commit messages for inspiration.
//
	message := commit.Parse(commitWithBody)
	fmt.Printf("Parsed commit message with body = \n%v\n", message)
//
//	message = commit.Parse("feat add new feature")
//	fmt.Printf("Invalid commit message = %v \n", message)
}


