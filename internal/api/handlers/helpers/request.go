package helpers

import (
	"errors"
	"time"
)

//
// ────────────────────────────────
// Validation errors & constants
// ────────────────────────────────
//

var (
	ErrInvalidTimeRange = errors.New("'to' must be after 'from'")
	ErrInvalidDeviceID  = errors.New("deviceId is required")
	ErrInvalidMetric    = errors.New("metric is required")
	ErrInvalidSeverity  = errors.New("invalid severity level")
)

const (
	SeverityInfo     = "INFO"
	SeverityWarning  = "WARNING"
	SeverityCritical = "CRITICAL"
)

//
// ────────────────────────────────
// Shared reusable helpers
// ────────────────────────────────
//

// TimeRange: reusable for any request that needs a time window.
type TimeRange struct {
	From string `form:"from"` // ISO8601
	To   string `form:"to"`   // ISO8601
}

func (tr *TimeRange) Validate() error {
	if tr.From == "" || tr.To == "" {
		return nil // open-ended allowed
	}
	from, err := time.Parse(time.RFC3339, tr.From)
	if err != nil {
		return errors.New("invalid 'from' time format; must be RFC3339")
	}
	to, err := time.Parse(time.RFC3339, tr.To)
	if err != nil {
		return errors.New("invalid 'to' time format; must be RFC3339")
	}
	if to.Before(from) {
		return ErrInvalidTimeRange
	}
	return nil
}

// PaginationParams: optional for any endpoint that supports paging.
type PaginationParams struct {
	From int `form:"page_from"` // start position
	Size int `form:"page_size"` // page size
}

func (pp *PaginationParams) Validate() error {
	if pp.From < 0 {
		return errors.New("'page_from' must be non-negative")
	}
	if pp.Size < 1 {
		return errors.New("'page_size' must be positive")
	}
	return nil
}

//
// ────────────────────────────────
// Telemetry domain
// ────────────────────────────────
//

// GET /telemetry
type TelemetryQuery struct {
	DeviceID string `form:"deviceId"`
	TimeRange
	PaginationParams
}

func (q *TelemetryQuery) Validate() error {
	if q.DeviceID == "" {
		return ErrInvalidDeviceID
	}
	if err := q.TimeRange.Validate(); err != nil {
		return err
	}
	// If both pagination parameters are 0, consider pagination optional
	if q.PaginationParams.From == 0 && q.PaginationParams.Size == 0 {
		// Default page size for backward compatibility
		q.PaginationParams.Size = 10
		return nil
	}
	// Otherwise validate pagination parameters
	return q.PaginationParams.Validate()
}

// GET /telemetry/aggregate
type TelemetryAggregateQuery struct {
	DeviceID string `form:"deviceId"`
	Metric   string `form:"metric"`
	TimeRange
}

func (q *TelemetryAggregateQuery) Validate() error {
	if q.DeviceID == "" {
		return ErrInvalidDeviceID
	}
	if q.Metric == "" {
		return ErrInvalidMetric
	}
	return q.TimeRange.Validate()
}

// POST /telemetry/search
type TelemetrySearchRequest struct {
	DeviceID string   `json:"deviceId"`
	From     string   `json:"from"`
	To       string   `json:"to"`
	Metrics  []string `json:"metrics"`
}

func (r *TelemetrySearchRequest) Validate() error {
	if r.DeviceID == "" {
		return ErrInvalidDeviceID
	}
	return (&TimeRange{From: r.From, To: r.To}).Validate()
}

//
// ────────────────────────────────
// Events domain
// ────────────────────────────────
//

// GET /events
type EventsQuery struct {
	DeviceID string `form:"deviceId"`
	Severity string `form:"severity"` // INFO, WARNING, CRITICAL
	TimeRange
	PaginationParams
}

func (q *EventsQuery) Validate() error {
	if q.Severity != "" {
		valid := map[string]bool{
			SeverityInfo:     true,
			SeverityWarning:  true,
			SeverityCritical: true,
		}
		if !valid[q.Severity] {
			return ErrInvalidSeverity
		}
	}
	return q.TimeRange.Validate()
}

//
// ────────────────────────────────
// Device Config domain
// ────────────────────────────────
//

// GET /device-config/:deviceId (optional filters)
type DeviceConfigQuery struct {
	DeviceID string `form:"deviceId"`
	TimeRange
}

func (q *DeviceConfigQuery) Validate() error {
	if q.DeviceID == "" {
		return ErrInvalidDeviceID
	}
	return q.TimeRange.Validate()
}

// PUT /device-config/:deviceId (or POST)
type DeviceConfigRequest struct {
	DeviceName string   `json:"deviceName"`
	Model      string   `json:"model"`
	Firmware   string   `json:"firmware"`
	Tags       []string `json:"tags"`
}

func (r *DeviceConfigRequest) Validate() error {
	if r.Model == "" {
		return errors.New("model is required")
	}
	return nil
}
