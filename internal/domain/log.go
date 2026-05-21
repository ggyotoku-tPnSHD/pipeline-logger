package domain

import (
	"time" 
	"context"
)

type LogEvent struct {
	ID string
	Timestamp time.Time
	Level string
	Service string
	Message string
	RawData string
}

type Storage interface {
	SaveBatch(ctx context.Context, events []LogEvent) error
	Close() error
}