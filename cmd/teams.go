/*
Copyright © 2021 Jason Cheung <jasoncheung94@hotmail.com>

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
	"github.com/jasoncheung94/go-football-cli/pkg/api/football"
	"github.com/spf13/cobra"
)

// teamsCmd represents the teams command
var teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "Displays the teams in a competition.",
	Long:  `Displays the teams for a given competition code.`,
	Run: func(cmd *cobra.Command, args []string) {
		code, err := cmd.Flags().GetString("competition")
		if err != nil {
			code = "PL"
		}
		football.FetchTeams(code).DisplayWithPterm()
	},
}

func init() {
	rootCmd.AddCommand(teamsCmd)
	teamsCmd.Flags().String("competition", "PL", "Find tables for competition")
}
