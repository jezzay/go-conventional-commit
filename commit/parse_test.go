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

func TestParseCommitWithMultiLineBody(t *testing.T) {
	commit := `feat: add new feature
Description of the new feature
more details
even more details`
	parsedCommit := Parse(commit)
	expected := `Description of the new feature
more details
even more details`
	if parsedCommit.body != expected {
		t.Errorf("Expected parsedCommit.body to equal \n%v \n\ngot:\n\n%v", expected, parsedCommit.body)
	}
}

func TestParseCommitBodyExcludesFooter(t *testing.T) {
	commit := `feat: add new feature
Description of the new feature
more details
BREAKING CHANGE: A new breaking change
Closes #1`
	parsedCommit := Parse(commit)
	expected := `Description of the new feature
more details`
	if parsedCommit.body != expected {
		t.Errorf("Expected parsedCommit.body to equal \n%v \n\ngot:\n\n%v", expected, parsedCommit.body)
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

func TestParseCommitWithNoteOnNewLine(t *testing.T) {
	commit := `feat: add new feature
Description of the new feature
BREAKING CHANGE:
A new breaking change on a new line`
	parsedCommit := Parse(commit)
	if len(parsedCommit.notes) == 1 {
		note := parsedCommit.notes[0]
		expected := "A new breaking change on a new line"
		if note.text != expected {
			t.Errorf(`Expected parsedCommit.Notes[0].text to equal %v got %v`, expected, note.title)
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

func TestParseCommitFooter(t *testing.T)  {
	commit := `feat: add new feature
Description of the new feature

BREAKING CHANGE: A new breaking change
Details on the breaking change

Closes #42`

	parsedCommit := Parse(commit)
	expected := `BREAKING CHANGE: A new breaking change
Details on the breaking change

Closes #42`

	if parsedCommit.footer != expected {
		t.Errorf(`Expected parsedCommit.footer to equal %v got %v`, expected, parsedCommit.footer)
	}

}
