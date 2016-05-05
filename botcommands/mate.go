package botcommands

import (
	"errors"

	"git.192k.pw/bake/telegobot/telegram"
)

// Mate Struct
type Mate struct{}

// Pattern defines the message-prefix
func (m *Mate) Pattern() string {
	return "/mate"
}

// Run generates the answer
func (m *Mate) Run(arg string, message telegram.Message) (string, error) {
	if len(arg) == 0 {
		return "", errors.New("What did you drink?")
	}

	return "You drank " + arg + "?", nil
}
