package main

import (
	"flag"
	"go_telegram_bot_later_read_links/clients/telegram"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	gtClient := telegram.New(tgBotHost, getToken())

	//fetcher = fetcher.New(gtClient)

	//processor = processor.New(gtClient)

	//consumer.Start(fetcher, processor)
}

func getToken() string {
	token := flag.String("telegram_api_token", "", "telegram_api_token")

	if *token == "" {
		log.Fatal("Token not found")
	}

	return *token
}
