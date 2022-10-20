/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func helloWorld() {
	fmt.Println("hello, world");
}

// helloWorldCmd represents the helloWorld command
var helloWorldCmd = &cobra.Command{
	Use:   "helloWorld",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		helloWorld();
	},
}

func init() {
	rootCmd.AddCommand(helloWorldCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloWorldCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloWorldCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
