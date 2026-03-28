package main

import (
	"context"
	"testing"
)

// nopLogger implements Logger for tests.
type nopLogger struct{}

func (nopLogger) Debug(ctx context.Context, msg string) {}
func (nopLogger) Info(ctx context.Context, msg string)  {}
func (nopLogger) Warn(ctx context.Context, msg string)  {}
func (nopLogger) Error(ctx context.Context, msg string) {}
func (nopLogger) Flush()                                {}

func TestNopLogger(t *testing.T) {
	var l Logger = nopLogger{}
	l.Debug(context.Background(), "test")
	l.Info(context.Background(), "test")
	l.Warn(context.Background(), "test")
	l.Error(context.Background(), "test")
	l.Flush()
}
