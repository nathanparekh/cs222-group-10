package data

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	//"errors"
	"github.com/gookit/color"
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
	Id            string         `gorm:"primaryKey"`
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
	Id               int            `gorm:"primaryKey"`
	CourseId         int            `db:"course_id"`
	CRN              int            `db:"crn"`
	SectionNumber    string         `db:"section_number"`
	StatusCode       string         `db:"status_code"`
	SectStatusCode   string         `db:"sect_status_code"`
	PartOfTerm       string         `db:"part_of_term"`
	EnrollmentStatus string         `db:"enrollment_status"`
	StartDate        string         `db:"start_date"`
	EndDate          string         `db:"end_date"`
	Description      sql.NullString `db:"description"`
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
	Id           string `gorm:"primaryKey"`
	CourseId     string
	InstructorId string
	SchedType    string `db:"sched_type"`
	APlus        int    `db:"a_plus"`
	A            int    `db:"a"`
	AMinus       int    `db:"a_minus"`
	BPlus        int    `db:"b_plus"`
	B            int    `db:"b"`
	BMinus       int    `db:"b_minus"`
	CPlus        int    `db:"c_plus"`
	C            int    `db:"c"`
	CMinus       int    `db:"c_minus"`
	DPlus        int    `db:"d_plus"`
	D            int    `db:"d"`
	DMinus       int    `db:"d_minus"`
	F            int    `db:"f"`
	W            int    `db:"w"`
}

type Instructor struct {
	Id        string `gorm:"primaryKey"`
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
				return courses, errors.New("GetCourses field \"" + key + "\" cannot be empty")
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
	queryDB.Where(queryString, argsMap).Find(&sections)
	return sections, nil
}

func GetInstructors(argsMap map[string]interface{}, clauses string) ([]Instructor, error) {
	var instructors []Instructor
	// this initializes the DB
	queryDB := DB.Session(&gorm.Session{})

	queryString := ""
	// build the query as a string, adding each key in the argsMap to the AND statement
	for key := range argsMap {
		queryString += key + " = @" + key + " AND "
	}
	// remove the trailing " AND ""
	if len(argsMap) > 0 {
		queryString = queryString[0 : len(queryString)-5]
	}
	// append any additional clauses to the end of query
	queryString += " " + clauses
	queryDB.Where(queryString, argsMap).Find(&instructors)
	return instructors, nil
}
func GetGPAEntry(argsMap map[string]interface{}, clauses string) ([]GPAEntry, error) {
	var gpas []GPAEntry
	// this initializes the DB
	queryDB := DB.Session(&gorm.Session{})

	queryString := ""
	// build the query as a string, adding each key in the argsMap to the AND statement
	for key := range argsMap {
		queryString += key + " = @" + key + " AND "
	}
	// remove the trailing " AND ""
	if len(argsMap) > 0 {
		queryString = queryString[0 : len(queryString)-5]
	}
	// append any additional clauses to the end of query
	queryString += " " + clauses
	if len(argsMap) > 0 {
		queryDB.Where(queryString, argsMap).Find(&gpas)
	} else if len(clauses) > 0 {
		// if argsmap is empty, assume using just raw sql
		rows, err := queryDB.Raw(queryString).Rows()
		if err != nil {
			return gpas, err
		}
		defer rows.Close()
		for rows.Next() {
			var gpa GPAEntry
			queryDB.ScanRows(rows, &gpa)
			gpas = append(gpas, gpa)
		}
	} else {
		return gpas, errors.New("must pass non-empty arguments")
	}

	return gpas, nil
}

func CoursesToString(courses []Course) string {
	var output string
	hline := "------------------------------------------------------\n"
	for _, course := range courses {
		curr_line := "| " + course.Term + " " + strconv.Itoa(course.Year) + "\t| " + course.Subject + " " + strconv.Itoa(course.Number) + " | " + course.Name + "\n"
		output += curr_line + hline
	}
	return output
}
func PrintGpas(gpas []GPAEntry) {
	if len(gpas) == 0 {
		return
	}

	var hline string = "----------------------------------------------------------------------------------------------------------------------------------------------------------\n"
	fmt.Print("| Term\t\t| A+\t| A\t| A-\t| B+\t| B\t| B-\t| C+\t| C\t| C-\t| D+\t| D\t| D-\t| F\t| W\t| Avg\t| Instructor\n" + hline)
	for _, entry := range gpas {
		course, err := GetCourses(map[string]interface{}{"id": entry.CourseId}, "")
		if err != nil {
			fmt.Println(err)
			return
		}
		instructor, err := GetInstructors(map[string]interface{}{"id": entry.InstructorId}, "")
		if err != nil {
			fmt.Println(err)
			return
		}
		size := GetClassSize(entry)
		g := reflect.ValueOf(entry)
		fmt.Print("| " + course[0].Term + " " + strconv.Itoa(course[0].Year) + "\t|")
		// average gpa
		var gpa float64
		// iterate through the fields of the struct (APlus is the 4th field, W is the last)
		for i := 4; i < g.NumField(); i++ {
			// print the percentage (float64)
			percentage := float64(g.Field(i).Int()) / size * 100
			// add up gpa points (as we increment i, each letter grade is 0.33 less than the previous one, offset by 4 since we start at i=4
			// multiply it by the number of students with that grade, we divide by total students at the very end)
			gpa += (4.0 - ((float64(i) - 4.0) * 0.33)) * float64(g.Field(i).Int())
			if percentage <= 0 {
				color.Gray.Printf(strconv.FormatFloat(percentage, 'f', 2, 32))
			} else if percentage < 10 {
				color.Red.Printf(strconv.FormatFloat(percentage, 'f', 2, 32))
			} else if percentage < 40 {
				color.Yellow.Printf(strconv.FormatFloat(percentage, 'f', 2, 32))
			} else {
				color.Green.Printf(strconv.FormatFloat(percentage, 'f', 2, 32))
			}
			fmt.Print("%\t|")
		}
		gpa = gpa / size
		fmt.Print(strconv.FormatFloat(gpa, 'f', 2, 64) + "\t| " + instructor[0].FirstName + " " + instructor[0].LastName + "\n" + hline)

	}
}

// given a gpa entry, return the total number of students who took that class
func GetClassSize(x GPAEntry) float64 {
	return float64(x.APlus + x.A + x.AMinus + x.BPlus + x.B + x.BMinus + x.CPlus + x.C + x.CMinus + x.DPlus + x.D + x.DMinus + x.F + x.W)
}
