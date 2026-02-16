# Soft Real-Time System

A soft real-time event processing system with a Go backend and React dashboard.

```
+-------------+     +------------------+     +----------------+
|   Event     | --> | Deadline-Aware   | --> |   WebSocket    |
|  Generator  |     |   Go Processor   |     |    Server      |
+-------------+     +------------------+     +----------------+
                                                     |
                                                     v
                                             +----------------+
                                             | React Dashboard|
                                             +----------------+
```

## Architecture

- **Backend**: Go with Gin framework - generates events at 50ms intervals with 100ms deadlines
- **Frontend**: React + TypeScript + Vite - real-time dashboard displaying event status
- **Communication**: WebSocket for real-time event streaming

## Quick Start with Docker (Recommended)

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Run with Docker Compose

```bash
# Build and start all services
docker compose up --build

# Or run in detached mode
docker compose up --build -d
```

**Access the application:**

- Frontend Dashboard: http://localhost:3000
- Backend Health Check: http://localhost:8080/health
- WebSocket Endpoint: ws://localhost:8080/api/v1/ws

### Stop Services

```bash
# Stop and remove containers
docker compose down

# Stop and remove containers with volumes
docker compose down -v
```

## Manual Setup (Development)

### Prerequisites

- [Go 1.24+](https://golang.org/dl/)
- [Node.js 20+](https://nodejs.org/)
- [npm](https://www.npmjs.com/)

### Backend

```bash
cd backend

# Install dependencies
go mod tidy

# Run the server
go run ./cmd/server
```

The server starts on port `8080` by default. Configure via `.env` file:

```env
SERVER_PORT=8080
```

### Frontend

```bash
cd frontend/dashboard-real-time

# Install dependencies
npm install

# Run development server
npm run dev
```

The frontend starts on http://localhost:5173

## Project Structure

```
soft-real-time-system/
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go           # Application entry point
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go         # Configuration management
│   │   ├── handler/
│   │   │   ├── health.go         # Health check handler
│   │   │   └── websocket.go      # WebSocket handler
│   │   ├── middleware/
│   │   │   ├── cors.go           # CORS middleware
│   │   │   ├── logger.go         # Request logging
│   │   │   └── recovery.go       # Panic recovery
│   │   ├── model/
│   │   │   └── event.go          # Event model
│   │   ├── router/
│   │   │   └── router.go         # Gin router setup
│   │   └── service/
│   │       ├── event_generator.go # Event generation service
│   │       └── event_processor.go # Event processing service
│   ├── .env                      # Environment variables
│   ├── Dockerfile                # Backend Docker image
│   ├── go.mod
│   └── go.sum
├── frontend/
│   └── dashboard-real-time/
│       ├── src/
│       │   ├── components/
│       │   │   ├── EventsTable.tsx
│       │   │   └── StatsBar.tsx
│       │   ├── hooks/
│       │   │   └── useWebSocket.ts
│       │   ├── types/
│       │   │   └── types.ts
│       │   ├── App.tsx
│       │   ├── App.css
│       │   └── main.tsx
│       ├── Dockerfile            # Frontend Docker image
│       ├── nginx.conf            # Nginx configuration
│       └── package.json
├── docker-compose.yml            # Docker Compose configuration
└── README.md
```

## API Endpoints

| Endpoint     | Method    | Description            |
| ------------ | --------- | ---------------------- |
| `/health`    | GET       | Health check           |
| `/api/v1/ws` | WebSocket | Real-time event stream |

## Event Model

```json
{
  "id": "uuid-string",
  "created_at": "2026-02-16T10:00:00Z",
  "deadline_ms": 100000000,
  "status": "on-time | late"
}
```

## Configuration

### Backend Environment Variables

| Variable      | Default | Description      |
| ------------- | ------- | ---------------- |
| `SERVER_PORT` | `8080`  | HTTP server port |

## Reference

This project was built following the tutorial:

- [Real-Time Systems for Web Developers – From Theory to a Live Go + React App](https://www.freecodecamp.org/news/real-time-systems-for-web-developers-from-theory-to-a-live-go-react-app/)

## License

MIT
