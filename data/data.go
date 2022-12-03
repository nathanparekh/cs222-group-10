package data

import (
	"database/sql"
	"errors"
	"reflect"
	"strconv"

	//"errors"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

// global db variable, opened and closed in main
var DB *gorm.DB

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by gorm
func (Course) TableName() string {
	return "course"
}

func (Section) TableName() string {
	return "section"
}

func (Meeting) TableName() string {
	return "meeting"
}

func (GPAEntry) TableName() string {
	return "GPA_Entry"
}

func (Instructor) TableName() string {
	return "instructor"
}

type Course struct {
	Name          string         `db:"name"`
	Year          int            `db:"year"`
	Term          string         `db:"term"`
	Subject       string         `db:"subject"`
	Number        int            `db:"number"`
	Description   string         `db:"description"`
	CreditHours   string         `db:"credit_hours"`
	AC            int            `db:"advanced_comp"`
	NW            int            `db:"non_western"`
	US            int            `db:"us_minority"`
	W             int            `db:"western"`
	HP            int            `db:"hist_phil"`
	LA            int            `db:"lit_arts"`
	LS            int            `db:"life_sci"`
	PS            int            `db:"phys_sci"`
	QR1           int            `db:"quant_res_1"`
	QR2           int            `db:"quant_res_2"`
	BS            int            `db:"behav_sci"`
	SS            int            `db:"social_sci"`
	SectionInfo   sql.NullString `db:"section_info"`
	DegreeAttribs sql.NullString `db:"degree_attribs"`
	ScheduleInfo  sql.NullString `db:"schedule_info"`
}

type Section struct {
	CRN              int            `db:"crn"`
	SectionNumber    sql.NullString `db:"number"`
	StatusCode       string         `db:"status_code"`
	Description      sql.NullString `db:"description"`
	PartOfTerm       string         `db:"part_of_term"`
	SectStatusCode   string         `db:"sect_status_code"`
	EnrollmentStatus string         `db:"enrollment_status"`
	StartDate        string         `db:"start_date"`
	EndDate          string         `db:"end_date"`
	Id               string         `db:"id"`
	CourseId         string         `db:"course_id"`
}

type Meeting struct {
	Type       string         `db:"type"`
	StartTime  string         `db:"start_time"`
	EndTime    sql.NullString `db:"end_time"`
	DaysOfWeek sql.NullString `db:"days_of_week"`
	RoomNum    sql.NullString `db:"room_num"`
	Id         string         `db:"id"`
	SectionId  string         `db:"section_id"`
}

type GPAEntry struct {
	Id        string `db:"id"`
	SchedType string `db:"sched_type"`
	APlus     int    `db:"a_plus"`
	A         int    `db:"a"`
	AMinus    int    `db:"a_minus"`
	BPlus     int    `db:"b_plus"`
	B         int    `db:"b"`
	BMinus    int    `db:"b_minus"`
	CPlus     int    `db:"c_plus"`
	C         int    `db:"c"`
	CMinus    int    `db:"c_minus"`
	DPlus     int    `db:"d_plus"`
	D         int    `db:"d"`
	DMinus    int    `db:"d_minus"`
	F         int    `db:"f"`
	W         int    `db:"w"`
}

type Instructor struct {
	Id        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

// getters take in two arguments:
// example: GetCourses(map[string]interface{}{"subject":"ACE", "number":161,}, "LIMIT 10") will find the first 10 courses where subject=ACE and number=161

// "argsMap" is a map[string]interface{} from column name to value to search for
// to create a map, do mapName := map[string]interface{}{"columnName1": val1, "columnName2": val2,}
// columnName must be a string, but val can be any type
// make sure val is the same type of the column as listed in the DB documentation

// "clauses" is a string with optional additional statements you want to add to the end of a query
// for example, if you want to sort by descending and only return 6 rows, pass
// clauses := "ORDER BY year DESC LIMIT 6"
// if you want no additional statements, just pass an empty string
func GetCourses(argsMap map[string]interface{}, clauses string) ([]Course, error) {
	var courses []Course
	// this initializes the DB
	queryDB := DB.Session(&gorm.Session{})

	queryString := ""
	// build the query as a string, adding each key in the argsMap to the AND statement
	for key, val := range argsMap {
		// do error checking of fields that aren't ints
		// this is the only part of the getter that we really need to edit when writing getters for different structs
		if key != "number" && key != "year" {
			if val == "" {
				return courses, errors.New("field cannot be empty")
			}
		}
		// do error checking for numbers
		if key == "number" {
			var value int
			if reflect.TypeOf(val).String() == "string" {
				value, _ = strconv.Atoi(val.(string))
			} else {
				value = val.(int)
			}
			if value < 0 || value > 799 {
				return courses, errors.New("number out of range")
			}
		}
		// if error checks are successful, add the statement to the query
		// queryString looks something like "columnName1 = @mapKey1 AND columnName2 = @mapKey2"
		// where columnName is the name of the column we want to search and mapKey is the key to the value we want it to be
		// this is equivalent to the statement "SELECT * FROM course WHERE columnName1 = argsMap[mapKey1] AND columnName2 = argsMap[mapKey2]"
		queryString += key + " = @" + key + " AND "
	}
	// remove the trailing " AND ""
	if len(argsMap) > 0 {
		queryString = queryString[0 : len(queryString)-5]
	}
	// append any additional clauses to the end of query
	queryString += " " + clauses
	// fmt.Println(queryString)

	// see section about Named Arguments to see how passing a map[string]interface{} works:
	// https://gorm.io/docs/sql_builder.html#Raw-SQL
	// this code uses the second example snippet

	// .Find(&courses) takes the preceding statement, executes it, and puts it into courses slice
	// nothing is returned
	queryDB.Where(queryString, argsMap).Find(&courses)
	return courses, nil
}

func CoursesToString(courses []Course) string {
	var output string
	for _, course := range courses {
		curr_line := "| " + course.Term + " " + strconv.Itoa(course.Year) + "\t| " + course.Subject + " " + strconv.Itoa(course.Number) + " | " + course.Name + "\n"
		output += curr_line
		for i := 0; i < len(curr_line); i++ {
			output += "-"
		}
		output += "\n"
	}
	return output
}

func SectionsToString(sections []Section) string {
	var output string
	for _, section := range sections {
		curr_line := "| " + strconv.Itoa(section.CRN) +
			"\t| " + section.SectionNumber.String +
			"\t| POT " + section.PartOfTerm +
			"\t| " + section.EnrollmentStatus

		if section.Description.Valid {
			curr_line += "\t| " + section.Description.String
		} else {
			curr_line += "\t|"
		}

		// get instructors here, soon

		curr_line += "\n"
		output += curr_line

		for i := 0; i < len(curr_line); i++ {
			output += "-"
		}

		output += "\n"
	}
	return output
}

func GetSections(argsMap map[string]interface{}, clauses string) ([]Section, error) {
	var sections []Section
	queryDB := DB
	queryString := ""
	for key, val := range argsMap {
		// check if crn is correct length
		if key == "crn" {
			if len(strconv.Itoa(val.(int))) != 5 {
				return []Section{}, errors.New("invalid CRN")
			}
		}
		queryString += key + " = @" + key + " AND "
	}
	// remove the trailing " AND ""
	if len(argsMap) > 0 {
		queryString = queryString[0 : len(queryString)-5]
	}
	// append any additional clauses to the end of query
	queryString += " " + clauses
	subQuery := queryDB.Table("course").Select("id").Where(queryString, argsMap)
	queryDB.Where("course_id = (?)", subQuery).Find(&sections)
	return sections, nil
}

func GetCourseByDatabaseId(databaseId string) ([]Course, error) {
	var courses []Course
	queryDB := DB

	queryDB.Where("id= (?)", databaseId).Limit(1).Find(&courses)

	if len(courses) == 0 {
		return []Course{}, errors.New("could not find course")
	}
	return courses, nil
}

func GetInstructorsBySectionId(sectionId string) ([]Instructor, error) {
	var instructors []Instructor
	queryDB := DB
	subQuery1 := queryDB.Table("Meeting").Select("meeting_id").Where("section_id= (?)", sectionId)
	subQuery2 := queryDB.Table("Class").Select("instructor_id").Where("meeting_id= (?)", subQuery1)
	queryDB.Where("instructor_id= (?)", subQuery2).Find(&instructors)

	return instructors, nil
}

func GetSectionIdSubquery(argsMap map[string]interface{}, clauses string) (*gorm.DB, error) {
	queryDB := DB.Session(&gorm.Session{})
	queryString := ""
	for key, val := range argsMap {
		if key != "number" && key != "year" {
			if val == "" {
				return nil, errors.New("field cannot be empty")
			}
		}
		if key == "number" {
			if val.(int) < 0 || val.(int) > 799 {
				return nil, errors.New("number out of range")
			}
		}
		queryString += key + " = @" + key + " AND "
	}
	if len(argsMap) > 0 {
		queryString = queryString[0 : len(queryString)-5]
	}
	queryString += " " + clauses
	subQuery := queryDB.Select("id").Where(queryString, argsMap)
	return subQuery, nil
}

//func GetMeetings(argsMap map[string]interface{}, clauses string) ([]Meeting, error) {
//	var meetings []Meeting
//	queryDB := DB
//	queryString := ""
//	var subQuery *gorm.DB
//	for key := range argsMap {
//		if key == "crn" {
//			subQuery, _ = GetSectionIdSubquery(argsMap, clauses)
//		}
//		queryString += key + " = @" + key + " AND "
//	}
//	if len(argsMap) > 0 {
//		queryString = queryString[0 : len(queryString)-5]
//	}
//	queryString += " " + clauses
//	queryDB.Where("section_id=(?)", subQuery).Find(&meetings)
//	return meetings, nil
//}

//func GetInstructors(argsMap map[string]interface{}, clauses string) ([]Instructor, error) {
//	var instructors []Instructor
//	queryDB := DB
//	queryString := ""
//	for key := range argsMap {
//		queryString += key + " = @" + key + " AND "
//	}
//	if len(argsMap) > 0 {
//		queryString = queryString[0 : len(queryString)-5]
//	}
//	queryString += " " + clauses
//	queryDB.Where(queryString, argsMap).Find(&instructors)
//	return instructors, nil
//}
