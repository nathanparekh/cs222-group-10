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

func historyFunc(cmd *cobra.Command, args []string) {
	// read in course argument
	course, err := getCourse(args)
	if err != nil {
		fmt.Println("Error getting course:", err)
		fmt.Println("Usage:")
		fmt.Println("course [course name] to get a course by name (eg. course \"Data Structures\")")
		fmt.Println("course [subject] [number] to get a course by number (eg. course CS 225)")
		return
	}

	// read in flags
	latest, _ := cmd.Flags().GetBool("latest")
	oldest, _ := cmd.Flags().GetBool("oldest")
	num, err := cmd.Flags().GetInt("number")
	if err != nil {
		fmt.Println("flag would not be read", err)
	}

	// initalize argsMap used by the getters
	argsMap := map[string]interface{}{"subject": course[0].Subject, "number": course[0].Number}

	// print depending on flags
	if (latest || oldest) && num != -1 {
		// get slice of all courses with specified number
		var courses []data.Course
		var err error
		if oldest {
			// if sorting by oldest, reverse the slice
			fmt.Println("Sorting by oldest")
			courses, err = data.GetCourses(argsMap, "ORDER BY year DESC"+" LIMIT "+strconv.Itoa(num))
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("could not get course")
			}
		} else if latest || num != -1 {
			// just fetch the correct number of courses
			courses, err = data.GetCourses(argsMap, "LIMIT "+strconv.Itoa(num))
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("could not get course")
			}
		}
		// print courses
		fmt.Println(data.CoursesToString(courses))

	} else {
		// by default, just print latest
		fmt.Println("Latest Offering: ")
		courses, err := data.GetCourses(argsMap, "LIMIT 1")
		if err != nil {
			fmt.Println(err)
			fmt.Println("could not get course")
		}
		fmt.Println(data.CoursesToString(courses))
	}
}

// coursesCmd represents the courses command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Lists when a course was previously offered",
	Long:  `Lists when a course was previously offered. By default it lists starting at latest offering. Only goes back to 2010`,
	Run:   historyFunc,
}

func init() {
	courseCmd.AddCommand(historyCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// var course string
	var latest bool
	var oldest bool
	var num int

	// historyCmd.Flags().StringVarP(&course, "course", "c", "", "Course to find")
	historyCmd.Flags().BoolVarP(&latest, "latest", "l", false, "Sort by latest first")
	historyCmd.Flags().BoolVarP(&oldest, "oldest", "o", false, "Sort by oldest first")
	historyCmd.Flags().IntVarP(&num, "number", "n", 8, "Number of results to list")
}
