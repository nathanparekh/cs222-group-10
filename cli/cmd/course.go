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

func course(cmd *cobra.Command, args []string) {
	db, err := sql.Open("sqlite3", "../python/gpa_dataset.db")

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("DB error")
	}
	query := `SELECT Subject FROM courses WHERE "Course Title" = @title;`
	courses, err := db.Query(query, args[0])

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Error getting courses")
	}
	var course string
	err = scan.Row(&course, courses)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Error parsing data")
	}
	jsonCourse, err := json.Marshal(course)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonCourse))
	// fmt.Println(courses)
	db.Close()
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
