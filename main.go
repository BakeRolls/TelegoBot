package main

import (
	"flag"
	"strings"

	"strconv"

	"git.192k.pw/bake/telegobot/botcommands"
	"git.192k.pw/bake/telegobot/telegram"
)

type Command interface {
	Pattern() string
	Run(string, telegram.Message) (string, error)
}

var (
	token = flag.String("token", "", "Token")
)

func main() {
	flag.Parse()

	telegram.Token = *token

	c := make(chan telegram.Update)
	commands := []Command{}

	//commands = append(commands, &botcommands.Kity{})
	commands = append(commands, &botcommands.Mate{})
	commands = append(commands, &botcommands.Tumblr{})

	go telegram.GetUpdatesChannel(c)

	for update := range c {
		if update.Message.ID > 0 {
			go processMessage(commands, update.Message)
		} else if len(update.InlineQuery.ID) > 0 {
			go processInlineQuery(update.InlineQuery)
		}
	}
}

func processMessage(commands []Command, message telegram.Message) {
	for _, command := range commands {
		if len(message.Text) < len(command.Pattern()) {
			continue
		}

		if message.Text[:len(command.Pattern())] != command.Pattern() {
			continue
		}

		text, err := command.Run(strings.Trim(message.Text[len(command.Pattern()):], " "), message)

		if err != nil {
			telegram.SendMessage(message.Chat.ID, err.Error())
			continue
		}

		telegram.SendMessage(message.Chat.ID, text)
	}
}

// TODO: Do more than just Tumblr
func processInlineQuery(inlineQuery telegram.InlineQuery) error {
	tumblr := botcommands.Tumblr{}
	posts, err := tumblr.Query(inlineQuery.Query)
	photos := []telegram.InlineQueryResultPhoto{}

	if err != nil {
		return err
	}

	for i, post := range posts {
		photos = append(photos, telegram.InlineQueryResultPhoto{
			Type:   "photo",
			ID:     strconv.Itoa(i),
			Photo:  post.High,
			Thumb:  post.Low,
			Width:  post.Width,
			Height: post.Height,
		})
	}

	return telegram.AnswerInlineQuery(inlineQuery.ID, photos)
}
