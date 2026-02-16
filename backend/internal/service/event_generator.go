package service

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"

	"s1thu/soft-real-time-system/backend/internal/model"
)

// EventGenerator handles the creation and distribution of events
type EventGenerator struct {
	interval time.Duration
	deadline time.Duration
	out      chan model.Event
	ctx      context.Context
	cancel   context.CancelFunc
	wg       sync.WaitGroup
}

// NewEventGenerator creates a new event generator with the specified interval and deadline
func NewEventGenerator(interval, deadline time.Duration, bufferSize int) *EventGenerator {
	ctx, cancel := context.WithCancel(context.Background())
	return &EventGenerator{
		interval: interval,
		deadline: deadline,
		out:      make(chan model.Event, bufferSize),
		ctx:      ctx,
		cancel:   cancel,
	}
}

// Start begins generating events at the configured interval
func (g *EventGenerator) Start() {
	g.wg.Add(1)
	go g.run()
}

// Stop gracefully stops the event generator
func (g *EventGenerator) Stop() {
	g.cancel()
	g.wg.Wait()
	close(g.out)
}

// Events returns the channel for receiving generated events
func (g *EventGenerator) Events() <-chan model.Event {
	return g.out
}

func (g *EventGenerator) run() {
	defer g.wg.Done()

	ticker := time.NewTicker(g.interval)
	defer ticker.Stop()

	for {
		select {
		case <-g.ctx.Done():
			return
		case <-ticker.C:
			event := model.Event{
				ID:         uuid.New().String(),
				CreatedAt:  time.Now(),
				DeadlineMs: g.deadline,
			}

			select {
			case g.out <- event:
			default:
				// Drop event when load peaks on the goroutine
			}
		}
	}
}
