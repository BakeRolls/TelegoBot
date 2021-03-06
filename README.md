# TelegoBot

Telegram bot using long polling written in Go.

The bots not complete yet, since the most `send-` and `update`methods aren't implemented (`SendMessage` and `AnswerInlineQuery` are).

# Example

```go
package main

import (
	"flag"
	"fmt"

	"github.com/bakerolls/telegobot"
)

var (
	bot   *telegobot.TelegoBot
	token = flag.String("token", "", "Token")
)

func main() {
	flag.Parse()

	bot = telegobot.NewBot(*token)

	// Don't register channels if you're not using them. It would block the process.
	bot.Messages = make(chan telegobot.Message)
	bot.Queries = make(chan telegobot.InlineQuery)

	go bot.GetUpdates()

	go messages()
	go queries()

	select {}
}

func messages() {
	for message := range bot.Messages {
		fmt.Println("Message: " + message.Text)

		bot.SendMessage(message.Chat.ID, "Echo " + message.Text)
	}
}

func queries() {
	for query := range bot.Queries {
		fmt.Println("Query: " + query.Query)

		posts, _ := tumblr.Query(query.Query)
		photos := []telegobot.InlineQueryResultPhoto{}

		for i, post := range posts {
			photos = append(photos, telegobot.InlineQueryResultPhoto{
				Type:   "photo",
				ID:     strconv.Itoa(i),
				Photo:  post.High,
				Thumb:  post.Low,
				Width:  post.Width,
				Height: post.Height,
			})
		}

		bot.AnswerInlineQuery(query.ID, photos)
	}
}
```

```shell
$ go run example.go --token 123456789:ABC...
Message: Hello World!
Query: Ha
Query: Har
Query: Harvey
Message: @yourbot Harvey
```
