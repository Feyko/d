package cmd

import (
	"d/pkg/fma"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var fmaRemoveCommand = &cobra.Command{
	Use:                    "remove",
	Aliases:                []string{"delete", "r", "d"},
	Short:                  "Remove an FMA",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("Could not get the 'name' flag: %v", err)
		}

		err = fma.RemoveFMA(name)
		if err != nil {
			log.Fatalf("Could not remove the FMA: %v", err)
		}
		fmt.Println("Success!")
	},
}

func init() {
	fmaCmd.AddCommand(fmaRemoveCommand)

	fmaRemoveCommand.Flags().StringP("name", "n", "", "The name of the action")
	err := fmaRemoveCommand.MarkFlagRequired("name")
	if err != nil {
		log.Fatalf("Could not mark the name flag as required: %v\n", err)
	}
}
