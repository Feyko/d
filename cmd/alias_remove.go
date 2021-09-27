package cmd

import (
	"d/pkg/alias"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var aliasRemoveCmd = &cobra.Command{
	Use:   "remove",
	Aliases: []string{"delete", "r", "d"},
	Short: "Remove an alias",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalf("This command requires at least one argument")
		}
		for _, name := range args {
			err := alias.Remove(name)
			if err != nil {
				log.Fatalf("Could not remove the alias %v: %v", name, err)
			}
			fmt.Printf("Alias %v removed\n", name)
		}

	},
}

func init() {
	aliasCmd.AddCommand(aliasRemoveCmd)
}
