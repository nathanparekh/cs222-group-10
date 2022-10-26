package data

import (
	"database/sql"
	"fmt"

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

func GetCourseByNum(subject string, num int) Course {
	db, err := sql.Open("sqlite3", "../../python/gpa_dataset.db")

	if err != nil {
		fmt.Println(err)
	}
	var course Course
	if err := db.QueryRow("SELECT * FROM courses WHERE Subject=@subject AND Number=@num", subject, num).Scan(&course.Year, &course.Term, &course.YearTerm, &course.Subject, &course.Number, &course.Name); err != nil {
		fmt.Println(err)
	}
	db.Close()
	return course
}

func GetCourseByName(name string) Course {
	db, err := sql.Open("sqlite3", "../../python/gpa_dataset.db")

	if err != nil {
		fmt.Println(err)
	}
	var course Course
	if err := db.QueryRow("SELECT * FROM courses WHERE `Course Title`=@name", name).Scan(&course.Year, &course.Term, &course.YearTerm, &course.Subject, &course.Number, &course.Name); err != nil {
		fmt.Println(err)
	}
	db.Close()
	return course
}
