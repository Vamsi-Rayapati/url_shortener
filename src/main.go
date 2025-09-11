package main

import (
	"fmt"

	"log"

	"github.com/smartbot/account/api"
	"github.com/smartbot/account/database"
	"github.com/smartbot/account/pkg/config"
	"github.com/smartbot/account/pkg/dbclient"
)

func main() {
	var err error
	config.LoadConfig()

	// config.ConnectDB()
	db, err := dbclient.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	err = db.AutoMigrate(&database.User{})

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	r := api.RegisterRoutes()
	r.Run(fmt.Sprintf("%s%d", ":", config.Config.Port))

}
