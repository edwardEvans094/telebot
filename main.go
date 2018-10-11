package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token: "454550958:AAHd8FQnm-x6uHjcIIKBHOfQhmh6TqRtsBY",
		Poller: &tb.LongPoller{
			Timeout: 10 * time.Second,
		},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hello world")
	})

	b.Start()
}
