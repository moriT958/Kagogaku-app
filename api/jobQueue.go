package api

import (
	"errors"
	"fmt"
	"log/slog"
	"sync"
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
	Data   JobData
}

type JobData struct {
	CharacterId   string
	SleepDuration float32
	Foods         []string
}

type JobStatus int

const (
	completed JobStatus = iota
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

// processJob は実際のジョブ処理を行います
func processJob(job *Job) {
	slog.Info(fmt.Sprintf("Processing job %s", job.Id))

	// キャラクターの現在の Appearance を取得
	mu.RLock()
	char, exists := characterStore[job.Data.CharacterId]
	if !exists {
		mu.RUnlock()
		slog.Error("Character not found")
		job.Status = failed
		return
	}
	currentAppearance := char.Appearance
	mu.RUnlock()

	prompt := buildPrompt(*job)
	editedImgPath, err := updateCharacterImage(currentAppearance, prompt)
	if err != nil {
		job.Status = failed
		slog.Error(fmt.Sprintf("Job %s failed: %v", job.Id, err))
		return
	}

	mu.Lock()
	char, exists = characterStore[job.Data.CharacterId]
	if !exists {
		mu.Unlock()
		slog.Error("Character not found")
		job.Status = failed
		return
	}
	char.Appearance = editedImgPath
	job.Status = completed
	mu.Unlock()
}

func EnqueueTrainJob(data JobData) (int, error) {
	mu.Lock()
	jobCounter++
	jobId := jobCounter

	job := &Job{
		Id:     fmt.Sprintf("%d", jobId),
		Status: processing,
		Data:   data,
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
