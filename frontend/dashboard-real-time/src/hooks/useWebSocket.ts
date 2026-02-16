import { useEffect, useRef, useState, useCallback } from "react";
import type { RealTimeEvent } from "../types/types";

// Use relative WebSocket URL for production (proxied through nginx)
// or localhost for development
const getWsUrl = () => {
  if (import.meta.env.PROD) {
    const protocol = window.location.protocol === "https:" ? "wss:" : "ws:";
    return `${protocol}//${window.location.host}/api/v1/ws`;
  }
  return "ws://localhost:8080/api/v1/ws";
};

const MAX_EVENTS = 100;

export function useWebSocket() {
  const [events, setEvents] = useState<RealTimeEvent[]>([]);
  const [isConnected, setIsConnected] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const socketRef = useRef<WebSocket | null>(null);
  const reconnectTimeoutRef = useRef<ReturnType<typeof setTimeout> | undefined>(
    undefined,
  );

  const connect = useCallback(() => {
    try {
      const socket = new WebSocket(getWsUrl());

      socket.onopen = () => {
        setIsConnected(true);
        setError(null);
        console.log("WebSocket connected");
      };

      socket.onmessage = (event) => {
        try {
          const data: RealTimeEvent = JSON.parse(event.data);
          // Add processed_at timestamp when we receive the event
          data.processed_at = new Date().toISOString();

          setEvents((prev) => {
            const newEvents = [data, ...prev];
            // Keep only the last MAX_EVENTS to prevent memory issues
            return newEvents.slice(0, MAX_EVENTS);
          });
        } catch (err) {
          console.error("Failed to parse message:", err);
        }
      };

      socket.onclose = () => {
        setIsConnected(false);
        console.log("WebSocket disconnected");
        // Attempt to reconnect after 3 seconds
        reconnectTimeoutRef.current = setTimeout(connect, 3000);
      };

      socket.onerror = (err) => {
        setError("WebSocket connection error");
        console.error("WebSocket error:", err);
      };

      socketRef.current = socket;
    } catch (err) {
      setError("Failed to create WebSocket connection");
      console.error("Failed to connect:", err);
    }
  }, []);

  useEffect(() => {
    connect();

    return () => {
      if (reconnectTimeoutRef.current) {
        clearTimeout(reconnectTimeoutRef.current);
      }
      if (socketRef.current) {
        socketRef.current.close();
      }
    };
  }, [connect]);

  const clearEvents = useCallback(() => {
    setEvents([]);
  }, []);

  return {
    events,
    isConnected,
    error,
    clearEvents,
  };
}
