package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
	"todo/api"

	"github.com/joho/godotenv"
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
	// ジョブワーカーを起動
	api.StartJobWorker()

	address := "0.0.0.0:8080"
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", helloHandler)
	mux.HandleFunc("POST /character/new", api.NewCharacterPost)
	mux.HandleFunc("GET /character/{id}", api.CharacterGet)
	mux.HandleFunc("PATCH /character/{id}/sleep", api.CharacterSleepPatch)
	mux.HandleFunc("PATCH /character/{id}/wake-up", api.CharacterWakeUpPatch)
	mux.HandleFunc("GET /train-status/{jobId}", api.TrainJobStatusGet)
	mux.HandleFunc("POST /character/{id}/eat", api.CharacterEatPost)
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
	// ロガーの設定
	handler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// .env ファイルからの環境変数読み込み
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// タイムゾーンを日本時間に設定
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal("Failed to load timezone:", err)
	}
	time.Local = loc
}
