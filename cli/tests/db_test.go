package tests

import (
	"database/sql"
	"testing"

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
	db, _ := sql.Open("sqlite3", "python/gpa_dataset.db")
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
