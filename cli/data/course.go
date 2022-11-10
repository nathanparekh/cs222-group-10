package data

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type Course struct {
	Year          int
	Term          string
	Subject       string
	Number        int
	Name          string
	Description   string
	CreditHours   string
	AdvancedComp  int
	NonWestern    int
	USMinority    int
	Western       int
	HistPhil      int
	LitArts       int
	LifeSci       int
	PhysSci       int
	QuantRes1     int
	QuantRes2     int
	BehavSci      int
	SocialSci     int
	Id            string
	SectionInfo   string
	DegreeAttribs string
	SchdeduleInfo string
}

func GetCourseByNum(subject string, num int) (Course, error) {
	db, err := sql.Open("sqlite3", "../../python/course.db")

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
	if err := db.QueryRow("SELECT year,term,subject,number,name,description,credit_hours,advanced_comp,non_western,us_minority,western,hist_phil,lit_arts,life_sci,phys_sci,quant_res_1, quant_res_2,behav_sci,social_sci,id,section_info,degree_attribs,quant_res_1,schedule_info FROM Course WHERE subject=@subject AND number=@num", subject, num).Scan(
		&course.Year,
		&course.Term,
		&course.Subject,
		&course.Number,
		&course.Name,
		&course.Description,
		&course.CreditHours,
		&course.AdvancedComp,
		&course.NonWestern,
		&course.USMinority,
		&course.Western,
		&course.HistPhil,
		&course.LitArts,
		&course.LifeSci,
		&course.PhysSci,
		&course.QuantRes1,
		&course.QuantRes2,
		&course.BehavSci,
		&course.SocialSci,
		&course.Id,
		&course.SectionInfo,
		&course.DegreeAttribs,
		&course.SchdeduleInfo,
	); err != nil {
		return Course{}, errors.New(err.Error())
	}
	db.Close()
	return course, nil
}

func GetCourseByName(name string) (Course, error) {
	db, err := sql.Open("sqlite3", "../../python/course.db")

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
	if err := db.QueryRow("SELECT year,term,subject,number,name,description,credit_hours,advanced_comp,non_western,us_minority,western,hist_phil,lit_arts,life_sci,phys_sci,quant_res_1,quant_res_2, behav_sci,social_sci,id,section_info,degree_attribs,schedule_info FROM Course WHERE name=@name", name).Scan(
		&course.Year,
		&course.Term,
		&course.Subject,
		&course.Number,
		&course.Name,
		&course.Description,
		&course.CreditHours,
		&course.AdvancedComp,
		&course.NonWestern,
		&course.USMinority,
		&course.Western,
		&course.HistPhil,
		&course.LitArts,
		&course.LifeSci,
		&course.PhysSci,
		&course.QuantRes1,
		&course.QuantRes2,
		&course.BehavSci,
		&course.SocialSci,
		&course.Id,
		&course.SectionInfo,
		&course.DegreeAttribs,
		&course.SchdeduleInfo,
	); err != nil {
		return Course{}, errors.New(err.Error())
	}
	db.Close()
	return course, nil
}
