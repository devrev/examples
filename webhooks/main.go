package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type WebhookPayload struct {
	ID        string    `json:"id"`
	WebhookID string    `json:"webhook_id"`
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
	Challenge string    `json:"challenge"`
}

func main() {
	port := "3000"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	http.HandleFunc("/", webhookHandler)
	log.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, body, "", "  "); err != nil {
		log.Printf("Failed to format JSON: %v\n", err)
		log.Printf("Raw payload: %s\n", string(body))
	} else {
		log.Printf("Received webhook payload:\n%s\n", prettyJSON.String())
	}

	var payload WebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "Failed to parse JSON payload", http.StatusBadRequest)
		return
	}
	if payload.Type == "verify" {
		log.Println("Received verification challenge")
		response := map[string]string{"challenge": payload.Challenge}
		json.NewEncoder(w).Encode(response)
		return
	}

	signature := r.Header.Get("X-DevRev-Signature")
	if !verifySignature(body, signature) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func verifySignature(payload []byte, signature string) bool {
	secret := os.Getenv("WEBHOOK_SECRET")
	if secret == "" {
		panic("WEBHOOK_SECRET not set")
	}
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(payload)
	calculatedSignature := hex.EncodeToString(h.Sum(nil))
	return hmac.Equal([]byte(signature), []byte(calculatedSignature))
}
