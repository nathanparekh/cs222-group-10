package tests

import (
	"testing"

	// "github.com/CS222-UIUC/course-project-group-10.git/cli/data"
	_ "github.com/mattn/go-sqlite3"
)

// This is what a test would look like for our CLI
func TestCoursesCommand(t *testing.T) {
	// we would call the function that runs in the CLI command
	// commandOut := courses_command()

	// and check if it has the desired output
	if false {
		t.Errorf("command doesn't match desired output")
	}
}
