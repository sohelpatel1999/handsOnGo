package database

import (
	"context"
	"handsOnGO/config"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
)

var Db *firestore.Client

func CreateFirestoreClient() {
	ctx := context.Background()
	file := config.Config.String("firestoreConfig")
	sa := option.WithCredentialsFile(file)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatal().Err(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal().Err(err)
	}

	log.Info().Msg("Connected to Firestore")

	Db = client

}
