package cmd

import (
	"d/pkg/PATH"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var PATHRemoveCmd = &cobra.Command{
	Use:   "remove",
	Aliases: []string{"delete", "r", "d"},
	Short: "Remove paths from the PATH",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalf("This command requires at least one argument")
		}
		for _, path := range args {
			err := PATH.Remove(path)
			if err != nil {
				log.Fatalf("Could not remove the path %v from the PATH: %v", path, err)
			}
			fmt.Printf("Path %v removed from the PATH\n", path)
		}
	},
}

func init() {
	PATHCmd.AddCommand(PATHRemoveCmd)
}
