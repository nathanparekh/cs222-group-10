package data

import (
	"database/sql"
	"errors"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Course struct {
	Name    string
	Year    int
	Term    string
	Subject string
	Number  int
}

func GetCoursesByNum(subject string, num int, limit int) ([]Course, error) {
	db, err := sql.Open("sqlite3", "../python/course.db")

	if subject == "" {
		return []Course{}, errors.New("empty subject")
	}

	if num < 0 || num > 799 {
		return []Course{}, errors.New("number out of range")
	}

	if subject == "" {
		return []Course{}, errors.New("empty subject")
	}

	if num < 0 || num > 799 {
		return []Course{}, errors.New("number out of range")
	}

	if err != nil {
		return []Course{}, errors.New(err.Error())
	}
	var courses []Course
	rows, err := db.Query("SELECT year, term, subject, number, name FROM Course WHERE subject=@subject AND number=@num LIMIT @limit", subject, num, limit)
	if err != nil {
		return courses, err
	}
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.Year, &course.Term, &course.Subject, &course.Number, &course.Name); err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}
	rows.Close()
	db.Close()
	return courses, nil
}

func GetCoursesByName(name string, limit int) ([]Course, error) {
	db, err := sql.Open("sqlite3", "../python/course.db")

	if name == "" {
		return []Course{}, errors.New("empty course name")
	}

	if name == "" {
		return []Course{}, errors.New("empty course name")
	}

	if err != nil {
		return []Course{}, errors.New(err.Error())
	}

	var courses []Course
	rows, err := db.Query("SELECT year, term, subject, number, name FROM Course WHERE `name`=@name LIMIT @limit", name, limit)
	if err != nil {
		return courses, err
	}
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.Year, &course.Term, &course.Subject, &course.Number, &course.Name); err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}
	rows.Close()
	db.Close()
	return courses, nil
}

func CourseToString(course Course) string {
	var output string = "| " + course.Term + " " + strconv.Itoa(course.Year) + "\t| " + course.Subject + " " + strconv.Itoa(course.Number) + " | " + course.Name
	return output
}
