package botcommands

import (
	"git.192k.pw/bake/telegobot/telegram"
)

type Mate struct {}

func (*Mate) Pattern() string {
	return "/mate"
}

func (*Mate) Run(arg string, message telegram.Message, callback func(text string, message telegram.Message)) {
	callback("You drank " + arg + "?", message)
}
