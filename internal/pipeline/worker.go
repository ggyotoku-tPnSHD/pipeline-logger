package worker

import (
	"domain.log-engine/internal/domain"
	"context"
)

type Worker struct {}

func NewWorker() *Worker {
	return &Worker{}
}

func (w *Worker) Start(ctx context.Context, ch <-chan string, rc chan<- domain.LogEvent) {}