package storage

import (
	"context"
	"telem.kmani/internal/api/handlers/helpers"
	"telem.kmani/internal/models"
)

// SearchTelemetry should query OpenSearch for telemetry docs by deviceId and time range.
func SearchTelemetry(ctx context.Context, query helpers.TelemetryQuery) ([]models.Telemetry, int, error) {
	return nil, 0, nil
}
