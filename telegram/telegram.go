package telegram

import (
	"encoding/json"
	"errors"
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

// GetUpdatesChannel loops over GetUpdates and sends the Update through a channel
func GetUpdatesChannel(c chan Update) error {
	offset := 0

	for {
		res, err := GetUpdates(offset, 100, 30)

		if err != nil {
			return err
		}

		for _, update := range res.Updates {
			c <- update

			offset = update.ID + 1
		}
	}
}

// GetUpdates waits until there is at least one new message
func GetUpdates(offset int, limit int, timeout int) (Response, error) {
	response := Response{}
	res, err := get("getUpdates", map[string]string{
		"offset":  strconv.Itoa(offset),
		"limit":   strconv.Itoa(limit),
		"timeout": strconv.Itoa(timeout),
	})

	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(res, &response); err != nil {
		return response, err
	}

	return response, nil
}

func GetMe() (User, error) {
	user := User{}
	res, err := get("getMe", map[string]string{})

	if err != nil {
		return user, err
	}

	if err := json.Unmarshal(res, &user); err != nil {
		return user, err
	}

	return user, nil
}

func ForwardMessage(chat string, fromChat string, disableNotification bool, messageID int) (Message, error) {
	message := Message{}
	res, err := get("forwardMessage", map[string]string{
		"chat_id":              chat,
		"action":               fromChat,
		"disable_notification": strconv.FormatBool(disableNotification),
		"message_id":           strconv.Itoa(messageID),
	})

	if err != nil {
		return message, err
	}

	if err := json.Unmarshal(res, &message); err != nil {
		return message, err
	}

	return message, nil
}

func GetUserProfilePhotos(user int, offset int, limit int) (UserProfilePhotos, error) {
	photos := UserProfilePhotos{}
	res, err := get("getUserProfilePhotos", map[string]string{
		"user_id": strconv.Itoa(user),
		"offset":  strconv.Itoa(offset),
		"limit":   strconv.Itoa(limit),
	})

	if err != nil {
		return photos, err
	}

	if err := json.Unmarshal(res, &photos); err != nil {
		return photos, err
	}

	return photos, nil
}

func GetFile(id string) (File, error) {
	file := File{}
	res, err := get("getFile", map[string]string{
		"file_id": id,
	})

	if err != nil {
		return file, err
	}

	if err := json.Unmarshal(res, &file); err != nil {
		return file, err
	}

	return file, nil
}

func KickChatMember(chat string, user int) error {
	res, err := get("kickChatMember", map[string]string{
		"chat_id": chat,
		"user_id": strconv.Itoa(user),
	})

	if err != nil {
		return err
	}

	success, err := strconv.ParseBool(string(res))

	if err != nil {
		return err
	}

	if !success {
		return errors.New(strconv.FormatBool(success))
	}

	return nil
}

func UnbanChatMember(chat string, user int) error {
	res, err := get("unbanChatMember", map[string]string{
		"chat_id": chat,
		"user_id": strconv.Itoa(user),
	})

	if err != nil {
		return err
	}

	success, err := strconv.ParseBool(string(res))

	if err != nil {
		return err
	}

	if !success {
		return errors.New(strconv.FormatBool(success))
	}

	return nil
}

func AnswerCallbackQuery(id string, text string, alert bool) error {
	res, err := get("answerCallbackQuery", map[string]string{
		"file_id":    id,
		"text":       text,
		"show_alert": strconv.FormatBool(alert),
	})

	if err != nil {
		return err
	}

	success, err := strconv.ParseBool(string(res))

	if err != nil {
		return err
	}

	if !success {
		return errors.New(strconv.FormatBool(success))
	}

	return nil
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

// AnswerInlineQuery answers an InlineQuery
func AnswerInlineQuery(id string, results []InlineQueryResultPhoto) error {
	jsonResults, err := json.Marshal(results)

	if err != nil {
		return err
	}

	params := map[string]string{
		"inline_query_id": id,
		"results":         string(jsonResults),
	}

	if _, err := get("answerInlineQuery", params); err != nil {
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
