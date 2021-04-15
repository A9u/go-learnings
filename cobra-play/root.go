package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "root",
		Short: "Root Command",
		Long:  "This gets executed immediately with main",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Inside rootCmd PersistentPreRun")
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Inside rootCmd PreRun")
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Inside rootCmd Run")
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Inside rootCmd PostRun")
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Inside rootCmd PersistentPostRun")
		},
	}

	var childCmd = &cobra.Command{
		Use:   "child",
		Short: "Child childcommand",
		Long:  "Need to pass child to make it run",
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Inside childCmd PreRun")
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Inside childCmd Run")
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Inside childCmd PostRun")
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Inside childCmd PersistentPostRun")
		},
	}

	rootCmd.AddCommand(childCmd)

	//rootCmd.SetArgs([]string{""})
	rootCmd.Execute()
	fmt.Println()

	// to execute child programatically
	// rootCmd.SetArgs([]string{"child"})
	// rootCmd.Execute()
}
