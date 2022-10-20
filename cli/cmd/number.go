/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"regexp"
)

func numberFunc(cmd *cobra.Command, args []string) {
	regex := regexp.MustCompile(`^(\D{0,4})(\d{0,3})$`)

	matches := regex.FindStringSubmatch(args[0])

	//fmt.Println(matches)

	courseSubject := matches[1]
	courseNumber := matches[2]

	//fmt.Println(courseSubject)
	//fmt.Println(courseNumber)

	db, err := sql.Open("sqlite3", "../python/gpa_dataset.db")

	if err != nil {
		fmt.Println(err)
	}

	query := `SELECT "Course Title" FROM courses WHERE ("Subject"=@courseSubject AND "Number"=@courseNumber);`

	var title string
	db.QueryRow(query, courseSubject, courseNumber).Scan(&title)

	fmt.Println(title)
}

// numberCmd represents the number command
var numberCmd = &cobra.Command{
	Use:   "number",
	Short: "get course information by course number",
	Long: `Gets course information by course number. For example: "CS225", "ECON415", "MUSC487" etc.

Will obtain information about historical GPA, title, etc.`,
	Run: numberFunc,
}

func init() {
	rootCmd.AddCommand(numberCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// numberCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// numberCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
