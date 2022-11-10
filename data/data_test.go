package data

import (
	"fmt"
	"testing"
)

func TestValidName(t *testing.T) {
	_, err := GetCoursesByName("Microcomputer Applications", 1)
	if err != nil {
		t.Errorf("error when getting valid course")
		fmt.Println(err)
	}

	_, err = GetCoursesByName("", 1)

	if err == nil {
		t.Errorf("should not be allowed to get empty course")
	}
}

func TestValidNumber(t *testing.T) {
	_, err := GetCoursesByNum("ACE", 161, 1)
	if err != nil {
		t.Errorf("error when getting valid course")
		fmt.Println(err)
	}

	_, err = GetCoursesByNum("", 225, 1)

	if err == nil {
		t.Errorf("should not be allowed to get empty subject")
	}

	_, err = GetCoursesByNum("CS", -3, 1)

	if err == nil {
		t.Errorf("should not be allowed to have non-three-digit number (too small)")
	}

	_, err = GetCoursesByNum("CS", 1000, 1)

	if err == nil {
		t.Errorf("should not be allowed to have non-three-digit number (too big)")
	}
}
