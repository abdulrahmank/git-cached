package parser

import "testing"

func TestParse(t *testing.T) {
	contents, e := Parse("/Users/kabdul/go/src/github.com/abdulrahmank/git-cached/test_resource/.gitignore")
	if e != nil {
		t.Errorf("expected no error but got %v", e)
	}

	expectedContents := StringList{"bin/", "/release/", ".idea"}

	if !expectedContents.compare(contents) {
		t.Errorf("Expected %v but was %v", expectedContents, contents)
	}
}