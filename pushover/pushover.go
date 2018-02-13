package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const url = "https://api.pushover.net/1/messages.json"

func Send(token string, user string, message string) (resp bool) {
	resp = true
	data := make(map[string]string)
	data["token"] = token
	data["user"] = user
	data["message"] = message
	mkjson, _ := json.Marshal(data)
	_, err := http.Post(url, "application/json", bytes.NewBuffer(mkjson))
	if err != nil {
		resp = false
	}
	return

}
