package botcommands

// Query for users
// curl https://www.tumblr.com/search/life+is+strange | pup '.tumblelogs_json attr{data-search-tumblelogs-json}' | sed -e 's/&#34;/"/g' | jq '.'

// Query for posts
// curl https://www.tumblr.com/search/life+is+strange | pup '.post_media[data-lightbox] attr{data-lightbox}' | sed -e 's/&#34;/"/g' | jq '.'

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"

	"git.192k.pw/bake/telegobot/telegram"
)

type TumblrPost struct {
	Low    string `json:"low_res"`
	High   string `json:"high_res"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Tumblr Struct
type Tumblr struct{}

// Pattern defines the message-prefix
func (t *Tumblr) Pattern() string {
	return "/tumblr"
}

// Run queries Tumblr and returns the first high-res image
func (t *Tumblr) Run(arg string, message telegram.Message) (string, error) {
	if len(arg) == 0 {
		return "", errors.New("What are you looking for?")
	}

	posts, err := t.Query(arg)

	if err != nil {
		return "", err
	}

	if len(posts) == 0 {
		return "", errors.New("No post found :(")
	}

	return posts[0].High, nil
}

// Query Tumblr
func (t *Tumblr) Query(query string) ([]TumblrPost, error) {
	res, err := http.Get("https://www.tumblr.com/search/" + url.QueryEscape(query))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	r := regexp.MustCompile(`lightbox='(.+?)'`)

	posts := []TumblrPost{}
	matches := r.FindAllStringSubmatch(string(body), -1)

	for _, match := range matches {
		post := TumblrPost{}

		if len(match) < 2 {
			continue
		}

		if err := json.Unmarshal([]byte(match[1]), &post); err != nil {
			continue
		}

		posts = append(posts, post)
	}

	return posts, nil
}
