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
	"github.com/jasoncheung94/go-football-cli/pkg/table"
	"github.com/spf13/cobra"
)

var competitionCodes = map[string]string{
	"fifa world cup":                "WC",
	"champions league":              "CL",
	"bundesliga":                    "BL1",
	"eredivisie":                    "DED",
	"campeonato brasileiro serie a": "BSA",
	"primera division":              "PD",
	"ligue 1":                       "FL1",
	"championship":                  "ELC",
	"primeira liga":                 "PPL",
	"european championsip":          "EC",
	"serie a":                       "SA",
	"premier league":                "PL",
	"copa liertadores":              "CLI",
}

// tablesCmd represents the tables command
var tablesCmd = &cobra.Command{
	Use:   "tables",
	Short: "Returns the current tables of leagues.",
	Long: `DITTO TODO:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		code, err := cmd.Flags().GetString("league")
		if err != nil {
			code = "PL"
		}
		table.Fetch(code).DisplayWithPterm()
	},
}

func init() {
	rootCmd.AddCommand(tablesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tablesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	tablesCmd.Flags().String("league", "PL", "Find tables for league")
}
