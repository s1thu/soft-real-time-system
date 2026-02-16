package model

import "time"

// Event represents a soft real-time event with deadline constraints
type Event struct {
	ID         string        `json:"id"`
	CreatedAt  time.Time     `json:"created_at"`
	DeadlineMs time.Duration `json:"deadline_ms"`
	Status     string        `json:"status,omitempty"`
}
