package main

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"todo/api"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func generateCharacterImageHandler() ([]openai.Image, error) {
	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_SECRET_KEY")),
	)

	resp, err := client.Images.Generate(
		context.TODO(),
		openai.ImageGenerateParams{
			Model:  openai.ImageModelGPTImage1,
			Prompt: "A cute Shiba Inu sitting on a cloud, watercolor style",
			Size:   "1024x1024",
		},
	)
	if err != nil {
		slog.Error(err.Error())
		// http.Error(w, "image generation failed", http.StatusInternalServerError)
		return nil, err
	}

	return resp.Data, nil
}

// openai.Images を Base64 で受け取り、保存する関数
func saveGeneratedImages(dirPath string, images []openai.Image) ([]string, error) {
	// ディレクトリを作成
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	var savedPaths []string

	// 各画像をBase64からデコードして保存
	for i, img := range images {
		if img.B64JSON == "" {
			slog.Error("Image B64JSON is empty")
			return nil, errors.New("base64 image response is empty")
		}

		// Base64をデコード
		imgData, err := base64.StdEncoding.DecodeString(img.B64JSON)
		if err != nil {
			return nil, fmt.Errorf("failed to decode base64 image %d: %w", i, err)
		}

		// ファイル名を生成（タイムスタンプ付き）
		filename := fmt.Sprintf("generated_%d_%s.png", i, time.Now().Format("20060102_150405"))
		filePath := filepath.Join(dirPath, filename)

		// ファイルを作成
		file, err := os.Create(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create file: %w", err)
		}
		defer file.Close()

		// 画像データをファイルに書き込み
		_, err = file.Write(imgData)
		if err != nil {
			return nil, fmt.Errorf("failed to write image to file: %w", err)
		}

		savedPaths = append(savedPaths, filePath)
		slog.Info(fmt.Sprintf("Image saved to %s", filePath))
	}

	return savedPaths, nil
}

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

	// imgs, err := generateCharacterImageHandler()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// imgPaths, err := saveGeneratedImages("images", imgs)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(imgPaths)

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
}
