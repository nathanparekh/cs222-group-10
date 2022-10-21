package tests

import (
	"database/sql"
	//"fmt"
	"testing"

	"github.com/CS222-UIUC/course-project-group-10.git/cli/data"

	_ "github.com/mattn/go-sqlite3"
)

func TestCanOpenDB(t *testing.T) {
	db, err := sql.Open("sqlite3", "../../python/gpa_dataset.db")

	if err != nil {
		t.Errorf("could not open database")
	}

	db.Close()
}

func TestCourseDataIntegrity(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../python/gpa_dataset.db")
	courses, err := db.Query("SELECT * FROM courses")
	if err != nil {
		t.Errorf("couldn't query courses")
	}

	cols, err := courses.Columns()

	if err != nil {
		t.Errorf("error parsing columns")
	}

	if len(cols) != 6 {
		t.Errorf("incorrect number of columns")
	}
	db.Close()
}

func TestGetCourseByNum(t *testing.T) {
	var course data.Course = data.GetCourseByNum("CS", 225)

	if course.Name != "Data Structures" {
		t.Errorf("incorrect name")
	}
	if course.Year != 2021 {
		t.Errorf("incorrect year")
	}
	if course.Term != "Fall" {
		t.Errorf("incorrect term")
	}
	if course.Subject != "CS" {
		t.Errorf("incorrect subject")
	}
	if course.Number != 225 {
		t.Errorf("incorrect name")
	}
}

func TestGetCourseByName(t *testing.T) {
	var course data.Course = data.GetCourseByName("Data Structures")

	if course.Name != "Data Structures" {
		t.Errorf("incorrect name")
	}
	if course.Year != 2021 {
		t.Errorf("incorrect year")
	}
	if course.Term != "Fall" {
		t.Errorf("incorrect term")
	}
	if course.Subject != "CS" {
		t.Errorf("incorrect subject")
	}
	if course.Number != 225 {
		t.Errorf("incorrect name")
	}
}
