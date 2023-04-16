package dataset

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Send(url string, token string, serverHost string, parser string, logFile string, msg string) (resp bool) {
	// Forward to Dataset/Scalyr
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(msg))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("parser", parser)
	req.Header.Set("server-host", serverHost)
	req.Header.Set("logfile", "emaginedintel")

	// Make HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return false
	}
	defer resp.Body.Close()

	// Read response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return false
	}
	return true
}
