package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader for HTTP -> Websocket forbindelse, default settings
var upgrade = websocket.Upgrader{}

// HTTP Server settings
var httpServer = &http.Server{
	Addr: ":8080",
}

// Websocket forbindelse
func ws(w http.ResponseWriter, r *http.Request) {
	// Oppgrader HTTP til Websocket
	conn, err := upgrade.Upgrade(w, r, nil)

	// Error sjekk av upgrade
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	// Lukke websocket når forbindelse er ferdig
	defer conn.Close()

	// Legg til Websocket funksjoner under her
}

// Klient fil som skal ligge på home/root (/)
func httpHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "client.html")
}

func main() {
	log.Println("Starting service...")

	// Koble funksjoner til URL
	http.HandleFunc("/", httpHome)
	http.HandleFunc("/ws", ws)

	// Starte HTTP server
	httpServer.ListenAndServe()
}
