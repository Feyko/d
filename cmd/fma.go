package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// fmaCmd represents the fma command
var fmaCmd = &cobra.Command{
	Use:   "fma",
	Short: "Manage FileManager Actions",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatalf("Could not display the help message: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(fmaCmd)
}
