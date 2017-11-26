package commit

import "testing"

func TestParseCommitWithBody(t *testing.T) {
	commit := `feat: add new feature
Description of the new feature`
	parsedCommit := Parse(commit)
	expected := "Description of the new feature"
	if parsedCommit.body != expected {
		t.Errorf(`Expected parsedCommit.body to equal %v got %v`, expected, parsedCommit.body)
	}
}

func TestParseCommitWithOutBody(t *testing.T) {
	commit := `feat: add new feature`
	parsedCommit := Parse(commit)
	expected := ""
	if parsedCommit.body != expected {
		t.Errorf(`Expected parsedCommit.body to equal %v got %v`, expected, parsedCommit.body)
	}
}
