package botcommands

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"

	"git.192k.pw/bake/telegobot/telegram"
)

const (
	help = "/kity Hello World\nText ..."
	base = "~/kity"
)

// Kity-Struct
type Kity struct{}

// Pattern defines the message-prefix
func (k *Kity) Pattern() string {
	return "/kity"
}

// Run creates a file in base called 2006-01-02-slug.md
func (k *Kity) Run(arg string, message telegram.Message) (string, error) {
	if len(arg) == 0 {
		return "", errors.New(help)
	}

	args := strings.SplitN(arg, "\n", 2)

	if len(args) < 2 {
		return "", errors.New(help)
	}

	r := regexp.MustCompile(`^a-zA-Z0-9\-\s]+`)

	date := time.Now().Format("2006-01-02")
	title := strings.Trim(args[0], " -")
	text := strings.Trim(args[1], " -")
	slug := strings.ToLower(strings.Replace(r.ReplaceAllString(title, ""), " ", "-", -1))
	file := fmt.Sprintf("%s/%s-%s.md", base, date, slug)

	if len(slug) == 0 {
		return "", errors.New("The slug would be empty.")
	}

	if err := ioutil.WriteFile(file, []byte(text), 0644); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s-%s.md", base, date, slug), nil
}
