package data

import (
	"testing"
)

func TestValidName(t *testing.T) {
	_, err := GetCourseByName("Microcomputer Applications")
	if err != nil {
		t.Errorf("error when getting valid course")
	}

	_, err = GetCourseByName("")

	if err == nil {
		t.Errorf("should not be allowed to get empty course")
	}
}

func TestValidNumber(t *testing.T) {
	_, err := GetCourseByNum("ACE", 161)
	if err != nil {
		t.Errorf("error when getting valid course")
	}

	_, err = GetCourseByNum("", 225)

	if err == nil {
		t.Errorf("should not be allowed to get empty subject")
	}

	_, err = GetCourseByNum("CS", -3)

	if err == nil {
		t.Errorf("should not be allowed to have non-three-digit number (too small)")
	}

	_, err = GetCourseByNum("CS", 1000)

	if err == nil {
		t.Errorf("should not be allowed to have non-three-digit number (too big)")
	}
}
