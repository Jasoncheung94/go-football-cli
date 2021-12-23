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
	Long:  `Displays the tables for a given competition code.`,
	Run: func(cmd *cobra.Command, args []string) {
		code, err := cmd.Flags().GetString("competition")
		if err != nil {
			code = "PL"
		}
		football.FetchTable(code).DisplayWithPterm()
	},
}

func init() {
	rootCmd.AddCommand(tablesCmd)
	tablesCmd.Flags().String("competition", "PL", "Find tables for competition")
}
