package telegram

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

var (
	// Limit limits the number of updates to be retrieved
	Limit = 100

	// Timeout in seconds for long polling.
	Timeout = 5

	// Base of Telegrams API
	Base = "https://api.telegram.org/bot"

	// Token of your bot
	Token = ""
)

// GetUpdates waits until there is at least one new message
func GetUpdates(offset int, limit int, timeout int) (Response, error) {
	params := map[string]string{
		"offset":  strconv.Itoa(offset),
		"limit":   strconv.Itoa(limit),
		"timeout": strconv.Itoa(timeout),
	}

	body, err := get("getUpdates", params)

	if err != nil {
		return Response{}, err
	}

	res, err := parse(body)

	if err != nil {
		return res, err
	}

	if len(res.Results) <= 0 {
		return res, nil
	}

	return res, nil
}

// GetUpdatesChannel loops over GetUpdates and sends the result Message through a channel
func GetUpdatesChannel(c chan Message) error {
	offset := 0

	for {
		res, err := GetUpdates(offset, 100, 30)

		if err != nil {
			return err
		}

		for _, result := range res.Results {
			c <- result.Message

			offset = result.ID + 1
		}
	}
}

// SendMessage sends a Telegram-message
func SendMessage(chat int, text string) error {
	params := map[string]string{
		"chat_id": strconv.Itoa(chat),
		"text":    text,
	}

	if _, err := get("sendMessage", params); err != nil {
		return err
	}

	return nil
}

func get(method string, params map[string]string) ([]byte, error) {
	vals := url.Values{}

	for key, val := range params {
		vals.Add(key, val)
	}

	res, err := http.Get(Base + Token + "/" + method + "?" + vals.Encode())

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	return body, nil
}

func parse(body []byte) (Response, error) {
	res := Response{}

	err := json.Unmarshal(body, &res)

	if err != nil {
		return res, err
	}

	return res, nil
}
