package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func setWebhook(botUrl string, host string, port string, certPath string) (*http.Response, error) {
	myaddr := "https://" + host + ":" + port
	fmt.Println("host:", myaddr)

	reqParams := map[string]string{"url": myaddr, "certificate": "@" + certPath}

	return sendCommonRequest(botUrl, "setWebhook", reqParams, true)
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
