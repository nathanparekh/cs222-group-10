package tests

import (
	"database/sql"
	"testing"

	"github.com/CS222-UIUC/course-project-group-10.git/cli/data"

	_ "github.com/mattn/go-sqlite3"
)

func TestCanOpenDB(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../python/course.db")
	err := db.Ping()
	if err != nil {
		t.Errorf("could not open database")
	}

	db.Close()
}

func TestCourseDataIntegrity(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../python/course.db")
	courses, err := db.Query("SELECT * FROM Course")
	if err != nil {
		t.Errorf("couldn't query courses")
	}

	cols, err := courses.Columns()
	if err != nil {
		t.Errorf("error parsing columns")
	}

	if len(cols) != 23 {
		t.Errorf("incorrect number of columns")
	}
	db.Close()
}

func TestGetCourseByNum(t *testing.T) {
	course, err := data.GetCourseByNum("ACE", 161)
	if err != nil {
		t.Errorf("error when getting course")
	}
	if course.Name != "Microcomputer Applications" {
		t.Errorf("incorrect name")
	}
	if course.Year != 2022 {
		t.Errorf("incorrect year")
	}
	if course.Term != "Winter" {
		t.Errorf("incorrect term")
	}
	if course.Subject != "ACE" {
		t.Errorf("incorrect subject")
	}
	if course.Number != 161 {
		t.Errorf("incorrect name")
	}
}

func TestGetCourseByName(t *testing.T) {
	course, err := data.GetCourseByName("Microcomputer Applications")
	if err != nil {
		t.Errorf("error when getting course")
	}
	if course.Name != "Microcomputer Applications" {
		t.Errorf("incorrect name")
	}
	if course.Year != 2022 {
		t.Errorf("incorrect year")
	}
	if course.Term != "Winter" {
		t.Errorf("incorrect term")
	}
	if course.Subject != "ACE" {
		t.Errorf("incorrect subject")
	}
	if course.Number != 161 {
		t.Errorf("incorrect name")
	}
}
