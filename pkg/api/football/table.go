package football

import (
	"fmt"

	"github.com/jasoncheung94/go-football-cli/pkg/client"
	"github.com/jasoncheung94/go-football-cli/pkg/entity"
)

func FetchTable(competitionCode string) entity.TableData {
	if competitionCode == "" {
		competitionCode = "PL"
	}
	result := entity.TableData{}
	client.RequestData(
		fmt.Sprintf("http://api.football-data.org/v2/competitions/%s/standings", competitionCode),
		&result,
		nil,
	)

	return result
}

// // leaguesCmd represents the leagues command
// var leaguesCmd = &cobra.Command{
// 	Use:   "leagues",
// 	Short: "A brief description of your command",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("leagues called")
// 		url := "http://api.football-data.org/v2/competitions/BL1/standings"

// 		req, _ := http.NewRequest("GET", url, nil)

// 		req.Header.Add("x-auth-token", viper.Get("apikey").(string))

// 		resp, _ := http.DefaultClient.Do(req)

// 		defer resp.Body.Close()
// 		body, _ := ioutil.ReadAll(resp.Body)

// 		fmt.Println(resp)
// 		fmt.Println(string(body))
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(leaguesCmd)

// 	// Here you will define your flags and configuration settings.

// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	// leaguesCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// leaguesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
