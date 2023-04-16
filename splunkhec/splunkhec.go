package splunkhec

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Send(uri string, key string, addr string, msg string) (resp bool) {
	// Forward to HEC Listener
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // Ignore certificate issues
	msg = strings.TrimRight(msg, "\x00")
	msg = strings.Replace(msg, "\n", " ", -1)
	msg = strings.Replace(msg, "\r", " ", -1)
	jsonData := fmt.Sprintf(`{"event": %q, "host": "%s"}`, msg, addr)

	client := &http.Client{Timeout: 3 * time.Second}
	r, _ := http.NewRequest("POST", uri, strings.NewReader(jsonData)) // URL-encoded payload
	r.Header.Add("Authorization", "Splunk "+key)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(jsonData)))
	_, err := client.Do(r)
	if err != nil {
		return false
	}
	return true
}
