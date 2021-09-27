/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"d/pkg/PATH"
	"log"

	"github.com/spf13/cobra"
)

// PATHListCmd represents the PATHAdd command
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
