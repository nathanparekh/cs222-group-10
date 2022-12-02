package cmd

import (
	"fmt"
	//"strconv"

	"github.com/CS222-UIUC/course-project-group-10.git/data"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func gpaFunc(cmd *cobra.Command, args []string) {
	courses, err := getCourse(args)
	if err != nil {
		fmt.Println("Error getting course:", err)
		fmt.Println("Usage:")
		fmt.Println("course [course name] to get a course by name (eg. course \"Data Structures\")")
		fmt.Println("course [subject] [number] to get a course by number (eg. course CS 225)")
		return
	}

	// read in flags
	latest, _ := cmd.Flags().GetBool("latest")
	oldest, _ := cmd.Flags().GetBool("oldest")
	num, err := cmd.Flags().GetInt("number")
	if err != nil {
		fmt.Println("flag would not be read", err)
	}

	var gpas []data.GPAEntry
	var tmp []data.GPAEntry
	// a course's id is different for every term, thus we need to search for each unique id of a course
	// (AAS 100 Spring 2020 has a different primary key and different gpa entry than Fall 2020)

	// query db and add all gpas from all terms to the slice
	// if two GPA entries have the same instructor id in the same term, SUM the grade
	for i := range courses {
		tmp, err = data.GetGPAEntry(map[string]interface{}{}, "SELECT id, course_id, instructor_id, sched_type, "+
			"SUM(a_plus) a_plus, SUM(a) a, SUM(a_minus) a_minus, "+
			"SUM(b_plus) b_plus, SUM(b) b, SUM(b_minus) b_minus, "+
			"SUM(c_plus) c_plus, SUM(c) c, SUM(c_minus) c_minus, "+
			"SUM(d_plus) d_plus, SUM(d) d, SUM(d_minus) d_minus, "+
			"SUM(f) f, SUM(w) w FROM GPA_Entry WHERE course_id=\""+courses[i].Id+
			"\" GROUP BY instructor_id")
		if err != nil {
			fmt.Println(err)
			return
		}
		gpas = append(gpas, tmp...)
	}

	// if we have less gpa entries than the limit, change the limit
	if num >= len(gpas) {
		num = len(gpas)
	}
	if (latest || oldest) && num != -1 {
		if oldest {
			// if sorting by oldest, reverse the slice
			fmt.Println("Sorting by oldest")
			// reverse slice
			var reversed []data.GPAEntry
			for i := len(gpas)-1; i >= 0; i-- {
				reversed = append(reversed, gpas[i])
				fmt.Println(gpas[i])
			}
			gpas = reversed
		}
		// print courses
		data.PrintGpas(gpas[:num])

	} else {
		// by default, just print latest
		fmt.Println("Latest Offering: ")
		data.PrintGpas(gpas[:1])
	}
}

// coursesCmd represents the courses command
var gpaCmd = &cobra.Command{
	Use:   "gpa",
	Short: "Lists GPAs of past sections of the class",
	Long: `Lists GPAs of past sections of the class, providing information on average GPA and 
			percentage of students receiving a certain letter grade. Can be sorted by instructor.`,
	Run: gpaFunc,
}

func init() {
	courseCmd.AddCommand(gpaCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// var course string
	var latest bool
	var oldest bool
	var num int

	// historyCmd.Flags().StringVarP(&course, "course", "c", "", "Course to find")
	gpaCmd.Flags().BoolVarP(&latest, "latest", "l", false, "Sort by latest first")
	gpaCmd.Flags().BoolVarP(&oldest, "oldest", "o", false, "Sort by oldest first")
	gpaCmd.Flags().IntVarP(&num, "number", "n", 8, "Number of results to list")
}
