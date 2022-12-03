package cmd

import (
	"errors"
	"fmt"
	"github.com/CS222-UIUC/course-project-group-10.git/data"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func getSections(args []string, semester string, crn int, useCRN bool) ([]data.Section, error) {
	tokens := strings.Split(semester, " ")
	year, err := strconv.Atoi(tokens[1])
	if (err != nil ||
		len(tokens) != 2) ||
		(tokens[0] != "Spring" && tokens[0] != "Fall" && tokens[0] != "Winter" && tokens[0] != "Summer") {
		return nil, errors.New("invalid semester provided")
	}
	if len(args) == 0 {
		if useCRN {
			sections, err := data.GetSections(map[string]interface{}{"crn": crn, "year": year, "term": tokens[0]}, "")
			return sections, err
		} else {
			return nil, errors.New("provide at least one argument or CRN flag")
		}
	} else if len(args) == 1 {
		sections, err := data.GetSections(map[string]interface{}{"name": args[0], "year": year, "term": tokens[0]}, "")
		return sections, err
	} else if len(args) == 2 {
		sections, err := data.GetSections(map[string]interface{}{"subject": args[0], "number": args[1], "year": year, "term": tokens[0]}, "")
		return sections, err
	}

	return []data.Section{}, errors.New("malformed arguments")
}

func printSections(cmd *cobra.Command, args []string) {
	crnLookup := cmd.Flags().Changed("crn")
	crn, _ := cmd.Flags().GetInt("crn")

	semester, _ := cmd.Flags().GetString("semester")
	sections, err := getSections(
		args,
		semester,
		crn,
		crnLookup,
	)
	if err != nil {
		fmt.Println("error while attempting to get section data")
		return
	}

	if len(sections) == 0 {
		fmt.Println("No sections found")
		return
	}

	course, err := data.GetCourseByDatabaseId(sections[0].CourseId)

	if err != nil {
		fmt.Println("error while attempting to get course info")
		return
	} else {
		fmt.Println(data.CoursesToString(course))
		fmt.Println(data.SectionsToString(sections))
	}
}

// coursesCmd represents the courses command
var sectionCmd = &cobra.Command{
	Use:   "sections",
	Short: "Lists all sections of a specified course for a specified term, OR for one specified CRN",
	Long: `When passed a specific course, prints its sections.
	Usage:
		sections [course name] to get a sections for a course by name (eg. course "Data Structures")
		sections [subject] [number] to get a course by number (eg. course CS 225)
		sections --crn [crn] to pull data for one specific section
`,
	Run: printSections,
}

func init() {
	rootCmd.AddCommand(sectionCmd)

	var semester string
	var sectionCrn int

	sectionCmd.Flags().StringVar(&semester, "semester", "Spring 2023", "specify a semester (default Spring 2023)")
	sectionCmd.Flags().IntVar(&sectionCrn, "crn", -1, "Specify a CRN")
}
