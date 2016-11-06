package telegobot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type TelegoBot struct {
	Token    string
	Limit    int
	Timeout  int
	Base     string
	Test     chan string
	Messages chan Message
	Queries  chan InlineQuery
}

// NewBot builds a new bot with specific token
func NewBot(token string) *TelegoBot {
	return &TelegoBot{
		Token:   token,
		Limit:   100,
		Timeout: 30,
		Base:    "https://api.telegram.org/bot",
	}
}

// GetUpdates loops over getUpdates and sends the Update through a channel
func (t *TelegoBot) GetUpdates() {
	offset := 0

	for {
		res, err := t.getUpdates(offset, t.Limit, t.Timeout)
		if err != nil {
			fmt.Printf("telegram: %s\n", err)
			time.Sleep(time.Second)
			continue
		}

		for _, update := range res.Updates {
			if t.Messages != nil && update.Message.ID > 0 {
				t.Messages <- update.Message
			} else if t.Queries != nil && update.InlineQuery.ID != "" {
				t.Queries <- update.InlineQuery
			}

			offset = update.ID + 1
		}
	}
}

// getUpdates waits until there is at least one new message
func (t *TelegoBot) getUpdates(offset int, limit int, timeout int) (*Response, *Error) {
	res, err := t.get("getUpdates", map[string]string{
		"offset":  strconv.Itoa(offset),
		"limit":   strconv.Itoa(limit),
		"timeout": strconv.Itoa(timeout),
	})
	if err != nil {
		return nil, &Error{false, 400, err.Error()}
	}

	rerr := &Error{}
	if err = json.Unmarshal(res, rerr); err == nil && rerr.Code > 0 {
		return nil, rerr
	}

	response := &Response{}
	if err = json.Unmarshal(res, response); err == nil {
		return response, nil
	}

	return nil, &Error{false, 400, "could not get updates"}
}

func (t *TelegoBot) GetMe() (User, error) {
	user := User{}
	res, err := t.get("getMe", map[string]string{})
	if err != nil {
		return user, err
	}
	if err := json.Unmarshal(res, &user); err != nil {
		return user, err
	}

	return user, nil
}

func (t *TelegoBot) ForwardMessage(chat string, fromChat string, disableNotification bool, messageID int) (Message, error) {
	message := Message{}
	res, err := t.get("forwardMessage", map[string]string{
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

func (t *TelegoBot) GetUserProfilePhotos(user int, offset int, limit int) (UserProfilePhotos, error) {
	photos := UserProfilePhotos{}
	res, err := t.get("getUserProfilePhotos", map[string]string{
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

func (t *TelegoBot) GetFile(id string) (File, error) {
	file := File{}
	res, err := t.get("getFile", map[string]string{
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

func (t *TelegoBot) KickChatMember(chat string, user int) error {
	res, err := t.get("kickChatMember", map[string]string{
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

func (t *TelegoBot) UnbanChatMember(chat string, user int) error {
	res, err := t.get("unbanChatMember", map[string]string{
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

func (t *TelegoBot) AnswerCallbackQuery(id string, text string, alert bool) error {
	res, err := t.get("answerCallbackQuery", map[string]string{
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
func (t *TelegoBot) SendMessage(chat int, text string) error {
	params := map[string]string{
		"chat_id": strconv.Itoa(chat),
		"text":    text,
	}
	if _, err := t.get("sendMessage", params); err != nil {
		return err
	}

	return nil
}

// AnswerInlineQuery answers an InlineQuery
func (t *TelegoBot) AnswerInlineQuery(id string, results []InlineQueryResultPhoto) error {
	jsonResults, err := json.Marshal(results)
	if err != nil {
		return err
	}

	params := map[string]string{
		"inline_query_id": id,
		"results":         string(jsonResults),
	}
	if _, err := t.get("answerInlineQuery", params); err != nil {
		return err
	}

	return nil
}

func (t *TelegoBot) get(method string, params map[string]string) ([]byte, error) {
	vals := url.Values{}
	for key, val := range params {
		vals.Add(key, val)
	}

	client := http.Client{Timeout: time.Duration(t.Timeout) * time.Second}
	res, err := client.Get(fmt.Sprintf("%s%s/%s?%s", t.Base, t.Token, method, vals.Encode()))
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
