package dataset

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func sendSyslog(proto string, server string, message string) (resp bool) {
	// Send to Syslog (UDP or TCP)
	// Server is a string in the form of "host:port"
	syslog, _ := net.Dial(proto, server)
	_, err := fmt.Fprintf(syslog, strings.Replace(message, `%`, `%%`, -1))
	if err != nil {
		log.Fatal(err)
	}
	defer syslog.Close()
	return
}
