package api

import (
	"errors"
	"fmt"
	"log/slog"
	"math/rand"
	"sync"
	"time"
)

var (
	mu sync.RWMutex

	// キャラクターを保存用のインメモリ KVS
	characterStore = make(map[string]*Character)

	// キャラクター画像変換ジョブキュー
	jobQueue   = make(map[string]*Job)
	jobCounter = 0

	// ジョブワーカー
	jobWorker = make(chan *Job, 100)
)

type Job struct {
	Id     string
	Status JobStatus
}

type JobStatus int

const (
	completed = iota
	processing
	failed
)

func StartJobWorker() {
	go func() {
		slog.Info("Job worker started")
		for job := range jobWorker {
			processJob(job)
		}
	}()
}

//	ジョブの内容:
//	  キャラの画像を OpenAI で編集
//	  昨日の食べたもの、睡眠時間、元の見た目の画像を元に変換

// processJob は実際のジョブ処理を行います（OpenAI API の代わりに time.Sleep でシミュレート）
func processJob(job *Job) {
	slog.Info(fmt.Sprintf("Processing job %s", job.Id))

	// 5〜10秒の処理時間をシミュレート
	processingTime := 5 + rand.Intn(6) // 5〜10秒
	time.Sleep(time.Duration(processingTime) * time.Second)

	// 90% の確率で成功、10% の確率で失敗
	mu.Lock()
	defer mu.Unlock()

	if rand.Float32() < 0.9 {
		job.Status = completed
		slog.Info(fmt.Sprintf("Job %s completed successfully", job.Id))
	} else {
		job.Status = failed
		slog.Error(fmt.Sprintf("Job %s failed", job.Id))
	}
}

func EnqueueTrainJob(base64Image []byte, sleepDuration float32, foods []string) (int, error) {
	mu.Lock()
	jobCounter++
	jobId := jobCounter

	job := &Job{
		Id:     fmt.Sprintf("%d", jobId),
		Status: processing,
	}
	jobQueue[fmt.Sprintf("%d", jobId)] = job
	mu.Unlock()

	jobWorker <- job

	return jobId, nil
}

func DequeueTrainJob(id string) (*Job, error) {
	mu.RLock()
	job, exists := jobQueue[id]
	mu.RUnlock()
	if !exists {
		return nil, errors.New("job not found")
	}
	return job, nil
}
