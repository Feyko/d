package cmd

import (
	"d/pkg/alias"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var aliasAddCmd = &cobra.Command{
	Use:   "add",
	Aliases: []string{"create", "a", "c"},
	Short: "Add an alias",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("Could not get the 'name' flag: %v", err)
		}

		exec, err := cmd.Flags().GetString("exec")
		if err != nil {
			log.Fatalf("Could not get the 'exec' flag: %v", err)
		}
		err = alias.Add(name, exec)
		if err != nil {
			log.Fatalf("Could not add the alias %v: %v", name, err)
		}
		fmt.Printf("Alias %v added\n", name)
	},
}

func init() {
	aliasCmd.AddCommand(aliasAddCmd)

	aliasAddCmd.Flags().StringP("name", "n", "",
		"The name of the alias")
	err := aliasAddCmd.MarkFlagRequired("name")
	if err != nil {
		log.Fatalf("Could not mark the name flag as required: %v\n", err)
	}

	aliasAddCmd.Flags().StringP("exec", "e", "", "The command the alias represents")
	err = aliasAddCmd.MarkFlagRequired("exec")
	if err != nil {
		log.Fatalf("Could not mark the exec flag as required: %v\n", err)
	}
}
