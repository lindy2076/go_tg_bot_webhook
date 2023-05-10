package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func sendCommonRequest(botUrl string, method string, form map[string]string, isMultiPartForm bool) (*http.Response, error) {
	query, err := url.JoinPath(botUrl, method)
	if err != nil {
		log.Println("joinpathsendrequest error")
	}

	if isMultiPartForm {
		contentType, body, err := createForm(form)
		if err != nil {
			log.Println("createformsendcommonrequest error")
		}

		resp, err := http.Post(query, contentType, body)
		if err != nil {
			log.Println("hmm", err)
		}

		return resp, err
	}

	var formJsoned bytes.Buffer
	err = json.NewEncoder(&formJsoned).Encode(form)

	resp, err := http.Post(query, "application/json", &formJsoned)

	if err != nil {
		log.Println("httppostsendrequest error")
	}

	return resp, err
}

func sendRequest(botUrl string, method string, form map[string]string) (*http.Response, error) {
	return sendCommonRequest(botUrl, method, form, false)
}

func statusAndDescription(resp *http.Response) TgResponse {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	tgResp := TgResponse{}
	if err := json.Unmarshal(body, &tgResp); err != nil {
		log.Fatal(err)
	}
	return tgResp
}

func replyWithParams(params map[string]string) *http.Response {
	resp, err := sendRequest(BOT_API, "sendMessage", params)
	if err != nil {
		log.Println("replywithtext error")
	}
	return resp
}

func replyWithText(chatId int, text string) *http.Response {
	responseParams := map[string]string{"chat_id": fmt.Sprintf("%v", chatId)}
	responseParams["text"] = text
	return replyWithParams(responseParams)
}

func replyWithReplyKeyboard(chatId int, text string, keyboard *ReplyKeyboard) *http.Response {
	responseParams := map[string]string{"chat_id": fmt.Sprintf("%v", chatId)}
	responseParams["text"] = text
	keyboardJsoned, err := json.Marshal(keyboard)
	responseParams["reply_markup"] = string(keyboardJsoned)
	resp, err := sendRequest(BOT_API, "sendMessage", responseParams)
	if err != nil {
		log.Println("replywithreplykeyboard error")
	}
	return resp
}

func replyWithInlineKeyboard(chatId int, text string, keyboard *InlineKeyboard) *http.Response {
	responseParams := map[string]string{"chat_id": fmt.Sprintf("%v", chatId)}
	responseParams["text"] = text
	keyboardJsoned, err := json.Marshal(keyboard)
	fmt.Println(string(keyboardJsoned))
	responseParams["reply_markup"] = string(keyboardJsoned)
	resp, err := sendRequest(BOT_API, "sendMessage", responseParams)
	if err != nil {
		log.Println("replywithinlinekeyboard error")
	}
	return resp
}

func changeMessage(chatId int, msgId int, text string, keyboard *InlineKeyboard) *http.Response {
	responseParams := map[string]string{"message_id": fmt.Sprintf("%v", msgId)}
	responseParams["chat_id"] = fmt.Sprintf("%v", chatId)
	responseParams["text"] = text
	fmt.Println(responseParams)

	keyboardJsoned, err := json.Marshal(keyboard)
	fmt.Println(string(keyboardJsoned))
	responseParams["reply_markup"] = string(keyboardJsoned)

	resp, err := sendRequest(BOT_API, "editMessageText", responseParams)
	if err != nil {
		log.Println("replywithinlinekeyboard error")
	}
	return resp
}

func answerCallback(queryId string) *http.Response {
	resp, err := sendRequest(BOT_API, "answerCallbackQuery", map[string]string{"callback_query_id": queryId})
	if err != nil {
		log.Println("err unmarshalling")
	}
	return resp
}
