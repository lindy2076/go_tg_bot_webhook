package main

import (
	"bytes"
	"encoding/json"
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
