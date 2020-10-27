package models

import (
	// _ "go.mongodb.org/mongo-driver@latest"
	"github.com/Kamva/mgm/v3"
)

type Schedule struct {
	mgm.DefaultModel `bson:",inline"`
	Game             *Game
	UpcomingGame     *UpcomingGame
}

type UpcomingGame struct {
	HomeTeam string `json:"home_team" bson:"home_team"`
	AwayTeam string `json:"away_team" bson:"away_team"`
}

type Game struct {
	HomeTeam      string `json:"home_team" bson:"home_team"`
	AwayTeam      string `json:"away_team" bson:"away_team"`
	HomeTeamScore int    `json:"homeTeam_Score" bson:"homeTeam_Score"`
	AwayTeamScore int    `json:"awayteam_Score" bson:"awayTeam_Score"`
}
