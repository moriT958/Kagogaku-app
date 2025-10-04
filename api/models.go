package api

import "time"

// キャラクター
type Character struct {
	Id         string
	Name       string
	Appearance []byte // Base64 Image
	Status     CharacterHealth
	SleepTime  time.Time
	WakeUpTime time.Time
	Foods      []string // １日に食べたもの
}

type CharacterHealth int

const (
	healthy = iota
	lackOfSleep
	overeating
)
