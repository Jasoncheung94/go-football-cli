/*
Copyright © 2021 Jason Cheung <jasoncheung94@hotmail.com>

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
	Short: "Displays the top scorers for a league",
	Long: `Using the competition id, this displays the top scorers for that
competition. 
`,
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
