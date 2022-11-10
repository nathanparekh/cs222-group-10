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
	Description      string
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
	if err := db.QueryRow("SELECT crn,number,status_code,description,part_of_term,sect_status_code,enrollment_status,start_date,end_date,id,course_id FROM Section WHERE CRN=@crn", crn).Scan(
		&section.Id,
		&section.CourseId,
		&section.CRN,
		&section.SectionNumber,
		&section.StatusCode,
		&section.SectStatusCode,
		&section.PartOfTerm,
		&section.EnrollmentStatus,
		&section.StartDate,
		&section.EndDate,
		&section.Description,
	); !(err == nil) {
		return Section{}, err
	}

	db.Close()
	return section, nil
}

func GetSectionsByCourse(course Course) ([]Section, error) {
	db, err := sql.Open("sqlite3", "../../python/gpa_dataset.db")

	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT crn,number,status_code,description,part_of_term,sect_status_code,enrollment_status,start_date,end_date,id,course_id FROM Section WHERE course_id=@course_id", course.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Section

	for rows.Next() {
		var section Section
		if err := rows.Scan(
			&section.Id,
			&section.CourseId,
			&section.CRN,
			&section.SectionNumber,
			&section.StatusCode,
			&section.SectStatusCode,
			&section.PartOfTerm,
			&section.EnrollmentStatus,
			&section.StartDate,
			&section.EndDate,
			&section.Description,
		); !(err == nil) {
			return nil, err
		}
		result = append(result, section)
	}

	err = db.Close()
	if err != nil {
		return nil, err
	}
	return result, nil
}
