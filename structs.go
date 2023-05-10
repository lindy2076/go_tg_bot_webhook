package main

// import (
// 	"encoding/json"
// )

type TgResponse struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
}

type BotMessage struct {
	UpdateId int `json:"update_id"`
	Message  struct {
		MessageId int `json:"message_id"`
		From      struct {
			Username  string
			Id        int
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		}
		Chat struct {
			Id int
		}
		Text string
	}
}
