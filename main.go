package main

import (
    "log"
    "os"
	"strconv"

    tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
    var (
        port      = os.Getenv("PORT")
        publicURL = os.Getenv("PUBLIC_URL") // you must add it to your config vars
        token     = os.Getenv("TOKEN")      // you must add it to your config vars
    )

    webhook := &tb.Webhook{
        Listen:   ":" + port,
        Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
    }

    pref := tb.Settings{
        Token:  token,
        Poller: webhook,
    }
type Recipient interface {
    // Must return legit Telegram chat_id or username
    Recipient() string
}

    b, err := tb.NewBot(pref)
    if err != nil {
        log.Fatal(err)
    }
	
	b.Handle("/roll", func (m *tb.message) {
	m.Dice.Send(b, strconv.FormatInt(m.Chat.ID, 10), &tb.SendOptions{})
	})
	
	b.Start()
}