package data

import (
	"database/sql"
	"errors"
)

type Section struct {
	Id               int
	CourseId         int
	CRN              int
	SectionNumber    string
	StatusCode       string
	SectStatusCode   string
	PartOfTerm       string
	EnrollmentStatus string
	StartDate        string
	EndDate          string
}

func GetSectionByCRN(crn int) (Section, error) {
	db, err := sql.Open("sqlite3", "../../python/gpa_dataset.db")

	if crn > 999999 {
		return Section{}, errors.New("CRN larger than 6 digits")
	}
	if !(err == nil) {
		return Section{}, err
	}

	var section Section
	if err := db.QueryRow("SELECT * from sections WHERE CRN=@crn", crn).Scan(
		&section.Id, &section.CourseId, &section.SectionNumber, &section.StatusCode,
		&section.SectStatusCode, &section.PartOfTerm, &section.EnrollmentStatus,
		&section.StartDate, &section.EndDate,
	); !(err == nil) {
		return Section{}, err
	}

	db.Close()
	return section, nil
}
