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
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// leaguesCmd represents the leagues command
var leaguesCmd = &cobra.Command{
	Use:   "leagues",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("leagues called")
		url := "http://api.football-data.org/v2/competitions/BL1/standings"

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("x-auth-token", viper.Get("apikey").(string))

		resp, _ := http.DefaultClient.Do(req)

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(resp)
		fmt.Println(string(body))
	},
}

func init() {
	rootCmd.AddCommand(leaguesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// leaguesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// leaguesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
