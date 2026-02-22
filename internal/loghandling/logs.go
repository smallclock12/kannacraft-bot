package loghandling

import (
	"log/slog"

	"github.com/smallclock12/kannacraft-bot/db"
	"github.com/smallclock12/kannacraft-bot/mc"
)

func ProcessLogs(logs []db.Log) {

	for _, l := range logs {
		ProcessLog(l)
	}
}

func ProcessLog(log db.Log) {
	slog.Info("Processing log", slog.String("log", log.Id))
	for _, h := range handlers {
		if h.CanProcess(log) {
			slog.Info("Handling log", slog.String("log", log.Id))
			err := h.Process(log)
			if err == nil {
				err := db.ActionLog(log.Id)
				if err != nil {
					slog.Error("Error updating log to actioned", slog.Any("error", err))
				}
			}
			return
		}
	}
}

var handlers = []LogHandler{nmc}

type LogHandler interface {
	CanProcess(log db.Log) bool
	Process(log db.Log) error
}

type NamemcLogHandler int
var nmc NamemcLogHandler = 1

func (h NamemcLogHandler) CanProcess(log db.Log) bool {
	return log.Type == db.Namemc
}

func (h NamemcLogHandler) Process(log db.Log) error {
	slog.Info("Handler processing log", slog.String("log", log.Id))
	return mc.LikedNamemc(log.User)
}


