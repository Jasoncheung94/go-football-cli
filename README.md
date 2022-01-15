## About The Project
go-football-cli is a project implemented in Go that utilizes the Cobra package. It uses an API to retrieve the latest updates of matches, top scorers, league standings and other information. This project requires an API Key to send HTTP requests otherwise it uses a demo mode which retrieves it's information from saved files located in the data folder. 

### Commands
`go run main.go competitions`

<img src="./assets/images/competitions.png" width=300>


`go run main.go matches`

<img src="./assets/images/matches.png" width=300>


`go run main.go scorers`

<img src="./assets/images/scorers.png" width=300>


`go run main.go tables`

<img src="./assets/images/competitions.png" width=300>

`go run main.go teams`

<img src="./assets/images/teams.png" width=300>




#### Links
* [Football API](https://www.football-data.org/)
* [Cobra Package](https://github.com/spf13/cobra)