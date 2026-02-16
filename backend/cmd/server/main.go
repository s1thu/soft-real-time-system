package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"s1thu/soft-real-time-system/backend/internal/config"
	"s1thu/soft-real-time-system/backend/internal/handler"
	"s1thu/soft-real-time-system/backend/internal/model"
	"s1thu/soft-real-time-system/backend/internal/router"
	"s1thu/soft-real-time-system/backend/internal/service"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize services
	generator := service.NewEventGenerator(
		cfg.Event.Interval,
		cfg.Event.Deadline,
		cfg.Event.BufferSize,
	)
	processor := service.NewEventProcessor(cfg.Processor.WorkDuration)

	// Create processed events channel
	processedEvents := make(chan model.Event, cfg.Event.BufferSize)

	// Start event generator
	generator.Start()

	// Start event processing pipeline
	go func() {
		for event := range generator.Events() {
			processed := processor.ProcessWithStatus(event)
			select {
			case processedEvents <- processed:
			default:
				// Drop if channel is full
			}
		}
		close(processedEvents)
	}()

	// Initialize handlers
	wsHandler := handler.NewWebSocketHandler(processedEvents)
	healthHandler := handler.NewHealthHandler()

	// Setup router
	r := router.New(&router.Config{
		WebSocketHandler: wsHandler,
		HealthHandler:    healthHandler,
	})

	// Create HTTP server
	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Stop event generator
	generator.Stop()

	// Shutdown HTTP server
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped")
}
