package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sikalabs/hello-world-server/version"
)

var Counter = 0

func index(w http.ResponseWriter, r *http.Request) {
	Counter++
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

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/version", versionAPI)
	http.HandleFunc("/version", versionAPI)
	http.HandleFunc("/api/livez", livez)
	http.HandleFunc("/livez", livez)
	http.HandleFunc("/api/readyz", readyz)
	http.HandleFunc("/readyz", readyz)
	fmt.Println("Server started.")
	http.ListenAndServe(":8000", nil)
}
