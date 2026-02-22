package mc

import (
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/smallclock12/kannacraft-bot/internal"
	"github.com/smallclock12/kannacraft-bot/mc/client"
)

type RconActions interface {
	LikedNamemc(uuid string) error
	Close()
}

func LikedNamemc(uuid string) error {
	return DefaultHandler.LikedNamemc(uuid)
}

func Close() {
	DefaultHandler.Close()
}

var DefaultHandler = newDefaultHandler()

func newDefaultHandler() RconActions {

	// Create a new client and connect to the server.
	client, err := client.NewClient(client.ClientOptions{
		Hostport: internal.Env.RconLocation,
		Timeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Send some commands.
	if err := client.Authenticate(internal.Env.RconPassword); err != nil {
		log.Fatal(err)
	}

	return &rconHandler{client: client}
}

type rconHandler struct {
	client *client.Client
}

func (h *rconHandler) Close() {
	h.client.Close()
}

func (h *rconHandler) LikedNamemc(uuid string) error {
	slog.Info("Processing liked namemc request", slog.String("uuid", uuid))
	resp,err := h.client.SendCommand(fmt.Sprintf("execute as %s run advancement grant @s only %s", uuid, internal.Env.NamemcAdvancement))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Body)
	slog.Info("Completed liked namemc request", slog.String("uuid", uuid))
	return nil
}
