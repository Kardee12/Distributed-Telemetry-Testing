package handlers

import (
	"net/http"
	"strconv"
	"telem.kmani/internal/api/handlers/helpers"
	"telem.kmani/internal/storage"
)

// Endpoints:
//   GET /api/telemetry
//     - Search telemetry docs by deviceId and time range.
//   GET /api/telemetry/aggregate
//     - Return avg/min/max for a metric over time.
//   POST /api/telemetry/search
//     - Run advanced search with JSON body.
//
// Uses:
//   - TelemetryQuery, TelemetryAggregateQuery, TelemetrySearchRequest
//   - storage.SearchTelemetry(), storage.AggregateTelemetry()
//   - api.WriteJSON, api.WriteError

func TelemetrySearch(w http.ResponseWriter, r *http.Request) {
	deviceID := r.URL.Query().Get("deviceId")
	if deviceID == "" {
		helpers.WriteError(w, http.StatusBadRequest, "deviceId is required")
		return
	}

	tr := helpers.TimeRange{
		From: r.URL.Query().Get("from"),
		To:   r.URL.Query().Get("to"),
	}

	// Extract pagination parameters
	var from, size int

	if fromParam := r.URL.Query().Get("page_from"); fromParam != "" {
		parsedFrom, err := strconv.Atoi(fromParam)
		if err != nil {
			helpers.WriteError(w, http.StatusBadRequest, "Invalid 'page_from' parameter")
			return
		}
		from = parsedFrom
	}

	if sizeParam := r.URL.Query().Get("page_size"); sizeParam != "" {
		parsedSize, err := strconv.Atoi(sizeParam)
		if err != nil {
			helpers.WriteError(w, http.StatusBadRequest, "Invalid 'page_size' parameter")
			return
		}
		size = parsedSize
	}

	query := helpers.TelemetryQuery{
		DeviceID:  deviceID,
		TimeRange: tr,
		PaginationParams: helpers.PaginationParams{
			From: from,
			Size: size,
		},
	}

	if err := query.Validate(); err != nil {
		helpers.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	results, total, err := storage.SearchTelemetry(r.Context(), query)
	if err != nil {
		helpers.WriteError(w, http.StatusInternalServerError, "Search failed")
		return
	}

	// Create response
	response := helpers.APIResponse{
		Data:  results,
		Count: len(results),
	}

	// Add pagination metadata only if pagination parameters were provided
	if size > 0 {
		meta := map[string]interface{}{
			"page":  from/size + 1,
			"size":  size,
			"total": total,
		}
		response.Meta = meta
	} else {
		// If no pagination, just include total count in metadata
		response.Meta = map[string]interface{}{
			"total": total,
		}
	}

	helpers.WriteJSON(w, http.StatusOK, response)
}
