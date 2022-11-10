package data

import "database/sql"

type Instructor struct {
	Id string
	FirstName string
	LastName string
}

func GetInstructorsByMeeting(meeting Meeting) ([]Instructor, error){
	db, err := sql.Open("sqlite3", "../../python/gpa_dataset.db")

	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, first_name, last_name FROM Instructor WHERE id IN (SELECT instructor_id FROM Class WHERE meeting_id=@meetingId)", meeting.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Instructor

	for rows.Next() {
		var instructor Instructor
		if err := rows.Scan(
			&instructor.Id,
			&instructor.FirstName,
			&instructor.LastName,
		); !(err == nil) {
			return nil, err
		}
		result = append(result, instructor)
	}

	err = db.Close()
	if err != nil {
		return nil, err
	}
	return result, nil
}
