package main

import (
	"handsOnGO/config"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"handsOnGO/database"
	"handsOnGO/routes"
)

func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Document Service Running")

	config.ConfigInit()
	port := config.Config.String("port")
	//database setting
	database.CreateFirestoreClient()
	//routes
	router := routes.DocumentRoutes()

	if err := router.Run(":" + port); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}

}
