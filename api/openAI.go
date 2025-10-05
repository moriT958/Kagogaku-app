package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
)

var (
	model           = "gpt-image-1"
	outputImageSize = "1024x1024"
)

// fetchImageFromURL は URL から画像をダウンロードして一時ファイルに保存します
func fetchImageFromURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch image from URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch image: HTTP %d", resp.StatusCode)
	}

	// Content-Type からファイル拡張子を判定
	contentType := resp.Header.Get("Content-Type")
	var ext string
	switch contentType {
	case "image/png":
		ext = ".png"
	case "image/jpeg", "image/jpg":
		ext = ".jpg"
	case "image/webp":
		ext = ".webp"
	default:
		// Content-Type が不明な場合は URL から拡張子を取得
		ext = filepath.Ext(url)
		if ext == "" {
			ext = ".png" // デフォルト
		}
	}

	// 一時ファイルに保存
	tmpFile, err := os.CreateTemp("images", "fetched-*"+ext)
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}
	defer tmpFile.Close()

	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		os.Remove(tmpFile.Name())
		return "", fmt.Errorf("failed to save image: %w", err)
	}

	return tmpFile.Name(), nil
}

func updateCharacterImage(imagePath string, prompt string) (string, error) {
	/*
		OpenAI の Edit で画像を変換する処理

		引数1: 変換前の画像パス (ローカルパスまたはURL)
		引数2: プロンプト

		返り値: 変換後の画像パス, エラー
	*/

	// URL かローカルパスかを判定
	var actualImagePath string
	if strings.HasPrefix(imagePath, "http://") || strings.HasPrefix(imagePath, "https://") {
		// URL の場合: ダウンロードして一時ファイルに保存
		tmpPath, err := fetchImageFromURL(imagePath)
		if err != nil {
			return "", err
		}
		actualImagePath = tmpPath
		defer os.Remove(tmpPath) // 処理後に削除
	} else {
		// ローカルパスの場合
		actualImagePath = imagePath
	}

	// multipart/form-data 用のリクエストボディ作成
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 変換前の画像ファイルを読み込み
	imageFile, err := os.Open(actualImagePath)
	if err != nil {
		return "", fmt.Errorf("failed to open image file: %w", err)
	}
	defer imageFile.Close()

	// ファイル拡張子からMIMEタイプを判定
	ext := strings.ToLower(filepath.Ext(actualImagePath))
	var contentType string
	switch ext {
	case ".png":
		contentType = "image/png"
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".webp":
		contentType = "image/webp"
	default:
		return "", fmt.Errorf("unsupported image format: %s (supported: .png, .jpg, .jpeg, .webp)", ext)
	}

	// 画像フィールドを追加（MIMEタイプを明示的に設定）
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="image"; filename="%s"`, filepath.Base(actualImagePath)))
	h.Set("Content-Type", contentType)
	imagePart, err := writer.CreatePart(h)
	if err != nil {
		return "", fmt.Errorf("failed to create image form field: %w", err)
	}
	if _, err := io.Copy(imagePart, imageFile); err != nil {
		return "", fmt.Errorf("failed to copy image data: %w", err)
	}

	// プロンプトフィールドを追加
	if err := writer.WriteField("prompt", prompt); err != nil {
		return "", fmt.Errorf("failed to write prompt field: %w", err)
	}

	// モデルフィールドを追加
	if err := writer.WriteField("model", model); err != nil {
		return "", fmt.Errorf("failed to write model field: %w", err)
	}

	// サイズフィールドを追加
	if err := writer.WriteField("size", outputImageSize); err != nil {
		return "", fmt.Errorf("failed to write size field: %w", err)
	}

	// 生成枚数フィールドを追加
	if err := writer.WriteField("n", "1"); err != nil {
		return "", fmt.Errorf("failed to write n field: %w", err)
	}

	// マルチパートライターをクローズ
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close multipart writer: %w", err)
	}

	// リクエストの作成
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/images/edits", &requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// ヘッダーを設定
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_SECRET_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// X-Request-ID（トレースID）を出力
	if requestID := resp.Header.Get("X-Request-ID"); requestID != "" {
		slog.Info(fmt.Sprintf("OpenAI Image Edit Request ID: %s", requestID))
	}

	// ステータスコードをチェック
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("OpenAI API returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	// レスポンスボディの処理
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var result struct {
		Created int64 `json:"created"`
		Data    []struct {
			B64Json string `json:"b64_json"`
		} `json:"data"`
	}

	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return "", fmt.Errorf("failed to decode response body: %w", err)
	}

	// データが空でないかチェック
	if len(result.Data) == 0 {
		return "", fmt.Errorf("no image data returned from OpenAI API")
	}

	// base64 画像を images に保存する
	imageData, err := base64.StdEncoding.DecodeString(result.Data[0].B64Json)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 image: %w", err)
	}

	// 出力ファイルパスを生成
	imgFileExt := filepath.Ext(actualImagePath)
	baseName := strings.TrimSuffix(filepath.Base(actualImagePath), imgFileExt)
	outputPath := filepath.Join("images", baseName+"-edited"+imgFileExt)

	// ファイルに保存
	if err := os.WriteFile(outputPath, imageData, 0644); err != nil {
		return "", fmt.Errorf("failed to save image file: %w", err)
	}

	slog.Info(fmt.Sprintf("Image saved to: %s", outputPath))
	return outputPath, nil
}

func buildPrompt(job Job) string {

	var foodsList strings.Builder
	for _, food := range job.Data.Foods {
		foodsList.WriteString(fmt.Sprintf("- %s\n", food))
	}

	prompt := fmt.Sprintf(`
# キャラクターの見た目を変換するタスク

画像で与えたキャラクターは、昨日１日、以下のような生活を送りました。
生活を踏まえて、以下の要素を画像に追加してキャラクターの体調を表してください。
健康的であれば健康そうな画像に変換してください。

- 寝不足度
- 食べ過ぎ度

## 睡眠時間

昨晩は %f 時間寝ました。

## 食事

昨日は以下の食事を食べました。

%s`, job.Data.SleepDuration, foodsList.String())

	return prompt
}
