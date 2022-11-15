/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/CS222-UIUC/course-project-group-10.git/data"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func historyFunc(cmd *cobra.Command, args []string) {
	// read in course argument
	var course []data.Course
	var err error
	if len(args) == 0 {
		log.Fatal("No arguments passed. Command usage: either a subject and number (ex CS 225) or a course name (ex \"Data Structures\")")
	} else if len(args) == 1 { // argument is probably a course name
		course, err = data.GetCoursesByName(args[0], 1)
	} else if len(args) == 2 { // argument is probably a course subject and number
		number, atoi_err := strconv.Atoi(args[1])
		if atoi_err != nil {
			log.Fatal("Not a valid course number.")
		}
		course, err = data.GetCoursesByNum(args[0], number, 1)
	}
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Error getting course")
	}

	// read in flags
	latest, err := cmd.Flags().GetBool("latest")
	if err != nil {
		fmt.Println("flag would not be read", err)
	}
	oldest, err := cmd.Flags().GetBool("oldest")
	if err != nil {
		fmt.Println("flag would not be read", err)
	}
	num, err := cmd.Flags().GetInt("number")
	if err != nil {
		fmt.Println("flag would not be read", err)
	}

	// print depending on flags
	if latest || oldest || num != -1 {
		// get slice of all courses with specified number
		var courses []data.Course
		var err error
		if oldest {
			// if sorting by oldest, reverse the slice
			fmt.Println("Sorting by oldest")
			// reverse courses slice
			courses, err = data.GetCoursesByNum(course[0].Subject, course[0].Number, 100)
			if err != nil {
				fmt.Println(err.Error())
			}
			var reversed []data.Course
			for i := len(courses) - 1; i >= len(courses)-num; i-- {
				reversed = append(reversed, courses[i])
			}
			courses = reversed
		} else if latest || num != -1 {
			courses, err = data.GetCoursesByNum(course[0].Subject, course[0].Number, num)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		// if no --number flag was passed, set it to the length of the courses slice
		// print the sorted results
		for i := 0; i < len(courses); i++ {
			output := data.CourseToString(courses[i])
			fmt.Println(output)
			for i := 0; i < len(output); i++ {
				fmt.Print("-")
			}
			fmt.Println()
		}

	} else {
		// by default, just print latest
		fmt.Println("Latest Offering: ")
		fmt.Println(data.CourseToString(course[0]))
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
	//var course string
	var latest bool
	var oldest bool
	var num int

	//historyCmd.Flags().StringVarP(&course, "course", "c", "", "Course to find")
	historyCmd.Flags().BoolVarP(&latest, "latest", "l", false, "Sort by latest first")
	historyCmd.Flags().BoolVarP(&oldest, "oldest", "o", false, "Sort by oldest first")
	historyCmd.Flags().IntVarP(&num, "number", "n", 8, "Number of results to list")
}
