/*
Copyright © 2021 Jason Cheung jasoncheung94@hotmail.com

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

// competitionsCmd represents the competitions command
var competitionsCmd = &cobra.Command{
	Use:   "competitions",
	Short: "Fetches all the available competitions",
	Long: `This command fetches all available competitions, display's their ID, name, code and location. 

This information is useful for other commands when the competition ID/Code is required.`,
	Run: func(cmd *cobra.Command, args []string) {
		football.FetchCompetitions().Display()
	},
}

func init() {
	rootCmd.AddCommand(competitionsCmd)
}
