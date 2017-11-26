package commit

import "testing"

func TestParseHeaderRaw(t *testing.T) {
	commitHeader := `feat: add new feature`
	header := parseHeader(commitHeader)
	expected := "feat: add new feature"
	if header.raw != expected {
		t.Errorf(`Expected header.raw to equal %v got %v`, expected, header.raw)
	}
}

func TestParseHeaderType(t *testing.T) {
	commitHeader := `feat: add new feature`
	header := parseHeader(commitHeader)
	expected := "feat"
	if header.commitType != expected {
		t.Errorf(`Expected header.commitType to equal %v got %v`, expected, header.commitType)
	}
}

func TestParseHeaderNoScope(t *testing.T) {
	commitHeader := `fix: add new feature`
	header := parseHeader(commitHeader)
	expected := ""
	if header.scope != expected {
		t.Errorf(`Expected header.scope to equal %v got %v`, expected, header.scope)
	}
}

func TestParseHeaderScope(t *testing.T) {
	commitHeader := `feat(scope): add new feature`
	header := parseHeader(commitHeader)
	expected := "scope"
	if header.scope != expected {
		t.Errorf(`Expected header.scope to equal %v got %v`, expected, header.scope)
	}
}

func TestParseHeaderInvalid(t *testing.T) {
	commitHeader := `add a new feature`
	header := parseHeader(commitHeader)
	if header.raw != "" {
		t.Errorf(`Expected header to be empty, got %v`, header)
	}
}

