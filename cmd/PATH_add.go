package cmd

import (
	"d/pkg/PATH"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var PATHAddCmd = &cobra.Command{
	Use:   "add",
	Aliases: []string{"create", "a", "c"},
	Short: "Add paths to the PATH",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalf("This command requires at least one argument")
		}
		for _, path := range args {
			err := PATH.Add(path)
			if err != nil {
				log.Fatalf("Could not add the path %v to the PATH: %v", path, err)
			}
			fmt.Printf("Path %v added to the PATH\n", path)
		}
	},
}

func init() {
	PATHCmd.AddCommand(PATHAddCmd)
}
