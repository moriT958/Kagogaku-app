package api

import "time"

// POST /character/new
type PostNewCharacterReq struct {
	Id         string `json:"id"` // UUID in local storage
	Name       string `json:"name"`
	Appearance []byte `json:"appearance"`
}

// GET /charactor/{id}
// PATCH /charactor/{id}/sleep
type CharacterResp struct {
	Name       string    `json:"name"`
	Status     string    `json:"status"` // health(default), lackOfSleep
	SleepTime  time.Time `json:"sleepTime"`
	WakeUpTime time.Time `json:"wakeUpTime"`
	Foods      []string  `json:"foods"` // foods per date
}

// PATCH character/{id}/wake-up
type PostWakeUpResp struct {
	JobId int `json:"jobId"`
}

// GET /train-status/{jobId}
type GetTrainStatusResp struct {
	JobStatus string `json:"jobStatus"` // completed or processing or failed
}

// POST /character/{id}/eat
type PostEatReq struct {
	Food string `json:"food"`
}
