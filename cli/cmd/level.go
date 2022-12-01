/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/CS222-UIUC/course-project-group-10.git/data"
	"github.com/spf13/cobra"
)

func levelFunc(cmd *cobra.Command, args []string) {
	// read in course argument
	if len(args) < 2 {
		fmt.Println("malformed arguments")
		return
	}

	level, err := strconv.Atoi(args[1])
	if err != nil || level%100 != 0 {
		fmt.Println("Level must be a multiple of 100")
		return
	}

	courses, err := getSubject(args[:1])
	if err != nil {
		return
	}

	// filter
	var n int
	for _, course := range courses {
		if course.Number >= level && course.Number < level+100 {
			courses[n] = course
			n++
		}
	}
	courses = courses[:n]

	fmt.Println(data.CoursesToString(courses))
}

// coursesCmd represents the courses command
var levelCmd = &cobra.Command{
	Use:   "level [int n: n is a multiple of 100]",
	Short: "Lists courses in a department at a certain level",
	Long: `Lists courses in a department which start with a given number, eg 100-level courses.
Usage:
subject level [subject] [level] eg. subject level ART 100`,
	Run: levelFunc,
}

func init() {
	subjectCmd.AddCommand(levelCmd)

}
