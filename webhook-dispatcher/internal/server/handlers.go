package server

import (
	"encoding/json"
	"net/http"

	"github.com/binit2-1/golang-dojo/webhook-dispatcher/internal/dispatcher"
)

type EventPayload struct {
	Message string `json:"message"`
}

func FireWebhookHandler(w http.ResponseWriter, r *http.Request) {

	var payload EventPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	urls := []string{
		"https://httpbin.org/post",
		"https://httpbin.org/post",
		"https://httpbin.org/post",
		"https://httpbin.org/post",
		"https://httpbin.org/post",
		"https://httpbin.org/post",
		"https://httpbin.org/post",
		"https://httpbin.org/post",
		"https://httpbin.org/post",
		"https://httpbin.org/post",
	}

	payloadBytes, _ := json.Marshal(payload)

	go dispatcher.DispatchWebhooks(urls, payloadBytes)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message":"Webhooks are being dispatched in the background"}`))

}
