package cmd

import (
	"d/pkg/bashrc"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var bashrcAddCmd = &cobra.Command{
	Use:   "add",
	Aliases: []string{"create", "a", "c"},
	Short: "Add a line or multiple lines to the bashrc file of the current user",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalf("This command requires at least one argument")
		}
		for _, s := range args {
			err := bashrc.Append("\n" + s)
			if err != nil {
				log.Fatalf("Could not add the string to the bashrc file: %v", s, err)
			}
			fmt.Printf("Successfuly appended the following text to the bashrc file:'%v'\n", s)
		}
	},
}

func init() {
	bashrcCmd.AddCommand(bashrcAddCmd)
}
