package models

import (
	// _ "go.mongodb.org/mongo-driver@latest"
	"github.com/Kamva/mgm/v3"
)

type Team struct {
	// DefaultModel add _id,created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	TeamName         string `json:"team_name" bson:"team_name"`
	TeamID           int    `json:"team_id" bson:"team_id"`
	League           string `json:"league" bson:"league"`
	Sport            string `json:"sport" bson:"sport"`
}

type Record struct {
	Wins   int `json:"wins" bson:"wins"`
	Losses int `json:"losses" bson:"losses"`
}

type Game struct {
	HomeTeam      string `json:"home_team" bson:"home_team"`
	AwayTeam      string `json:"away_team" bson:"away_team"`
	HomeTeamScore int    `json:"homeTeam_Score" bson:"homeTeam_Score"`
	AwayTeamScore int    `json:"awayteam_Score" bson:"awayTeam_Score"`
}

type Score struct {
	// DefaultModel add _id,created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Team             *Team
	Game             *Game
	Record           *Record
}

type ScoresCollection struct {
	Scores []Score `json:"scores"`
}
