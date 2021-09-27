package cmd

import (
	"d/pkg/PATH"
	"log"

	"github.com/spf13/cobra"
)

var PATHAddCmd = &cobra.Command{
	Use:   "add",
	Aliases: []string{"create", "a", "c"},
	Short: "Add a path to the PATH",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatalf("Could not get the 'path' flag: %v", err)
		}
		err = PATH.Add(path)
		if err != nil {
			log.Fatalf("Could not add the path %v to the PATH: %v", path, err)
		}
		log.Printf("Success! Path %v added to the PATH", path)
	},
}

func init() {
	PATHCmd.AddCommand(PATHAddCmd)

	PATHAddCmd.Flags().StringP("path", "p", "", "The path to add")
	err := PATHAddCmd.MarkFlagRequired("path")
	if err != nil {
		log.Fatalf("Could not make the 'path' flag required in the PATH_add command: %v", err)
	}
}
