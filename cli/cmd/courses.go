/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func coursesFunc(cmd *cobra.Command, args []string) {
	fmt.Println("Entire list of courses")
}

// coursesCmd represents the courses command
var coursesCmd = &cobra.Command{
	Use:   "courses",
	Short: "Lists all courses",
	Long:  `Entire listing of courses for the University of Illinois at Urbana-Champaign`,
	Run:   coursesFunc,
}

func init() {
	rootCmd.AddCommand(coursesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// coursesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// coursesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
