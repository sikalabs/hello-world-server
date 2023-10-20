package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/sikalabs/hello-world-server/version"
)

var Counter = 0
var Color = "black"
var BackgroundColor = "white"
var RunTimestamp time.Time
var Text = "Hello World!"
var Logger zerolog.Logger

type StatusResponse struct {
	Counter              int    `json:"counter"`
	BackgroundColor      string `json:"background_color"`
	Hostname             string `json:"hostname"`
	RunTimestampUnix     int    `json:"run_timestamp_unix"`
	RunTimestampRFC3339  string `json:"run_timestamp_rfc3339"`
	RunTimestampUnixDate string `json:"run_timestamp_unixdate"`
}

func indexPlainText(w http.ResponseWriter, hostname string) {
	fmt.Fprintf(w, "%s %s", Text, hostname)
	if BackgroundColor != "white" {
		fmt.Fprintf(w, " (%s)", BackgroundColor)
	}
	fmt.Fprint(w, "\n")
}

func indexHTML(w http.ResponseWriter, hostname string) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<style>
	html, body {
		height: 100%;
		color: `+Color+`;
		background-color: `+BackgroundColor+`
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
	h2 {
		font-family: Arial;
		font-size: 2em;
	}
	</style>
	<section class="center-parent">
		<div class="center-child">
			<h1>
				`+Text+`
			</h1>
			<h2>`+hostname+`</h2>
		</div>
	</section>
	`)
}

func index(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	Counter++
	// Check if User-Agent header exists
	if userAgentList, ok := r.Header["User-Agent"]; ok {
		// Check if User-Agent header has some data
		if len(userAgentList) > 0 {
			// If User-Agent starts with curl, use plain text
			if strings.HasPrefix(userAgentList[0], "curl") {
				indexPlainText(w, hostname)
			} else {
				// If User-Agent header presents and not starts with curl
				// use HTML (Chrome, Safari, Firefox, ...)
				indexHTML(w, hostname)
			}
		}
	} else {
		// If User-Agent header doesn't exists, use plain text
		indexPlainText(w, hostname)
	}
	Logger.Info().Str("hostname", hostname).Str("method", r.Method).Str("path", r.URL.Path).Msg(r.Method + " " + r.URL.Path)
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
		BackgroundColor:      BackgroundColor,
		RunTimestampUnix:     int(RunTimestamp.Unix()),
		RunTimestampRFC3339:  RunTimestamp.Format(time.RFC3339),
		RunTimestampUnixDate: RunTimestamp.Format(time.UnixDate),
	})
	fmt.Fprint(w, string(data))
}

func main() {
	backgroundCounterEnv := os.Getenv("BACKGROUND_COLOR")
	if backgroundCounterEnv != "" {
		BackgroundColor = backgroundCounterEnv
	}

	counterEnv := os.Getenv("COLOR")
	if counterEnv != "" {
		Color = counterEnv
	}

	textEnv := os.Getenv("TEXT")
	if textEnv != "" {
		Text = textEnv
	}

	port := "8000"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		port = envPort
	}

	Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
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
	hostname, _ := os.Hostname()
	Logger.Info().Str("hostname", hostname).Msg("Server started on 0.0.0.0:" + port + ", see http://127.0.0.1:" + port)
	http.ListenAndServe("0.0.0.0:"+port, nil)
}
