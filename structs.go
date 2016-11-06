package telegobot

import (
	"fmt"
)

type Error struct {
	OK          bool   `json:"ok"`
	Code        int    `json:"error_code"`
	Description string `json:"description"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Description)
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Chat struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Message struct {
	ID             int             `json:"message_id"`
	From           User            `json:"from"`
	Date           int             `json:"date"`
	Chat           User            `json:"chat"`
	ForwardFrom    User            `json:"forward_from"`
	ForwardDate    int             `json:"forward_date"`
	ReplyToMessage *Message        `json:"reply_to_message"`
	Text           string          `json:"text"`
	Entities       []MessageEntity `json:"entities"`
	Audio          Audio           `json:"audio"`
}

type Audio struct {
	ID        string `json:"id"`
	Size      int    `json:"file_size"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
	MimeType  string `json:"mime_type"`
}

type Document struct {
	ID       string    `json:"id"`
	Size     int       `json:"file_size"`
	Thumb    PhotoSize `json:"thumb"`
	FileName string    `json:"file_name"`
	MimeType string    `json:"mime_type"`
}

type Sticker struct {
	ID     string    `json:"id"`
	Size   int       `json:"file_size"`
	Width  int       `json:"width"`
	Height int       `json:"height"`
	Thumb  PhotoSize `json:"thumb"`
}

type Video struct {
	ID       string    `json:"id"`
	Size     int       `json:"file_size"`
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Duration int       `json:"duration"`
	Thumb    PhotoSize `json:"thumb"`
	MimeType string    `json:"mime_type"`
}

type Voice struct {
	ID       string `json:"id"`
	Size     int    `json:"file_size"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
}

type Contact struct {
	ID          int    `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

type Location struct {
	longitude float32 `json:"longitude"`
	latitude  float32 `json:"latitude"`
}

type Venue struct {
	Location     Location `json:"location"`
	Title        string   `json:"title"`
	Address      string   `json:"address"`
	FoursquareID string   `json:"foursquare_id"`
}

type UserProfilePhotos struct {
	TotalCount int         `json:"total_count"`
	Photos     []PhotoSize `json:"photos"`
}

type File struct {
	ID   string `json:"file_id"`
	Size int    `json:"file_size"`
	Path string `json:"file_path"`
}

type ReplyKeyboardMarkup struct {
	Keyboard        []KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool             `json:"resize_keyboard"`
	OneTimeKeyboard bool             `json:"one_time_keyboard"`
	Selective       bool             `json:"selective"`
}

type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact"`
	RequestLocation bool   `json:"request_location"`
}

type ReplyKeyboardHide struct {
	HideKeyboard bool `json:"hide_keyboard"`
	Selective    bool `json:"selective"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard []KeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text              string `json:"text"`
	URL               string `json:"url"`
	CallbackData      string `json:"callback_data"`
	SwitchInlineQuery string `json:"switch_inline_query"`
}

type CallbackQuery struct {
	ID              string  `json:"id"`
	From            User    `json:"from"`
	Message         Message `json:"message"`
	InlineMessageID string  `json:"inline_message_id"`
	Data            string  `json:"data"`
}

type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective"`
}

type InputFile struct{}

type PhotoSize struct {
	ID     string `json:"file_id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Size   int    `json:"file_size"`
}

type Response struct {
	OK      bool     `json:"ok"`
	Updates []Update `json:"result"`
}

type Update struct {
	ID          int         `json:"update_id"`
	Message     Message     `json:"message"`
	InlineQuery InlineQuery `json:"inline_query"`
}

type MessageEntity struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}

type InlineQuery struct {
	ID       string   `json:"id"`
	From     User     `json:"from"`
	Location Location `json:"location"`
	Query    string   `json:""`
	Offset   string   `json:"offset"`
}

type InlineQueryResultPhoto struct {
	Type   string `json:"type"`
	ID     string `json:"id"`
	Photo  string `json:"photo_url"`
	Thumb  string `json:"thumb_url"`
	Width  int    `json:"photo_width"`
	Height int    `json:"photo_height"`
}
