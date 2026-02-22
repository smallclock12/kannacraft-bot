package main

import (
	"os"
	"os/signal"

	"github.com/smallclock12/kannacraft-bot/db"
	"github.com/smallclock12/kannacraft-bot/internal"
	"github.com/smallclock12/kannacraft-bot/jobs"
)

func main() {
	sigch := make(chan os.Signal, 1)

	db.Connect(internal.Env.ConnectionString)
	jobs.CheckNamemc()
	
	signal.Notify(sigch, os.Interrupt)
	<-sigch
}
