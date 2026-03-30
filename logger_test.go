package main

import (
	"context"
	"testing"
)

// nopLogger implements Logger for tests.
type nopLogger struct{}

func (nopLogger) Debug(ctx context.Context, msg string) {
	// intentionally empty
}
func (nopLogger) Info(ctx context.Context, msg string) {
	// intentionally empty
}
func (nopLogger) Warn(ctx context.Context, msg string) {
	// intentionally empty
}
func (nopLogger) Error(ctx context.Context, msg string) {
	// intentionally empty
}
func (nopLogger) Flush() {
	// intentionally empty
}

func TestNopLogger(t *testing.T) {
	var l Logger = nopLogger{}
	l.Debug(context.Background(), "test")
	l.Info(context.Background(), "test")
	l.Warn(context.Background(), "test")
	l.Error(context.Background(), "test")
	l.Flush()
}
