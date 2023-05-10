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
	tgResp := statusAndDescription(resp)
	if tgResp.Ok {
		infoLog.Println("Webhook set")
	} else {
		infoLog.Fatal("Webhook was not set: ", tgResp.Description)
	}

	http.HandleFunc("/", handleEvent)
	http.ListenAndServeTLS(fmt.Sprintf("%s:%s", HOST_ADRESS, HOST_PORT), CERT_PATH, CERT_KEY_PATH, nil)

	// deleting webhook
	time.Sleep(time.Second * 3)

	resp, err = deleteWebhook(BOT_API)
	if err != nil {
		infoLog.Println(err)
	}
	tgResp = statusAndDescription(resp)
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

	message := messageFromJson(text).Message
	log.Println(message, "msg")

	if message.Chat.Id == 0 && message.From.Id == 0 && message.From.Username == "" {
		log.Println("callback incoming")
		handleCallback(text)
	} else {
		handleCommand(message)
	}
}

func handleCommand(msg Message) {
	username := msg.From.Username
	chatId := msg.From.Id
	chatGroup := msg.Chat.Id
	messageId := msg.MessageId
	botCommand := strings.ToLower(msg.Text)

	log.Printf("uname: %s chatid:%d chatgroup:%d msgid:%d cmd:%s", username, chatId, chatGroup, messageId, botCommand)

	firstName := msg.From.FirstName

	switch strings.ToLower(msg.Text) {
	case "привет":
		reply := "привет, " + firstName + "!"
		resp := replyWithText(chatId, reply)
		log.Println(statusAndDescription(resp))

	case "/start":
		reply := STARTUP_MESSAGE
		kb := createWelcomeKeyboard()
		resp := replyWithReplyKeyboard(chatId, reply, &kb)
		log.Println(statusAndDescription(resp))
	case "команды":
		reply := COMMANDS
		resp := replyWithText(chatId, reply)
		log.Println(statusAndDescription(resp))
	case "нуклеотиды":
		reply := "Нажми на нуклеотид, о котором ты хочешь узнать подробнее.."
		kb := createFirstLayerKeyboard()
		resp := replyWithInlineKeyboard(chatId, reply, &kb)
		log.Println(statusAndDescription(resp))
	default:
		reply := "не понял..."
		resp := replyWithText(chatId, reply)
		log.Println(statusAndDescription(resp))
	}
}

func handleCallback(text []byte) {
	query := queryFromJson(text).CallbackQuery
	log.Println(query, "query")

	queryId := query.Id

	callbackData := query.Data

	chatId := query.Message.Chat.Id
	msgId := query.Message.MessageId
	fmt.Println(query.Message)
	inlineMessageId := query.InlineMessageId
	fmt.Println("INLMMESD", chatId, msgId, inlineMessageId)

	switch callbackData {
	case "layer1button1":
		reply := LAYER1BUTTON1TEXT
		kb := createSecondLayerKeyboard1()
		resp := changeMessage(chatId, msgId, reply, &kb)
		log.Println(statusAndDescription(resp))

	case "layer1button2":
		reply := LAYER1BUTTON2TEXT
		kb := createSecondLayerKeyboard2()
		resp := changeMessage(chatId, msgId, reply, &kb)
		log.Println(statusAndDescription(resp))

	case "layer1button3":
		reply := LAYER1BUTTON3TEXT
		kb := createSecondLayerKeyboard3()
		resp := changeMessage(chatId, msgId, reply, &kb)
		log.Println(statusAndDescription(resp))

	case "layer1button4":
		reply := LAYER1BUTTON4TEXT
		kb := createSecondLayerKeyboard4()
		resp := changeMessage(chatId, msgId, reply, &kb)
		log.Println(statusAndDescription(resp))

	case "layer1":
		reply := "Нажми на нуклеотид, о котором ты хочешь узнать подробнее.."
		kb := createFirstLayerKeyboard()
		resp := changeMessage(chatId, msgId, reply, &kb)
		log.Println(statusAndDescription(resp))

	case "layer2adenin":
		reply := ADENIN_WIKI
		resp := replyWithText(chatId, reply)
		log.Println(statusAndDescription(resp))
	case "layer2guanin":
		reply := GUANIN_WIKI
		resp := replyWithText(chatId, reply)
		log.Println(statusAndDescription(resp))
	case "layer2timin":
		reply := TIMIN_WIKI
		resp := replyWithText(chatId, reply)
		log.Println(statusAndDescription(resp))
	case "layer2citozin":
		reply := CITOZIN_WIKI
		resp := replyWithText(chatId, reply)
		log.Println(statusAndDescription(resp))

	}
	resp := answerCallback(queryId)
	log.Println(statusAndDescription(resp))

}

func messageFromJson(text []byte) BotMessage {
	var input BotMessage
	err := json.Unmarshal(text, &input)
	if err != nil {
		log.Println("err unmarshalling")
	}
	return input
}

func queryFromJson(text []byte) BotQuery {
	var input BotQuery
	err := json.Unmarshal(text, &input)
	if err != nil {
		log.Println("err unmarshalling")
	}
	return input
}
