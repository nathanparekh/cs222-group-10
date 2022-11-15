package data

import (
	"fmt"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestValidName(t *testing.T) {
	var err error
	DB, err = gorm.Open(sqlite.Open("../python/course.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	_, err = GetCourses(map[string]interface{}{"name": "Microcomputer Applications"}, "LIMIT 1")
	if err != nil {
		t.Errorf("error when getting valid course")
		fmt.Println(err)
	}

	_, err = GetCourses(map[string]interface{}{"name": ""}, "LIMIT 1")

	if err == nil {
		t.Errorf("should not be allowed to get empty course")
	}

	_, err = GetCourses(map[string]interface{}{}, "")

	if err != nil {
		t.Errorf("should be allowed to query with no qualifiers")
	}
}

func TestValidNumber(t *testing.T) {
	var err error
	DB, err = gorm.Open(sqlite.Open("../python/course.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	_, err = GetCourses(map[string]interface{}{"subject": "ACE", "number": 161}, "LIMIT 1")
	if err != nil {
		t.Errorf("error when getting valid course")
		fmt.Println(err)
	}

	_, err = GetCourses(map[string]interface{}{"subject": "", "number": 161}, "LIMIT 10")

	if err == nil {
		t.Errorf("should not be allowed to get empty subject")
	}

	_, err = GetCourses(map[string]interface{}{"subject": "ACE", "number": -3}, "LIMIT 10")

	if err == nil {
		t.Errorf("should not be allowed to have non-three-digit number (too small)")
	}

	_, err = GetCourses(map[string]interface{}{"subject": "ACE", "number": 1000}, "LIMIT 10")

	if err == nil {
		t.Errorf("should not be allowed to have non-three-digit number (too big)")
	}
}
