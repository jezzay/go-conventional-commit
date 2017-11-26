package commit

import "regexp"


type Header struct {
	raw         string
	commitType  string
	scope       string
	description string
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

