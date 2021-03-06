package cmd

import (
	"d/pkg/fma"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var fmaAddCmd = &cobra.Command{
	Use:                    "add",
	Aliases:                []string{"create", "a", "c"},
	Short:                  "Add an FMA",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("Could not get the 'name' flag: %v", err)
		}

		command, err := cmd.Flags().GetString("exec")
		if err != nil {
			log.Fatalf("Could not get the 'exec' flag: %v", err)
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

	fmaAddCmd.Flags().StringP("exec", "e", "",
		"The command the action will execute. List of parameters here -> https://askubuntu.com/a/783313")
	err = fmaAddCmd.MarkFlagRequired("exec")
	if err != nil {
		log.Fatalf("Could not mark the exec flag as required: %v\n", err)
	}

	fmaAddCmd.Flags().BoolP("targetsContext", "f", false,
		"Can the action act on selected files and folders")

	fmaAddCmd.Flags().BoolP("targetsLocation", "l", false,
		"Can the action act on a location")
}
