package main

import (
    "log"
    "os"
	"fmt"

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

    b, err := tb.NewBot(pref)
    if err != nil {
        log.Fatal(err)
    }
	

	b.Handle("/roll", func(m *tb.Message) {
		kek := tb.Recipient(m.Chat)
		m.Dice = tb.Cube
		m.Dice.Send(b, kek, &tb.SendOptions{
			ReplyTo:m,
			})
	})
	b.Handle("/darts", func(m *tb.Message) {
		kek := tb.Recipient(m.Chat)
		m.Dice = tb.Dart
		m.Dice.Send(b, kek, &tb.SendOptions{
			ReplyTo:m,
			})
	})
	b.Handle("/fag", func(m *tb.Message) {
	b.Send(m.Send, tb.SendOptions{
		ReplyTo"Hello fag!"}
		)
	})
	
	b.Start()

}