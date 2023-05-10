package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func setWebhook(botUrl string, host string, port string) (*http.Response, error) {
	myaddr := "https://" + host + ":" + port
	fmt.Println("host:", myaddr)

	reqParams := url.Values{}
	reqParams.Set("url", myaddr)

	addr, err := url.JoinPath(BOT_API, "setWebhook")
	if err != nil {
		log.Println("hmm", err)
	}

	query := addr + "?" + reqParams.Encode()

	resp, err := http.Get(query)
	if err != nil {
		log.Println("hmm", err)
	}

	return resp, err
}

func deleteWebhook(botUrl string) (*http.Response, error) {
	query, err := url.JoinPath(botUrl, "deleteWebhook")
	if err != nil {
		log.Println("hmm", err)
	}
	resp, err := http.Get(query)
	if err != nil {
		log.Println("hmm", err)
	}

	return resp, err
}
