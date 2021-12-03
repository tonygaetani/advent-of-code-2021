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

func init() {
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
			switch day {
			case "1":
				adventofcode.Day1()
				break
			case "2":
				adventofcode.Day2()
				break
			default:
				fmt.Printf("No implementation for day %s\n", day)
			}
		},
	}

	rootCmd.AddCommand(dayCmd)
}
