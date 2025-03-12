package server

import (
	sha1 "crypto/sha1"
	b64 "encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

//type WebsocketClient struct {
//	Client *http.Client
//	LOG    *log.Logger
//}
//
//func NewClient() *WebsocketClient {
//	client := http.Client{
//		Transport: &http.Transport{
//			MaxConnsPerHost: 10,
//		},
//	}
//
//	return &WebsocketClient{
//		Client: &client,
//		LOG:    log.New(os.Stdout, "Websocket test - ", log.LstdFlags),
//	}
//}

func StartServer() {
	http.HandleFunc("/socket", webSocketHandler)
	port := 3001
	log.Printf("Starting server on port %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	guid, err := uuid.NewUUID()
	if err != nil {
		log.Printf("Error generating GUID: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	websocketKey := r.Header.Get("Sec-WebSocket-Key")
	if websocketKey == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	serverWsResKey := encodeWebsocketResponseKey(guid.String(), websocketKey)

	w.Header().Set("Upgrade", "websocket")
	w.Header().Set("Connection", "Upgrade")
	w.Header().Set("Sec-WebSocket-Accept", serverWsResKey)
	w.Header().Set("Set-websocket-Protocol", "chat")

	w.WriteHeader(http.StatusSwitchingProtocols)
}

// the server response takes the websocket key from a client
// generates a guid concatenates the two
// sha-1 encodes the result and passes to client
func encodeWebsocketResponseKey(guid string, websocketClientKey string) string {
	hash := sha1.New()
	if _, err := hash.Write([]byte(websocketClientKey + guid)); err != nil {
		log.Fatal(err)
	}
	return b64.StdEncoding.EncodeToString(hash.Sum(nil))
}
