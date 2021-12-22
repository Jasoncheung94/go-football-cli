/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/jasoncheung94/go-football-cli/pkg/api/football"
	"github.com/spf13/cobra"
)

// scorersCmd represents the scorers command
var scorersCmd = &cobra.Command{
	Use:   "scorers",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("competitionId")
		if err != nil {
			fmt.Println(err)
		}
		football.FetchScorers(id).DisplayWithPterm()
	},
}

func init() {
	rootCmd.AddCommand(scorersCmd)
	scorersCmd.Flags().String("competitionId", "2021", "Find top scorers for competition id")
}
