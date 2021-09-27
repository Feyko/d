package cmd

import (
	"d/pkg/alias"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var aliasExistsCmd = &cobra.Command{
	Use:   "exists",
	Aliases: []string{"ex", "e", "check", "includes", "in", "i"},
	Short: "Check if an alias exists",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalf("This command requires at least one argument")
		}
		for _, name := range args {
			b, err := alias.Exists(name)
			if err != nil {
				log.Fatalf("Could not check if the alias %v exists: %v", name, err)
			}
			exists := "exists"
			if !b {
				exists = "doesn't exist"
			}
			fmt.Printf("The alias %v %v in the ~/.bashrc file\n", name, exists)
		}
	},
}

func init() {
	aliasCmd.AddCommand(aliasExistsCmd)
}
