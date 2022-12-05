package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/CS222-UIUC/course-project-group-10.git/data"
	"github.com/spf13/cobra"
)

func getSubject(args []string) ([]data.Course, error) {
	if len(args) == 1 {
		course, err := data.GetCourses(map[string]interface{}{"subject": args[0]}, "")
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

			// remove courses that are not in the desired semester
			semester, _ := cmd.Flags().GetString("semester")
			if semester != "all" {
				var n int
				for _, course := range courses {
					courseSemseter := course.Term + " " + strconv.Itoa(course.Year)
					if courseSemseter == semester {
						courses[n] = course
						n++
					}
				}
				courses = courses[:n]
			}

			// remove courses that are not in the specified range
			level, _ := cmd.Flags().GetInt("level")
			if level != 0 {
				if level%100 != 0 {
					fmt.Println("Level must be a multiple of 100")
					return
				}

				var n int
				for _, course := range courses {
					if course.Number >= level && course.Number < level+100 {
						courses[n] = course
						n++
					}
				}
				courses = courses[:n]
			}

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
subject [subject] to get all courses in that subject (eg subject CS)`,

	Run: printSubject,
}

func init() {
	rootCmd.AddCommand(subjectCmd)

	var level int
	var semester string

	subjectCmd.Flags().IntVarP(&level, "level", "l", 0, "Course level")
	subjectCmd.Flags().StringVarP(&semester, "semester", "s", "Spring 2023", "Semester (type all for all semesters in dataset)")
}
