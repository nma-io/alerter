package slack

import (
    "net/http"
    "encoding/json"
    "bytes"
)

func Send(hook string, channel string, username string, message string) (resp bool) {
    resp = true
    data := make(map[string]string)
    data["channel"] = channel
    data["username"] = username
    data["text"] = message
    mkjson, _ := json.Marshal(data)
    _, err := http.Post(hook, "application/json", bytes.NewBuffer(mkjson))
    if err != nil{
       resp = false
    }
    return
}

