package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("Starting HelloServer...")
	hostname, _ := os.Hostname()
	_ = http.ListenAndServe(":8180", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello! Welcome to %s, you requested: %s\n", hostname, r.URL.Path)
		log.Printf("Received request for path: %s", r.URL.Path)
	}))
}
