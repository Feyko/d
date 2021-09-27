package cmd

import (
	"d/pkg/PATH"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var PATHListCmd = &cobra.Command{
	Use:   "list",
	Aliases: []string{"l"},
	Short: "List all the paths in the PATH",
	Run: func(cmd *cobra.Command, args []string) {
		list, err := PATH.List()
		if err != nil {
			log.Fatalf("Could not remove the paths in the PATH: %v", err)
		}
		for _, e := range list {
			fmt.Println(e)
		}
	},
}

func init() {
	PATHCmd.AddCommand(PATHListCmd)
}
