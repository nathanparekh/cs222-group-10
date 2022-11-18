package cmd

import (
	"errors"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"

	"github.com/spf13/cobra"

	"github.com/CS222-UIUC/course-project-group-10.git/data"
)

func getCourse(args []string) (data.Course, error) {
	if len(args) == 0 {
	} else if len(args) == 1 { // argument is probably a course name
		courses, err := data.GetCoursesByName(args[0], 1)
		return courses[0], err
	} else if len(args) == 2 { // argument is probably a course subject and number
		number, err := strconv.Atoi(args[1])
		if err != nil {
			return data.Course{}, errors.New("given course number is not a number")
		}
		courses, err := data.GetCoursesByNum(args[0], number, 1)
		return courses[0], err
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
}
