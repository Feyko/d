package cmd

import (
	"d/pkg/PATH"
	"log"

	"github.com/spf13/cobra"
)

var PATHRemoveCmd = &cobra.Command{
	Use:   "remove",
	Aliases: []string{"delete", "r", "d"},
	Short: "Remove a path to the PATH",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatalf("Could not get the 'path' flag: %v", err)
		}
		err = PATH.Remove(path)
		if err != nil {
			log.Fatalf("Could not remove the path %v from the PATH: %v", path, err)
		}
		log.Printf("Success! Path %v removed from the PATH", path)
	},
}

func init() {
	PATHCmd.AddCommand(PATHRemoveCmd)

	PATHRemoveCmd.Flags().StringP("path", "p", "", "The path to remove")
	err := PATHRemoveCmd.MarkFlagRequired("path")
	if err != nil {
		log.Fatalf("Could not make the 'path' flag required in the PATH_remove command: %v", err)
	}
}
