package teams

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Send(hook string, message string) (resp bool) {
	resp = true
	data := make(map[string]string)
	data["text"] = message
	mkjson, _ := json.Marshal(data)
	_, err := http.Post(hook, "application/json", bytes.NewBuffer(mkjson))
	if err != nil {
		resp = false
	}
	return
}
