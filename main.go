package main

import (
	"fmt"

	"log"

	"github.com/url_shortener/api"
	"github.com/url_shortener/database"
	"github.com/url_shortener/pkg/config"
	"github.com/url_shortener/pkg/dbclient"
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

	err = db.AutoMigrate(&database.UrlMapping{})

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	r := api.RegisterRoutes()
	r.Run(fmt.Sprintf("%s%d", ":", config.Config.Port))

}
