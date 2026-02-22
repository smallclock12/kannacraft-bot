package internal

import (
	"fmt"
	"os"
)

const envPrefix = "KANNACRAFT"

func envKey(s string) string {
	return fmt.Sprintf("%s_%s", envPrefix, s)
}

func getEnv(s string) string {
	return os.Getenv(envKey(s))
}

var Env = struct {
	ConnectionString string
	DiscordToken     string
	DiscordUser      string
	ServerName string
	RconPassword string
	RconLocation string
	NamemcAdvancement string 
}{
	ConnectionString: getEnv("CS"),
	DiscordToken:     getEnv("TOKEN"),
	DiscordUser:      getEnv("USER"),
	ServerName: getEnv("SERVER_NAME"),
	RconPassword: os.Getenv("RCON_PASSWORD"),
	RconLocation: os.Getenv("RCON_LOCATION"),
	NamemcAdvancement: getEnv("NAMEMC_ADVANCEMENT"),
}

