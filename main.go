package main

import (
	"fmt"
	"net/http"
	"time"
)

var startTime = time.Now()

func uptime() time.Duration {
	return time.Since(startTime)
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintf(w, "%s Uptime %v", t.Format("2006-01-02T15:04:05"), uptime())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
