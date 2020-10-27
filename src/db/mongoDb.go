package db

import (
	"github.com/Kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver@latest"
)

func InitDB() {
	// Setup mgm default config
	err := mgm.SetDefaultConfig(nil, "", options.Client().ApplyURI("mongodb://mongo:12345@localhost:27017"))
	if err != nil {
		panic(err)
	}
}

// DB CRUD Operations
// Create

// Read

// Update

// Delete
