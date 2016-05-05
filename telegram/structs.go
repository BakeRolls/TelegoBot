package telegram

type Response struct {
	OK      bool     `json:"ok"`
	Updates []Update `json:"result"`
}

type Update struct {
	ID          int         `json:"update_id"`
	Message     Message     `json:"message"`
	InlineQuery InlineQuery `json:"inline_query"`
}

type Message struct {
	ID       int      `json:"message_id"`
	From     User     `json:"from"`
	Chat     User     `json:"chat"`
	Date     int      `json:"date"`
	Text     string   `json:"text"`
	Entities []Entity `json:"entities"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"first_name"`
	Username string `json:"username"`
	Type     string `json:"type"`
}

type Entity struct {
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

type Location struct {
	longitude float32 `json:"longitude"`
	latitude  float32 `json:"latitude"`
}
