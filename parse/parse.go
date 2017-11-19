package parse

import (
	"fmt"
	"strings"
)

// feat: added a new feature

// ->
// 	type: feat
//	description: added a new feature

type CommitMessage struct {
	commitType  string
	description string
	body        string
}

func Parse(c string) CommitMessage {
	split := strings.Split(c, "\n")
	fmt.Printf("%+q\n", split)
	if len(split) > 1 {
		return CommitMessage{split[0], split[1], ""}
	}
	return CommitMessage{}
}
