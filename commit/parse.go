package commit

import (
	"regexp"
	"strings"
)

// feat: added a new feature

// ->
// 	type: feat
//	description: added a new feature

type Parsed struct {
	header      Header
	body        string
}

func Parse(c string) Parsed {

	lines := regexp.MustCompile("\r?\n").Split(c, -1)

	if len(lines) >= 1 {
		header := parseHeader(lines[0])
		body := strings.Join(lines[1:], "")
		return Parsed{header, body}
	}
	return Parsed{}
}

