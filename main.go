package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
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

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server started.")
	http.ListenAndServe(":8000", nil)
}
