package cmd

import (
	"d/pkg/PATH"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var PATHIncludesCmd = &cobra.Command{
	Use:   "includes",
	Aliases: []string{"include", "in", "i"},
	Short: "Check if certain paths are in the PATH",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalf("This command requires at least one argument")
		}
		for _, path := range args {
			b, err := PATH.Includes(path)
			if err != nil {
				log.Fatalf("Could not check if the path %v is included in the PATH: %v", path, err)
			}
			included := "is"
			if !b {
				included = "isn't"
			}
			fmt.Printf("The path %v %v in the PATH\n", path, included)
		}
	},
}

func init() {
	PATHCmd.AddCommand(PATHIncludesCmd)
}
