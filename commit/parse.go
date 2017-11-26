package commit

import (
	"fmt"
	"regexp"
)

// feat: added a new feature

// ->
// 	type: feat
//	description: added a new feature

type Parsed struct {
	header      Header
	description string
	body        string
}

type Header struct {
	raw         string
	commitType  string
	scope       string
	description string
}

func Parse(c string) Parsed {

	lines := regexp.MustCompile("\r?\n").Split(c, -1)

	fmt.Printf("%+q\n", lines)

	if len(lines) > 1 {
		header := parseHeader(lines[0])
		fmt.Printf("Parsed commit message with body = \n%v\n", header)
		return Parsed{header, "", ""}
	}
	return Parsed{}
}

func parseHeader(h string) Header {
	headerPattern := regexp.MustCompile(`^(\w*)(?:\((.*)\))?: (.*)$`)
	if headerPattern.MatchString(h) {
		matches := headerPattern.FindAllStringSubmatch(h, -1)
		if len(matches) == 1 {
			parts := matches[0]
			raw := parts[0]
			typ := parts[1]
			scope := parts[2]
			desc := parts[3]
			return Header{raw, typ, scope, desc}
		}
	}
	return Header{}
}
