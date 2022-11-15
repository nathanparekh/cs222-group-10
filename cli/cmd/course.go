/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"

	"github.com/spf13/cobra"

	"github.com/CS222-UIUC/course-project-group-10.git/cli/data"
)

func getCourse(args []string) (data.Course, error) {
	if len(args) == 0 {
	} else if len(args) == 1 { // argument is probably a course name
		return data.GetCourseByName(args[0])
	} else if len(args) == 2 { // argument is probably a course subject and number
		number, err := strconv.Atoi(args[1])
		if err != nil {
			return data.Course{}, errors.New("given course number is not a number")
		}
		return data.GetCourseByNum(args[0], number)
	}

	return data.Course{}, errors.New("malformed arguments")
}

func printCourse(cmd *cobra.Command, args []string) {
	course, err := getCourse(args)

	if err != nil {
		fmt.Println("Error getting course:", err)
		fmt.Println("Usage:")
		fmt.Println("course [course name] to get a course by name (eg. course \"Data Structures\")")
		fmt.Println("course [subject] [number] to get a course by number (eg. course CS 225)")
	} else {
		fmt.Println(course)
	}
}

// coursesCmd represents the courses command
var courseCmd = &cobra.Command{
	Use:   "course",
	Short: "Lists a course",
	Long: `When passed a specific course, prints its details.
Usage:
course [course name] to get a course by name (eg. course "Data Structures")
course [subject] [number] to get a course by number (eg. course CS 225)`,
	Run: printCourse,
}

func init() {
	rootCmd.AddCommand(courseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// coursesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// coursesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
