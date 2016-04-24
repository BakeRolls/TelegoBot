package botcommands

import (
	"git.192k.pw/bake/telegobot/telegram"
)

type Mate struct {}

func (*Mate) Pattern() string {
	return "/mate"
}

func (*Mate) Run(arg string, message telegram.Message) string {
	return "You drank " + arg + "?"
}
