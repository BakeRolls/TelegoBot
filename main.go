package main

import (
	"flag"
	"strings"

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

	c := make(chan telegram.Message)
	commands := []Command{}

	//commands = append(commands, &botcommands.Kity{})
	commands = append(commands, &botcommands.Mate{})

	go telegram.GetUpdatesChannel(c)

	for message := range c {
		go processMessage(commands, message)
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
