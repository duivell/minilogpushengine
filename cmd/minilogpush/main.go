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

func logsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "logs"})
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/v1/", http.StripPrefix("/v1", v1Routes()))
	return mux
}

func v1Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/logs", logsHandler)
	mux.HandleFunc("/hello", helloHandler)
	return mux
}

func main() {
	cfg := loadConfig()
	cfg.setupLogger()

	serve(&cfg)
}

func serve(cfg *Config) {
	mux := newMux()

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
