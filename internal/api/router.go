package api

import (
	"net/http"
	"telem.kmani/internal/api/handlers"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Register health check
	mux.HandleFunc("/api/health/{$}", handlers.HealthCheck)

	// Register Telemetry endpoints
	mux.HandleFunc("GET /api/telemetry/{$}", handlers.TelemetrySearch)
	return mux
}
