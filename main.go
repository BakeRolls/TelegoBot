package main

import (
	"flag"
	"log"
	"strings"
	"git.192k.pw/bake/telegobot/botcommands"
	"git.192k.pw/bake/telegobot/telegram"
)

type Command interface {
	Pattern() string
	Run(string, telegram.Message) string
}

var (
	token = flag.String("token", "", "Token")
)

func main() {
	flag.Parse()

	telegram.Token = *token

	loop()
}

func loop() {
	offset := 0
	commands := []Command{}

	commands = append(commands, &botcommands.Mate{})

	for {
		res, err := telegram.GetUpdates(offset, 100, 30)

		if err != nil {
			log.Println(err.Error())
			return
		}

		for _, result := range res.Results {
			for _, command := range commands {
				if len(result.Message.Text) < len(command.Pattern()) {
					break
				}

				if result.Message.Text[:len(command.Pattern())] != command.Pattern() {
					break
				}

				go runCommand(command, result.Message.Text[len(command.Pattern()):], result.Message)
			}
		}

		if len(res.Results) > 0 {
			offset = res.Results[len(res.Results)-1].ID + 1
		}
	}
}

func runCommand(command Command, arg string, message telegram.Message) {
	text := command.Run(strings.Trim(arg, " "), message)

	telegram.SendMessage(message.Chat.ID, text)
}
