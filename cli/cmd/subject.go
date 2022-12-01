package cmd

import (
	"errors"
	"fmt"

	"github.com/CS222-UIUC/course-project-group-10.git/data"
	"github.com/spf13/cobra"
)

func getSubject(args []string) ([]data.Course, error) {
	if len(args) == 1 {
		course, err := data.GetCourses(map[string]interface{}{"subject": args[0]}, "")
		return course, err
	} else if len(args) == 2 {
		limit := "LIMIT " + args[1]
		course, err := data.GetCourses(map[string]interface{}{"subject": args[0]}, limit)
		return course, err
	}

	return []data.Course{}, errors.New("malformed arguments")
}

func printSubject(cmd *cobra.Command, args []string) {
	courses, err := getSubject(args)

	if err != nil {
		fmt.Println("Error getting courses:", err)
	} else {
		if len(courses) == 0 {
			fmt.Println("This subject does not have courses in the dataset.")
		} else {
			fmt.Println(data.CoursesToString(courses))
		}
	}
}

// subjectCmd represents the subject command
var subjectCmd = &cobra.Command{
	Use:   "subject",
	Short: "Lists all courses of a given subject",
	Long: `When passed a subject, lists courses in that subject.
Usage:
subject [subject shorthand] to get all courses in that subject (eg subject ART)
subject [subject shorthand] [int n] to get n courses in that subject (eg subject ART 2)`,

	Run: printSubject,
}

func init() {
	rootCmd.AddCommand(subjectCmd)
}
