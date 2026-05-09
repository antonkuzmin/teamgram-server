package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type telegramRequest struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func main() {
	listen := flag.String("listen", ":8181", "listen address")
	flag.Parse()

	token := os.Getenv("TG_BOT_TOKEN")
	chatID := os.Getenv("TG_CHAT_ID")
	if chatID == "" {
		chatID = "384578982"
	}
	if token == "" {
		log.Fatal("TG_BOT_TOKEN env is required")
	}

	http.HandleFunc("/code", func(w http.ResponseWriter, r *http.Request) {
		phone := r.URL.Query().Get("phone")
		code := r.URL.Query().Get("code")

		log.Printf("OTP request: phone=%s code=%s", phone, code)

		text := fmt.Sprintf("📱 <b>%s</b>\nКод: <code>%s</code>", phone, code)

		body, _ := json.Marshal(telegramRequest{
			ChatID:    chatID,
			Text:      text,
			ParseMode: "HTML",
		})

		url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
		resp, err := http.Post(url, "application/json", bytes.NewReader(body))
		if err != nil {
			log.Printf("telegram error: %v", err)
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
		log.Printf("telegram response: %s", respBody)

		if resp.StatusCode != http.StatusOK {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	log.Printf("listening on %s", *listen)
	log.Fatal(http.ListenAndServe(*listen, nil))
}
