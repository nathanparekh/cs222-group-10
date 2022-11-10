package data

import "database/sql"

type Meeting struct {
	Type       string
	StartTime  string
	EndTime    string
	DaysOfWeek string
	RoomNum    string
	Id         string
	SectionId  string
}

func GetMeetingsBySection(section Section) ([]Meeting, error) {
	db, err := sql.Open("sqlite3", "../../python/gpa_dataset.db")

	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM Meeting WHERE section_id=@section_id", section.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Meeting

	for rows.Next() {
		var meeting Meeting
		if err := rows.Scan(
			&meeting.Type,
			&meeting.StartTime,
			&meeting.EndTime,
			&meeting.DaysOfWeek,
			&meeting.RoomNum,
			&meeting.Id,
			&meeting.SectionId,
		); !(err == nil) {
			return nil, err
		}
		result = append(result, meeting)
	}

	err = db.Close()
	if err != nil {
		return nil, err
	}
	return result, nil
}
