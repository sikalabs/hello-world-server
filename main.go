package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sikalabs/hello-world-server/version"
)

var Counter = 0
var RunTimestamp time.Time

type StatusResponse struct {
	Counter              int    `json:"counter"`
	Hostname             string `json:"hostname"`
	RunTimestampUnix     int    `json:"run_timestamp_unix"`
	RunTimestampRFC3339  string `json:"run_timestamp_rfc3339"`
	RunTimestampUnixDate string `json:"run_timestamp_unixdate"`
}

func indexPlainText(w http.ResponseWriter) {
	fmt.Fprint(w, "Hello World\n")
}

func indexHTML(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<style>
	html, body {
		height: 100%;
	}
	.center-parent {
		width: 100%;
		height: 100%;
		display: table;
		text-align: center;
	}
	.center-parent > .center-child {
		display: table-cell;
		vertical-align: middle;
	}
	</style>
	<style>
	h1 {
		font-family: Arial;
		font-size: 5em;
	}
	</style>
	<section class="center-parent">
		<div class="center-child">
			<h1>
				Hello World
			</h1>
		</div>
	</section>
	`)
}

func index(w http.ResponseWriter, r *http.Request) {
	Counter++
	// Check if User-Agent header exists
	if userAgentList, ok := r.Header["User-Agent"]; ok {
		// Check if User-Agent header has some data
		if len(userAgentList) > 0 {
			// If User-Agent starts with curl, use plain text
			if strings.HasPrefix(userAgentList[0], "curl") {
				indexPlainText(w)
			} else {
				// If User-Agent header presents and not starts with curl
				// use HTML (Chrome, Safari, Firefox, ...)
				indexHTML(w)
			}
		}
	} else {
		// If User-Agent header doesn't exists, use plain text
		indexPlainText(w)
	}
}

func versionAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(map[string]string{
		"version": version.Version,
	})
	fmt.Fprint(w, string(data))
}

func livez(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(map[string]bool{
		"live": true,
	})
	fmt.Fprint(w, string(data))
}

func readyz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(map[string]bool{
		"ready": true,
	})
	fmt.Fprint(w, string(data))
}

func status(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(StatusResponse{
		Counter:              Counter,
		Hostname:             hostname,
		RunTimestampUnix:     int(RunTimestamp.Unix()),
		RunTimestampRFC3339:  RunTimestamp.Format(time.RFC3339),
		RunTimestampUnixDate: RunTimestamp.Format(time.UnixDate),
	})
	fmt.Fprint(w, string(data))
}

func main() {
	RunTimestamp = time.Now()
	http.HandleFunc("/", index)
	http.HandleFunc("/api/version", versionAPI)
	http.HandleFunc("/version", versionAPI)
	http.HandleFunc("/api/livez", livez)
	http.HandleFunc("/livez", livez)
	http.HandleFunc("/api/readyz", readyz)
	http.HandleFunc("/readyz", readyz)
	http.HandleFunc("/api/status", status)
	http.HandleFunc("/status", status)
	fmt.Println("Server started.")
	http.ListenAndServe(":8000", nil)
}
