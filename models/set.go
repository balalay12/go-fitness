package models

import "time"

type Set struct {
	ID         int64     `json:"id"`
	Date       time.Time `json:"date"`
	ExerciseID int64     `json:"exercise_id"`
}

type SetsOrdered struct {
	Date time.Time `json:"date"`
	Data []*Set    `json:"data"`
}

type Sets []*Set

