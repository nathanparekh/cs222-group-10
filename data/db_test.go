package data

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCanOpenDB(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../python/course.db")
	err := db.Ping()
	if err != nil {
		t.Errorf("could not open database")
	}

	db.Close()
}

func TestCourseDataIntegrity(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../python/course.db")
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
	course, err := GetCourses(map[string]interface{}{"subject":"ACE", "number":161,}, "LIMIT 1")
	if err != nil {
		t.Errorf("error when getting course")
		fmt.Println(err)
	}
	if course[0].Name != "Microcomputer Applications" {
		t.Errorf("incorrect name")
	}
	if course[0].Year != 2022 {
		t.Errorf("incorrect year")
	}
	if course[0].Term != "Winter" {
		t.Errorf("incorrect term")
	}
	if course[0].Subject != "ACE" {
		t.Errorf("incorrect subject")
	}
	if course[0].Number != 161 {
		t.Errorf("incorrect name")
	}
}

func TestGetCourseByName(t *testing.T) {
	course, err := GetCourses(map[string]interface{}{"name":"Microcomputer Applications",}, "LIMIT 1")
	if err != nil {
		t.Errorf("error when getting course")
		fmt.Println(err)
	}
	if course[0].Name != "Microcomputer Applications" {
		t.Errorf("incorrect name")
	}
	if course[0].Year != 2022 {
		t.Errorf("incorrect year")
	}
	if course[0].Term != "Winter" {
		t.Errorf("incorrect term")
	}
	if course[0].Subject != "ACE" {
		t.Errorf("incorrect subject")
	}
	if course[0].Number != 161 {
		t.Errorf("incorrect name")
	}
}
