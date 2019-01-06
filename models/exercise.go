package models

type Exercise struct {
	ID         int64    `json:"id"`
	Name       string   `json:"name"`
	CategoryID int64    `json:"-"`
	Category   Category `json:"category"`
}

type Exercises []*Exercise
