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

func TestParseCommitWithNote(t *testing.T) {
	commit := `feat: add new feature
Description of the new feature
BREAKING CHANGE: A new breaking change`
	parsedCommit := Parse(commit)
	if len(parsedCommit.notes) != 1 {
		t.Errorf(`Expected len(parsedCommit.Notes) to equal %v got %v`, 1, parsedCommit.notes)
	}
}

func TestParseCommitWithNoteInLowercase(t *testing.T) {
	commit := `feat: add new feature
Description of the new feature
breaking change: A new breaking change`
	parsedCommit := Parse(commit)
	if len(parsedCommit.notes) != 1 {
		t.Errorf(`Expected len(parsedCommit.Notes) to equal %v got %v`, 1, parsedCommit.notes)
	}
}

func TestParseCommitWithNoteTitle(t *testing.T) {
	commit := `feat: add new feature
Description of the new feature
BREAKING CHANGE: A new breaking change`
	parsedCommit := Parse(commit)
	if len(parsedCommit.notes) == 1 {
		note := parsedCommit.notes[0]
		if note.title != "BREAKING CHANGE" {
			t.Errorf(`Expected parsedCommit.Notes[0].title to equal %v got %v`, "BREAKING CHANGE", note.title)
		}
	} else {
		t.Error("Expected one Note")
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
