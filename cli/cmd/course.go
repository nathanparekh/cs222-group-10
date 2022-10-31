/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"

	"github.com/spf13/cobra"

	"github.com/CS222-UIUC/course-project-group-10.git/cli/data"
)

func course(cmd *cobra.Command, args []string) {
	var course data.Course
	var err error
	if len(args) == 0 {
		log.Fatal("No arguments passed. Command usage: either a subject and number (ex CS 225) or a course name (ex \"Data Structures\")")
	} else if len(args) == 1 { // argument is probably a course name
		course, err = data.GetCourseByName(args[0])
	} else if len(args) == 2 { // argument is probably a course subject and number
		number, atoi_err := strconv.Atoi(args[1])
		if atoi_err != nil {
			log.Fatal("Not a valid course number.")
		}
		course, err = data.GetCourseByNum(args[0], number)
	}
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Error getting course")
	} else {
		fmt.Println(course)
	}
}

// coursesCmd represents the courses command
var courseCmd = &cobra.Command{
	Use:   "course",
	Short: "Lists a course",
	Long:  `When passed a specific course, prints its details`,
	Run:   course,
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
