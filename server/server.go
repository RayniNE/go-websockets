package server

import "net/http"

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "static/html/index.html")
}

func ServerWebSocket(w http.ResponseWriter, r *http.Request) {}
