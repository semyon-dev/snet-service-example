package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Request/Response structures matching the proto definitions
type STTInput struct {
	ModelID string `json:"model_id"`
	Speech  []byte `json:"speech"`
}

type BasicSTTInput struct {
	Text string `json:"text"`
}

type STTResp struct {
	Result string `json:"result"`
}

// HTTP handler for /stt endpoint
func handleSTT(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input STTInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("Invalid JSON input: %v", err), http.StatusBadRequest)
		return
	}

	// Process the STT request (placeholder implementation)
	log.Printf("Processing STT request - ModelID: %s, Speech length: %d bytes", input.ModelID, len(input.Speech))

	response := STTResp{
		Result: fmt.Sprintf("Processed speech with model %s, received %d bytes", input.ModelID, len(input.Speech)),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// HTTP handler for /basic_stt endpoint
func handleBasicSTT(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input BasicSTTInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("Invalid JSON input: %v", err), http.StatusBadRequest)
		return
	}

	// Process the basic STT request (placeholder implementation)
	log.Printf("Processing Basic STT request - Text: %s", input.Text)

	response := STTResp{
		Result: fmt.Sprintf("Processed text: %s", input.Text),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Register HTTP handlers
	http.HandleFunc("/stt", handleSTT)
	http.HandleFunc("/basic_stt", handleBasicSTT)

	// Start HTTP server
	port := ":5001"
	log.Printf("Starting HTTP service on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
