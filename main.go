package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info(fmt.Sprintf("[%s] %s", r.Method, r.URL.Path))
	resp := map[string]string{
		"message": "バックエンドからこんにちは 😎",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error(err.Error())
		http.Error(w, "Failed to say hello", http.StatusInternalServerError)
		return
	}
}

func main() {
	address := "0.0.0.0:8080"
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", helloHandler)
	svr := http.Server{
		Addr:    address,
		Handler: mux,
	}

	// TODO: Gracefull Shutdown
	slog.Info(fmt.Sprintf("Server started at port %s 🚀", address))
	if err := svr.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	handler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
