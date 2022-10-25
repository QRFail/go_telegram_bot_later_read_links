package telegram

import "go_telegram_bot_later_read_links/clients/telegram"

type Processor struct {
	client *telegram.Client
	offset int
}

func New(client *telegram.Client) {

}
