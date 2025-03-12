package main

import (
	b64 "encoding/base64"
	"log"
	"net/http"
)

func main() {
	log.Print("creating websocket request")
	client := http.Client{}
	nonce := b64.StdEncoding.EncodeToString([]byte("test-nonce-123"))
	url := "http://localhost:3001/socket"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Adding headers")

	req.Header.Add("Sec-WebSocket-Key", nonce)
	req.Header.Add("Upgrade", "websocket")
	req.Header.Add("Connection", "Upgrade")
	req.Header.Add("Sec-WebSocket-Protocol", "chat, superchat")
	req.Header.Add("Sec-WebSocket-Version", "13")

	log.Print("sending websocket request")
	res, err := client.Do(req)

	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	log.Print("res: %v", res)
}
