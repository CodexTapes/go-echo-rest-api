package models

import (
	"github.com/Kamva/mgm/v3"
	_ "go.mongodb.org/mongo-driver@latest"
)

type Standings struct {
	mgm.DefaultModel `bson:",inline"`
	Team             *Team
	Rank             int `json:"rank" bson:"rank"`
	Record           *Record
}
