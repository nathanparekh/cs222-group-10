/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	//"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"github.com/spf13/cobra"

	"github.com/CS222-UIUC/course-project-group-10.git/data"

	"github.com/disiqueira/gotree"
	//"github.com/k0kubun/pp/v3"
)

func getCoursePrereqs(course data.Course, courseToPrereqs *map[data.Course][]data.Course) {
	if _, ok := (*courseToPrereqs)[course]; ok {
		return // if course is already in the map
	}

	if !course.SectionInfo.Valid {
		return // if course does not have section info
	}

	prereqIdx := strings.Index(course.SectionInfo.String, "Prerequisite")

	if prereqIdx == -1 {
		return // if section info does not contain "Prerequisite"
	}

	prereqInfo := course.SectionInfo.String[prereqIdx+14:]
	parts := strings.Split(prereqInfo, " ")

	// naively parse section info to find prerequisites:
	for idx := range parts[:len(parts)-1] {
		subject := parts[idx]     // possible subject
		courseNum := parts[idx+1] // possible course num

		re := regexp.MustCompile(`[^\w\s]`) // remove punctuation
		subject = string(re.ReplaceAll([]byte(subject), []byte("")))
		courseNum = string(re.ReplaceAll([]byte(courseNum), []byte("")))

		// naively validate subject & courseNum (by their lengths)
		if (len(subject) <= 4 && len(subject) >= 2) && len(courseNum) == 3 {
			prereq, err := getCourse([]string{subject, courseNum})
			if err == nil && prereq != (data.Course{}) {
				// subject & courseNum are valid & represent a course
				(*courseToPrereqs)[course] = append((*courseToPrereqs)[course], prereq)
				getCoursePrereqs(prereq, courseToPrereqs)
			}
		}
	}
}

func fillPrereqTree(node *gotree.Tree, course data.Course, courseToPrereqs *map[data.Course][]data.Course) {
	prereqs := (*courseToPrereqs)[course]
	(*courseToPrereqs)[course] = []data.Course{} // empty array to avoid cycles & repetitive subtrees

	for _, prereq := range prereqs {
		prereqNode := (*node).Add(prereq.Subject + " " + strconv.Itoa(prereq.Number))
		fillPrereqTree(&prereqNode, prereq, courseToPrereqs)
	}
}

func printPrereqs(cmd *cobra.Command, args []string) {
	course, err := getCourse(args)
	if err != nil {
		fmt.Println("Error getting course:", err)
		fmt.Println("Usage:")
		fmt.Println("prereqs [course name] to get a course by name (eg. course \"Data Structures\")")
		fmt.Println("prereqs [subject] [number] to get a course by number (eg. course CS 225)")
		return
	}

	courseToPrereqs := map[data.Course][]data.Course{}
	getCoursePrereqs(course, &courseToPrereqs)

	root := gotree.New(course.Subject + " " + strconv.Itoa(course.Number))
	fillPrereqTree(&root, course, &courseToPrereqs)
	fmt.Print(root.Print())
}

var prereqsCmd = &cobra.Command{
	Use:   "prereqs",
	Short: "Display prerequisites of a course",
	Long:  `Search for a course and print a tree of its prerequisites`,
	Run:   printPrereqs,
}

func init() {
	rootCmd.AddCommand(prereqsCmd)
}
