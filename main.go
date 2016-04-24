package main

import (
	"flag"
	"git.192k.pw/bake/telegobot/botcommands"
	"git.192k.pw/bake/telegobot/telegram"
	"log"
)

type Command interface {
	Pattern() string
	Run(string, telegram.Message) string
}

func main() {
	token := flag.String("token", "", "Token")

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

				text := command.Run(result.Message.Text[len(command.Pattern())+1:], result.Message)

				log.Println(result.Message.Chat.Username + " drank " + result.Message.Text[len(command.Pattern())+1:])

				telegram.SendMessage(result.Message.Chat.ID, text)
			}
		}

		if len(res.Results) > 0 {
			offset = res.Results[len(res.Results)-1].ID + 1
		}
	}
}
