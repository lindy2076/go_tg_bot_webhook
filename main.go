package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	TG_API_URL    = os.Getenv("TG_API_URL")
	BOT_TOKEN     = os.Getenv("BOT_TOKEN")
	HOST_ADRESS   = os.Getenv("HOST_ADRESS")
	HOST_PORT     = os.Getenv("HOST_PORT")
	BOT_API, _    = url.JoinPath(TG_API_URL, "bot"+BOT_TOKEN)
	CERT_KEY_PATH = os.Getenv("CERT_KEY_PATH")
	CERT_PATH     = os.Getenv("CERT_PATH")
)

func main() {
	if BOT_TOKEN == "" {
		log.Fatal("No token provided")
	}
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// setting webhook
	resp, err := setWebhook(BOT_API, HOST_ADRESS, HOST_PORT, CERT_PATH)
	if err != nil {
		infoLog.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body, resp)
	tgResp := TgResponse{}
	if err := json.Unmarshal(body, &tgResp); err != nil {
		infoLog.Fatal(err)
	}
	if tgResp.Ok {
		infoLog.Println("Webhook set")
	} else {
		infoLog.Println(body, "sf")
		infoLog.Fatal("Webhook was not set", tgResp.Ok)
	}
	infoLog.Println("webhook:", tgResp.Ok)

	http.HandleFunc("/", handleEvent)
	http.ListenAndServeTLS(fmt.Sprintf("%s:%s", HOST_ADRESS, HOST_PORT), CERT_PATH, CERT_KEY_PATH, nil)

	// deleting webhook
	time.Sleep(time.Second * 3)

	resp, err = deleteWebhook(BOT_API)
	if err != nil {
		infoLog.Println(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	tgResp = TgResponse{}
	if err := json.Unmarshal(body, &tgResp); err != nil {
		infoLog.Fatal(err)
	}
	if tgResp.Ok {
		infoLog.Println("Webhook unset")
	} else {
		infoLog.Fatal("Webhook can not be unset")
	}
}

// TODO
func handleEvent(w http.ResponseWriter, r *http.Request) {
	text, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var botText BotMessage
	err = json.Unmarshal(text, &botText)
	username := botText.Message.From.Username
	chatUser := botText.Message.From.Id
	chatGroup := botText.Message.Chat.Id
	messageId := botText.Message.MessageId
	botCommand := strings.Split(botText.Message.Text, "@")[0]
	commandText := strings.Split(botText.Message.Text, " ")
	log.Printf("%s %d %d %d %s %s", username, chatUser, chatGroup, messageId, botCommand, commandText)
}
