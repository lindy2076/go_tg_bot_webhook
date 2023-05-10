package main

type TgResponse struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
}

type BotMessage struct {
	UpdateId int `json:"update_id"`
	Message  Message
}

type Message struct {
	MessageId int `json:"message_id"`
	From      User
	Chat      struct {
		Id int
	}
	Text string
}

type User struct {
	Username  string
	Id        int
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ReplyKeyboard struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard"`
	OneTimeKeyboard bool               `json:"one_time_keyboard"`
}

type KeyboardButton struct {
	Text string `json:"text"`
}

type InlineKeyboard struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type CallbackQuery struct {
	Id              string
	From            User
	Message         Message `json:"message"`
	Data            string
	InlineMessageId string `json:"inline_message_id"`
}

type BotQuery struct {
	UpdateId      int           `json:"update_id"`
	CallbackQuery CallbackQuery `json:"callback_query"`
}
