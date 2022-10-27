package tests

import (
	"testing"

	"github.com/CS222-UIUC/course-project-group-10.git/cli/data"
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

func TestValidName(t *testing.T) {
	_, err := data.GetCourseByName("Data Structures")

	if err != nil {
		t.Errorf("error when getting valid course")
	}

	_, err = data.GetCourseByName("")

	if err == nil {
		t.Errorf("should not be allowed to get empty course")
	}
}

func TestValidNumber(t *testing.T) {
	_, err := data.GetCourseByNum("CS", 225)

	if err != nil {
		t.Errorf("error when getting valid course")
	}

	_, err = data.GetCourseByNum("", 225)

	if err == nil {
		t.Errorf("should not be allowed to get empty subject")
	}

	_, err = data.GetCourseByNum("CS", -3)

	if err == nil {
		t.Errorf("should not be allowed to have non-three-digit number (too small)")
	}
	_, err = data.GetCourseByNum("CS", 0)

	if err == nil {
		t.Errorf("should not be allowed to have non-three-digit number (too small)")
	}
	_, err = data.GetCourseByNum("CS", 99)

	if err == nil {
		t.Errorf("should not be allowed to have non-three-digit number (too small)")
	}

	_, err = data.GetCourseByNum("CS", 1000)

	if err == nil {
		t.Errorf("should not be allowed to have non-three-digit number (too big)")
	}
}
