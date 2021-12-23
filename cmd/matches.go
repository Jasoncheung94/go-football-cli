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
	"fmt"

	"github.com/jasoncheung94/go-football-cli/pkg/api/football"
	"github.com/spf13/cobra"
)

// matchesCmd represents the matches command
var matchesCmd = &cobra.Command{
	Use:   "matches",
	Short: "Displays the matches for a list of competition ids",
	Long: `Displays the matches for a list of competition ids.
The Date, Home team and Away team are shown for the upcoming matches within the
next week.
`,
	Run: func(cmd *cobra.Command, args []string) {
		ids, err := cmd.Flags().GetStringSlice("competitionIds")
		if err != nil {
			fmt.Println(err)
		}
		football.FetchMatches(ids).Display()
	},
}

func init() {
	rootCmd.AddCommand(matchesCmd)
	matchesCmd.Flags().StringSlice("competitionIds", []string{"2021"}, "Find tables for competition ids")
}
