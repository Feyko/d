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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fmaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fmaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
