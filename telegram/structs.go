package telegram

type Response struct {
	OK      bool     `json:"ok"`
	Results []Result `json:"result"`
}

type Result struct {
	ID      int     `json:"update_id"`
	Message Message `json:"message"`
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
