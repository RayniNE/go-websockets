package main

import (
	"net/http"

	"github.com/raynine/go-websockets/server"
)

func main() {
	http.HandleFunc("/", server.ServeIndex)
	http.HandleFunc("/ws", server.ServerWebSocket)
}
