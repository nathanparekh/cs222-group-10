package cmd

import (
	"errors"
	"fmt"

	"github.com/CS222-UIUC/course-project-group-10.git/data"
	"github.com/spf13/cobra"
)

func getCourse(args []string) ([]data.Course, error) {
	if len(args) == 0 {
	} else if len(args) == 1 { // argument is probably a course name
		course, err := data.GetCourses(map[string]interface{}{"name": args[0]}, "")
		if len(course) > 0 {
			return course[0], err
		} else {
			return data.Course{}, err
		}
	} else if len(args) == 2 { // argument is probably a course subject and number
		course, err := data.GetCourses(map[string]interface{}{"subject": args[0], "number": args[1]}, "")
		if len(course) > 0 {
			return course[0], err
		} else {
			return data.Course{}, err
		}
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
		fmt.Println(data.CoursesToString([]data.Course{course}))
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
