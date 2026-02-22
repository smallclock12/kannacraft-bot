package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var migrations map[string]func(*mongo.Database) error = map[string]func(*mongo.Database) error{
	"scaffold": func(db *mongo.Database) error {
		return nil
	},
}

func runMigrations() {
	for k, v := range migrations {

		_, err := db.
			Collection("migrations").
			InsertOne(context.TODO(), bson.D{{Key: "_id", Value: k}})

		if err == nil {
			log.Printf("Running migration %s", k)
			e := v(db)
			if e != nil {
				log.Print(e)
				panic(e)
			}
			log.Printf("Migration %s applied", k)
		}
	}
}
