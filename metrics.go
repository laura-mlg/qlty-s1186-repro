package main

// Metrics defines the contract for reporting scheduler metrics.
type Metrics interface {
	RecordElectedAsPrimary(elected bool)
	RecordSchedulerRun(eventLabel string)
	RecordSchedulerRunDuration(duration float64)
	RecordSchedulerError(eventType string)
}

// NewMetricsNoop returns a no-op Metrics implementation.
func NewMetricsNoop() Metrics {
	return metricsNoop{}
}

type metricsNoop struct{}

func (m metricsNoop) RecordElectedAsPrimary(elected bool)         {}
func (m metricsNoop) RecordSchedulerRun(eventLabel string)        {}
func (m metricsNoop) RecordSchedulerRunDuration(duration float64) {}
func (m metricsNoop) RecordSchedulerError(eventType string)       {}
