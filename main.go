package main

import (
	"fmt"
	"log"
	"net/http"

	"golang/cmd/students-api/internal/config"
)

func main() {
	// Load config
	cfg := config.MustLoad()
	fmt.Println("✅ Config loaded successfully")

	// Setup router
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Student API"))
	})

	// Setup server
	server := &http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	fmt.Println("🚀 Server starting on", cfg.HTTPServer.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
