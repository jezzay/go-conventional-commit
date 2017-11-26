package commit

import (
	"regexp"
	"strings"
)

type Parsed struct {
	header Header
	body   string
	notes  []Note
}

type Note struct {
	title string
	text  string
}

func Parse(c string) Parsed {

	lines := regexp.MustCompile("\r?\n").Split(c, -1)

	if len(lines) >= 1 {
		header := parseHeader(lines[0])
		body := strings.Join(lines[1:], "")

		notesRegex := regexp.MustCompile(`^(?i)[\s|*]*(BREAKING CHANGE)[:\s]+(.*)`)
		notes := make([]Note, 0, 1)

		for _, l := range lines[1:] {
			if notesRegex.MatchString(l) {
				matches := notesRegex.FindAllStringSubmatch(l, -1)
				if len(matches) == 1 {
					parts := matches[0]
					note := Note{parts[1], parts[2]}
					notes = append(notes, note)
				}
			}
		}
		return Parsed{header, body, notes}
	}
	return Parsed{}
}
