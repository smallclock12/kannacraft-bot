package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type LogType string

const (
	Namemc LogType = "namemc"
	Discord LogType = "discord"
)

type Log struct {
	Id string `bson:"_id"`
	Created int64 `bson:"created"`
	Type LogType `bson:"type"`
	User string `bson:"user"`
	Actioned bool `bson:"actioned"`
	Data bson.D `bson:"data"`
}

func NewLog(id string, logType LogType, data bson.D) Log {
	return Log{
		Id: fmt.Sprintf("%s/%s", id, logType),
		Created: time.Now().Unix(),
		Type: logType,
		User: id,
		Actioned: false,
		Data: data,
	}
}


func InsertManyLogs(logs []Log) (inserted []string, err error) {
	res, err := db.Collection(collections.logs).InsertMany(context.TODO(), logs)
	return convertIds[string](res.InsertedIDs), err
}

func ActionLog(id string) error {
	u := bson.D{
		bson.E{Key: "$set", Value: bson.D{{Key:"actioned", Value:true}}},
	}

	_, err := db.Collection(collections.logs).UpdateByID(context.TODO(), id, u)
	return err
}

