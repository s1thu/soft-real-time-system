package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"s1thu/soft-real-time-system/backend/internal/model"
)

// WebSocketHandler handles WebSocket connections for real-time event streaming
type WebSocketHandler struct {
	upgrader websocket.Upgrader
	eventsCh <-chan model.Event
}

// NewWebSocketHandler creates a new WebSocket handler
func NewWebSocketHandler(eventsCh <-chan model.Event) *WebSocketHandler {
	return &WebSocketHandler{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true // Allow all origins in development
			},
		},
		eventsCh: eventsCh,
	}
}

// Handle upgrades HTTP connection to WebSocket and streams events
func (h *WebSocketHandler) Handle(c *gin.Context) {
	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	log.Println("WebSocket client connected")

	for event := range h.eventsCh {
		if err := conn.WriteJSON(event); err != nil {
			log.Printf("Failed to write message: %v", err)
			break
		}
	}

	log.Println("WebSocket client disconnected")
}
