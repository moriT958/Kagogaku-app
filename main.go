package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
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
	// MySQL の疎通確認
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("MYSQL_HOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", username, password, host, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("fail to connect DB:", err)
		return
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping MySQL: ", err)
	}

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
	// ロガーの設定
	handler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// .env ファイルからの環境変数読み込み
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
