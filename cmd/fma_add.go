package cmd

import (
	"d/pkg/fma"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// fmaRemoveCommand represents the add command
var fmaAddCmd = &cobra.Command{
	Use:                    "add",
	Aliases:                []string{"create"},
	Short:                  "Add an FMA",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("Could not get the 'name' flag: %v", err)
		}

		command, err := cmd.Flags().GetString("command")
		if err != nil {
			log.Fatalf("Could not get the 'command' flag: %v", err)
		}

		targetsContext, err := cmd.Flags().GetBool("targetsContext")
		if err != nil {
			log.Fatalf("Could not get the 'targetsContext' flag: %v", err)
		}

		targetsLocation, err := cmd.Flags().GetBool("targetsLocation")
		if err != nil {
			log.Fatalf("Could not get the 'targetsLocation' flag: %v", err)
		}

		err = fma.AddFMA(name, command, targetsContext, targetsLocation)
		if err != nil {
			log.Fatalf("Could not add the FMA: %v", err)
		}
		fmt.Println("Success!")
	},
}

func init() {
	fmaCmd.AddCommand(fmaAddCmd)

	fmaAddCmd.Flags().StringP("name", "n", "",
		"The name of the action")
	err := fmaAddCmd.MarkFlagRequired("name")
	if err != nil {
		log.Fatalf("Could not mark the name flag as required: %v\n", err)
	}

	fmaAddCmd.Flags().StringP("command", "c", "",
		"The command the action will run. List of parameters here -> https://askubuntu.com/a/783313")
	err = fmaAddCmd.MarkFlagRequired("command")
	if err != nil {
		log.Fatalf("Could not mark the command flag as required: %v\n", err)
	}

	fmaAddCmd.Flags().BoolP("targetsContext", "f", false,
		"Can the action act on selected files and folders")

	fmaAddCmd.Flags().BoolP("targetsLocation", "l", false,
		"Can the action act on a location")
}
