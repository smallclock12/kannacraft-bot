package db

import (
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var db *mongo.Database

const database string = "kannacraftdb"
const DOCUMENT_NOT_FOUND = "document not found"

var collections = struct {
	logs string
}{
	logs: "logs",
}

func Connect(connectionString string) {
	client, err := mongo.Connect(options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}

	db = client.Database(database)
	runMigrations()
}

func convertIds[T any](ids []any) []T {
	n := make([]T, len(ids))

	for i, v := range ids {
		n[i] = v.(T)
	}

	return n
}
