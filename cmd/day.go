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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tonygaetani/advent-of-code-2021/adventofcode"
)

// dayCmd represents the day command
var dayCmd = &cobra.Command{
	Use:   "day",
	Short: "Runs the advent-of-code solution for the given day, if it exists",
	Long:  `Runs the advent-of-code solution for the given day, if it exists`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Day expects exactly 1 argument, an integer between 1 and 25")
			return
		}

		var day = args[0]
		fmt.Printf("Running code for day %s\n", day)
		if day == "1" {
			adventofcode.Day1()
		}
	},
}

func init() {
	rootCmd.AddCommand(dayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
