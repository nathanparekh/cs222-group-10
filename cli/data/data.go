package data

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type Course struct {
	Name     string
	Year     int
	Term     string
	YearTerm string
	Subject  string
	Number   int
}

func GetCourseByNum(subject string, num int) (Course, error) {
	db, err := sql.Open("sqlite3", "../python/gpa_dataset.db")

	if subject == "" {
		return Course{}, errors.New("empty subject")
	}

	if num < 0 || num > 799 {
		return Course{}, errors.New("number out of range")
	}

	if subject == "" {
		return Course{}, errors.New("empty subject")
	}

	if num < 0 || num > 799 {
		return Course{}, errors.New("number out of range")
	}

	if err != nil {
		return Course{}, errors.New(err.Error())
	}
	var course Course
	if err := db.QueryRow("SELECT * FROM courses WHERE Subject=@subject AND Number=@num", subject, num).Scan(&course.Year, &course.Term, &course.YearTerm, &course.Subject, &course.Number, &course.Name); err != nil {
		return Course{}, errors.New(err.Error())
	}
	db.Close()
	return course, nil
}

func GetCourseByName(name string) (Course, error) {
	db, err := sql.Open("sqlite3", "../python/gpa_dataset.db")

	if name == "" {
		return Course{}, errors.New("empty course name")
	}

	if name == "" {
		return Course{}, errors.New("empty course name")
	}

	if err != nil {
		return Course{}, errors.New(err.Error())
	}
	var course Course
	if err := db.QueryRow("SELECT * FROM courses WHERE `Course Title`=@name", name).Scan(&course.Year, &course.Term, &course.YearTerm, &course.Subject, &course.Number, &course.Name); err != nil {
		return Course{}, errors.New(err.Error())
	}
	db.Close()
	return course, nil
}
