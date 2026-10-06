package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "hello"})
}

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/hello", helloHandler)
}

func main() {
	cfg := loadConfig()
	cfg.setupLogger()

	serve(&cfg)
}

func serve(cfg *Config) {
	mux := http.NewServeMux()
	registerRoutes(mux)

	srv := &http.Server{
		Addr:         ":" + cfg.Addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	cfg.InfoLog.Printf("Listening on %s", cfg.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		cfg.ErrorLog.Printf("Error: %s", err)
	}
}
