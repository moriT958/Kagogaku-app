package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func NewCharacterPost(w http.ResponseWriter, r *http.Request) {
	slog.Info(fmt.Sprintf("[%s] %s", r.Method, r.URL.Path))

	// リクエストのバリデーション
	var req PostNewCharacterReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error(err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// KVS に保存
	mu.Lock()
	characterStore[req.Id] = &Character{
		Id:         req.Id,
		Name:       req.Name,
		Appearance: req.Appearance,
		Status:     healthy,
		Foods:      []string{},
	}
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
}

func CharacterGet(w http.ResponseWriter, r *http.Request) {
	slog.Info(fmt.Sprintf("[%s] %s", r.Method, r.URL.Path))

	// キャラクターを取得
	id := r.PathValue("id")
	mu.RLock()
	char, exists := characterStore[id]
	mu.RUnlock()

	if !exists {
		slog.Error("Character not found")
		http.Error(w, "Character not found", http.StatusNotFound)
		return
	}

	// キャラクターの健康状態
	var charHealthResp string
	switch char.Status {
	case healthy:
		charHealthResp = "健康"
	case lackOfSleep:
		charHealthResp = "寝不足"
	case overeating:
		charHealthResp = "食べ過ぎ"
	}

	// CharacterResp を JSON で返す
	resp := CharacterResp{
		Name:       char.Name,
		Status:     charHealthResp,
		SleepTime:  char.SleepTime,
		WakeUpTime: char.WakeUpTime,
		Foods:      char.Foods,
		Appearance: char.Appearance,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error(err.Error())
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func CharacterSleepPatch(w http.ResponseWriter, r *http.Request) {
	slog.Info(fmt.Sprintf("[%s] %s", r.Method, r.URL.Path))

	// キャラクターを取得
	id := r.PathValue("id")
	mu.Lock()
	defer mu.Unlock()
	char, exists := characterStore[id]
	if !exists {
		slog.Error("character not found")
		http.Error(w, "Character not found", http.StatusNotFound)
		return
	}

	// 現在時刻を取得しキャラデータを更新
	char.SleepTime = time.Now()

	w.WriteHeader(http.StatusOK)
}

func CharacterWakeUpPatch(w http.ResponseWriter, r *http.Request) {
	slog.Info(fmt.Sprintf("[%s] %s", r.Method, r.URL.Path))

	// キャラクターを取得
	id := r.PathValue("id")

	// ロックスコープを限定（デッドロック回避）
	mu.Lock()
	char, exists := characterStore[id]
	if !exists {
		mu.Unlock()
		slog.Error("Character not found")
		http.Error(w, "Character not found", http.StatusNotFound)
		return
	}

	// 現在時刻を取得し、睡眠時間を計算
	char.WakeUpTime = time.Now()
	sleepDuration := float32(char.WakeUpTime.Sub(char.SleepTime).Hours())

	foods := char.Foods

	// 日が更新されたので、昨日までの情報 (Id, Name 以外) を初期化する
	char.Status = healthy
	char.SleepTime = time.Time{}
	char.WakeUpTime = time.Time{}
	char.Foods = []string{}
	mu.Unlock()

	// NOTE: ロック外でジョブ登録しないとデッドロックする
	jobId, err := EnqueueTrainJob(JobData{
		CharacterId:   char.Id,
		SleepDuration: sleepDuration,
		Foods:         foods,
	})
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Failed to create job", http.StatusInternalServerError)
		return
	}

	resp := PostWakeUpResp{
		JobId: jobId,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error(err.Error())
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func TrainJobStatusGet(w http.ResponseWriter, r *http.Request) {
	slog.Info(fmt.Sprintf("[%s] %s", r.Method, r.URL.Path))

	// ジョブキューからジョブを取得
	jobId := r.PathValue("jobId")
	job, err := DequeueTrainJob(jobId)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Failed to get job", http.StatusInternalServerError)
		return
	}

	// ジョブのステータスを返す
	var resp GetTrainStatusResp
	switch job.Status {
	case processing:
		resp = GetTrainStatusResp{
			JobStatus: "Processing",
		}
	case failed:
		resp = GetTrainStatusResp{
			JobStatus: "Failed",
		}
	case completed:
		resp = GetTrainStatusResp{
			JobStatus: "Completed",
		}
	default:
		slog.Error("Job status is invalid")
		http.Error(w, "Job status is invalid", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error(err.Error())
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func CharacterEatPost(w http.ResponseWriter, r *http.Request) {
	slog.Info(fmt.Sprintf("[%s] %s", r.Method, r.URL.Path))

	// キャラクターを取得
	id := r.PathValue("id")
	mu.Lock()
	defer mu.Unlock()
	char, exists := characterStore[id]
	if !exists {
		slog.Error("Character not found")
		http.Error(w, "Character not found", http.StatusNotFound)
		return
	}

	// リクエストのバリデーション
	var req PostEatReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error(err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 今日食べたものを追加する
	char.Foods = append(char.Foods, req.Food)

	w.WriteHeader(http.StatusOK)
}
