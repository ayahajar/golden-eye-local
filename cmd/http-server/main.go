package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	log.Printf("Received request: %s\n", string(body))

	fmt.Fprintf(w, "Received: %s\n", string(body))
}

func main() {
	http.HandleFunc("/echo", echoHandler)

	fmt.Printf("Starting server at port 9090\n")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
