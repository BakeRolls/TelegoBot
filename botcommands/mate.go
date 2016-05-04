package botcommands

import (
	"errors"

	"git.192k.pw/bake/telegobot/telegram"
)

type Mate struct{}

func (*Mate) Pattern() string {
	return "/mate"
}

func (*Mate) Run(arg string, message telegram.Message) (string, error) {
	if len(arg) == 0 {
		return "", errors.New("What did you drink?")
	}

	return "You drank " + arg + "?", nil
}
