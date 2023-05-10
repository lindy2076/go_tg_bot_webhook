package main

// import (
// 	"encoding/json"
// )

type TgResponse struct {
	Ok bool `json:"ok"`
}

type BotMessage struct {
	UpdateId int
	Message  struct {
		MessageId int
		From      struct {
			Username  string
			Id        int
			FirstName string
			LastName  string
		}
		Chat struct {
			Id int
		}
		Text string
	}
}
