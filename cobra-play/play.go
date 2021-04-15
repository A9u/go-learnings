package main

import (
	"github.com/spf13/cobra"
)

func main() {
	cmd := newCommand()
	cmd.AddCommand(newNestedCommand())

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd)

	if err := rootCmd.Execute(); err != nil {
		println(err.Error())
	}
}

func newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			println(`Foo`)
		},
		Use:   `foo`,
		Short: "Command foo",
		Long:  "This is a command",
	}

	return cmd
}

func newNestedCommand() *cobra.Command {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			println(`Bar`)
		},
		Use:   `bar`,
		Short: "Command bar",
		Long:  "This is a nested command",
	}

	return cmd
}
