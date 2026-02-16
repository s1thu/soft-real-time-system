package service

import (
	"context"
	"time"

	"s1thu/soft-real-time-system/backend/internal/model"
)

// EventProcessor handles the processing of events with deadline constraints
type EventProcessor struct {
	workDuration time.Duration
}

// NewEventProcessor creates a new event processor
func NewEventProcessor(workDuration time.Duration) *EventProcessor {
	return &EventProcessor{
		workDuration: workDuration,
	}
}

// Process processes an event and returns its status based on deadline compliance
func (p *EventProcessor) Process(event model.Event) string {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		event.DeadlineMs,
	)
	defer cancel()

	workDone := make(chan struct{})

	go func() {
		time.Sleep(p.workDuration)
		close(workDone)
	}()

	select {
	case <-workDone:
		return "on-time"
	case <-ctx.Done():
		return "late"
	}
}

// ProcessWithStatus processes an event and returns the event with status populated
func (p *EventProcessor) ProcessWithStatus(event model.Event) model.Event {
	event.Status = p.Process(event)
	return event
}
