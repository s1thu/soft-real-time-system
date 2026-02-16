package router

import (
	"github.com/gin-gonic/gin"

	"s1thu/soft-real-time-system/backend/internal/handler"
	"s1thu/soft-real-time-system/backend/internal/middleware"
)

// Config holds the router configuration
type Config struct {
	WebSocketHandler *handler.WebSocketHandler
	HealthHandler    *handler.HealthHandler
}

// New creates and configures the Gin router
func New(cfg *Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// Apply global middleware
	r.Use(middleware.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())

	// Register routes
	setupRoutes(r, cfg)

	return r
}

func setupRoutes(r *gin.Engine, cfg *Config) {
	// Health check endpoint
	r.GET("/health", cfg.HealthHandler.Check)

	// API v1 group
	v1 := r.Group("/api/v1")
	{
		// WebSocket endpoint for real-time events
		v1.GET("/ws", cfg.WebSocketHandler.Handle)
	}
}
