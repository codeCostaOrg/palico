package utils

import (
	"log"
)

func LogCommand(user string, message string) {
	log.Printf("-> %s: %s", user, message)
}

// --- Errors --- //

func LogAPIError(method string, endpoint string, err string) {
	log.Printf("[API ERROR] %s %s: %s", method, endpoint, err)
}

func LogSystemError(loc string, err string) {
	log.Printf("[SYSTEM ERROR] %s: %s", loc, err)
}

func LogDiscordError(loc string, err string) {
	log.Printf("[DISCORD ERROR] %s: %s", loc, err)
}
