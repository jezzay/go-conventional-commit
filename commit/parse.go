package commit

import (
	"regexp"
)

type Parsed struct {
	header Header
	body   string
	footer string
	notes  []Note
}

type Note struct {
	title string
	text  string
}

func Parse(c string) Parsed {

	lines := regexp.MustCompile("\r?\n").Split(c, -1)

	body := ""
	footer := ""
	continueNote := false
	isBody := true
	if len(lines) >= 1 {
		header := parseHeader(lines[0])
		//body = strings.Join(lines[1:], "")
		notesRegex := regexp.MustCompile(`^(?i)[\s|*]*(BREAKING CHANGE)[:\s]+(.*)`)
		notes := make([]Note, 0, 1)

		for _, l := range lines[1:] {
			if notesRegex.MatchString(l) {
				matches := notesRegex.FindAllStringSubmatch(l, -1)
				if len(matches) == 1 {

					isBody = false
					continueNote = true
					footer = appendLine(footer, l)

					parts := matches[0]
					note := Note{parts[1], parts[2]}
					notes = append(notes, note)
					continue
				}
			}
			if continueNote {
				previousNote := len(notes) - 1
				notes[previousNote].text = notes[previousNote].text + l
				footer = appendLine(footer, l)
				continue
				//	footer = append(footer, line);
			}
			if isBody {
				body = appendLine(body, l)
			}

		}
		return Parsed{header, body, footer, notes}
	}
	return Parsed{}
}

func appendLine(src string, line string) string {
	if len(src) >= 1 {
		return src + "\n" + line
	}
	return line
}
