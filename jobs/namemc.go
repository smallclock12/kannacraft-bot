package jobs

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"slices"

	"github.com/smallclock12/kannacraft-bot/db"
	"github.com/smallclock12/kannacraft-bot/internal/loghandling"
	"github.com/smallclock12/kannacraft-bot/internal"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const namemcUrl string = "https://api.namemc.com/server/%s/likes"

func CheckNamemc() {
	url := fmt.Sprintf(namemcUrl, internal.Env.ServerName)
	client := http.DefaultClient

	res, err := client.Get(url)
	if err != nil {
		slog.Error("Error making request to namemc", slog.Any("error", err))
		return
	}

	v := []string{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&v)
	if err != nil {
		slog.Error("Error parsing response from namemc", slog.Any("error", err))
		return
	}

	slog.Info("Found users", slog.Any("users", v))

	logs := []db.Log{}
	for _, uuid := range v {
		log := db.NewLog(uuid, db.Namemc, bson.D{})
		logs = append(logs, log)
	}

	ids, err := db.InsertManyLogs(logs)
	if err != nil {
		slog.Error("Error inserting some or all records into mongo", slog.Any("error", err))
	}
	slog.Info("Inserted logs", slog.Any("logs", ids))

	processLogs := []db.Log{}
	for _, l := range logs {
		if slices.Contains(ids, l.Id) {
			processLogs = append(processLogs, l)
		}
	}

	if len(processLogs) == 0 {
		return
	}

	slog.Info("Processing logs", slog.Any("logs", processLogs))
	go loghandling.ProcessLogs(processLogs)
}
