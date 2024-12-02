package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/sikalabs/hello-world-server/version"
)

//go:embed favicon.ico
var favicon []byte

var Counter = 0
var Color = "black"
var BackgroundColor = "white"
var RunTimestamp time.Time
var Text = "Hello World!"
var Logger zerolog.Logger

var promRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: "hello_world_server",
		Name:      "requests_total",
		Help:      "Total number of requests per endpoint",
	}, []string{"path"})

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
	fmt.Fprint(w, `<!DOCTYPE html>
	<html lang="en"><head>
  <meta charset="UTF-8">
  <title>`+Text+`</title>
	<style>
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
	<link rel="icon" href="/favicon.ico">
	</head>
	<body>
	<section class="center-parent">
		<div class="center-child">
			<h1>
				`+Text+`
			</h1>
			<h2>`+hostname+`</h2>
		</div>
	</section>
	</body></html>
	`)
}

func index(w http.ResponseWriter, r *http.Request) {
	promRequestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()

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
	Logger.Info().Str("server", "main").Str("hostname", hostname).Str("method", r.Method).Str("path", r.URL.Path).Msg(r.Method + " " + r.URL.Path)
}

func versionAPI(w http.ResponseWriter, r *http.Request) {
	promRequestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(map[string]string{
		"version": version.Version,
	})
	fmt.Fprint(w, string(data))
}

func livez(w http.ResponseWriter, r *http.Request) {
	promRequestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(map[string]bool{
		"live": true,
	})
	fmt.Fprint(w, string(data))
}

func readyz(w http.ResponseWriter, r *http.Request) {
	promRequestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(map[string]bool{
		"ready": true,
	})
	fmt.Fprint(w, string(data))
}

func status(w http.ResponseWriter, r *http.Request) {
	promRequestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
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

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	promRequestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	w.Header().Set("Content-Type", "image/x-icon")
	w.WriteHeader(http.StatusOK)
	w.Write(favicon)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	promRequestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
	promhttp.Handler().ServeHTTP(w, r)
}

var MainServer http.Server
var MetricsServer http.Server

func mainServer(port string, hostname string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/api/version", versionAPI)
	mux.HandleFunc("/version", versionAPI)
	mux.HandleFunc("/api/livez", livez)
	mux.HandleFunc("/livez", livez)
	mux.HandleFunc("/api/readyz", readyz)
	mux.HandleFunc("/readyz", readyz)
	mux.HandleFunc("/api/status", status)
	mux.HandleFunc("/status", status)
	mux.HandleFunc("/favicon.ico", faviconHandler)
	MainServer.Handler = mux
	MainServer.Addr = ":" + port
	Logger.Info().Str("hostname", hostname).Str("server", "main").Msg("Server started on 0.0.0.0:" + port + ", see http://127.0.0.1:" + port)
	if err := MainServer.ListenAndServe(); err != http.ErrServerClosed {
		Logger.Fatal().Str("hostname", hostname).Str("server", "main").Msg(err.Error())
	}
}

func metricsServer(port string, hostname string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/metrics", metricsHandler)
	MetricsServer.Addr = ":" + port
	MetricsServer.Handler = mux
	Logger.Info().Str("hostname", hostname).Str("server", "metrics").Msg("Server started on 0.0.0.0:" + port + ", see http://127.0.0.1:" + port + "/metrics")
	if err := MetricsServer.ListenAndServe(); err != http.ErrServerClosed {
		Logger.Fatal().Str("hostname", hostname).Str("server", "main").Msg(err.Error())
	}
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

	textSuffixEnv := os.Getenv("TEXT_SUFFIX")
	if textSuffixEnv != "" {
		Text = Text + " " + textSuffixEnv
	}

	port := "8000"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		port = envPort
	}

	portMetrics := "8001"
	envPortMetrics := os.Getenv("PORT_METRICS")
	if envPortMetrics != "" {
		portMetrics = envPortMetrics
	}

	prometheus.MustRegister(promRequestsTotal)

	Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	RunTimestamp = time.Now()

	hostname, _ := os.Hostname()

	go mainServer(port, hostname)
	go metricsServer(portMetrics, hostname)

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-stop

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	Logger.Info().Str("hostname", hostname).Msg("Shutting down gracefully, press Ctrl+C again to force")
	if err := MainServer.Shutdown(ctx); err != nil {
		Logger.Error().Str("hostname", hostname).Str("server", "main").Msgf("Could not gracefully shutdown the server: %v\n", err)
	}
	if err := MetricsServer.Shutdown(ctx); err != nil {
		Logger.Error().Str("hostname", hostname).Str("server", "metrics").Msgf("Could not gracefully shutdown the server: %v\n", err)
	}
}
