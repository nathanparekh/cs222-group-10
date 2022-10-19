/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func departmentFunc(cmd *cobra.Command, args []string) {
	fmt.Println("Course listing for " + args[0])
}

// departmentCmd represents the department command
var departmentCmd = &cobra.Command{
	Use:   "department",
	Short: "gets course listing for department",
	Long:  `Gets the list of courses available for the provided department (ex. CS or CWL or ECON)`,
	Run: func(cmd *cobra.Command, args []string) {
		departmentFunc(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(departmentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// departmentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// departmentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
