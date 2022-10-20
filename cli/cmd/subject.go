/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/blockloop/scan"
	"github.com/spf13/cobra"
)

func subject(cmd *cobra.Command, args []string) {
	db, err := sql.Open("sqlite3", "../python/gpa_dataset.db")

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("DB error")
	}

	fmt.Println(args)
	query := `SELECT "Course Title" FROM courses WHERE Subject = @subject;`
	courses, err := db.Query(query, args[0])

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Error getting courses")
	}
	var coursesList []string
	err = scan.Rows(&coursesList, courses)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Error parsing data")
	}
	jsonCourses, err := json.Marshal(coursesList)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonCourses))
	// fmt.Println(courses)
	db.Close()
}

// coursesCmd represents the courses command
var subjectCmd = &cobra.Command{
	Use:   "subject",
	Short: "Lists courses belonging to a particular subject",
	Long:  `When passed a string subject, finds all courses belonging to that subject`,
	Run:   subject,
}

func init() {
	rootCmd.AddCommand(subjectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// coursesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// coursesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
