package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var PATHCmd = &cobra.Command{
	Use:   "PATH",
	Short: "Manage the paths in the PATH",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalf("Could not print the help for this command: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(PATHCmd)
}
