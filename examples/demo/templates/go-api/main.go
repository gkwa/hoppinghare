package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// {{.ProjectName}} API Server
// Built with Go {{.GoVersion}}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to {{.ProjectName}} API!")
	})

	log.Printf("Starting {{.ProjectName}} on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

