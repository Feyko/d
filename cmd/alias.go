package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Manage the aliases",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalf("Could not print the help for this command: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(aliasCmd)
}
