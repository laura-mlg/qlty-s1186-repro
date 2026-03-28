package main

import (
	"context"
	"fmt"
)

func main() {
	var l Logger = nopLogger{}
	l.Debug(context.Background(), "starting")
	l.Info(context.Background(), "running")

	m := NewMetricsNoop()
	m.RecordSchedulerRun("startup")

	fmt.Println("ok")
}
