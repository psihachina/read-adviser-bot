package main

import (
	"flag"
	"log"

	"github.com/psihachina/telegrambot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	tgClient := telegram.New(tgBotHost, mustToken())

	// fetcher := fetcher.New(token)

	// processor := processor.New(token)

	// consumer.Start(fetcher, processor)

}

func mustToken() string {
	token := flag.String(
		"token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
