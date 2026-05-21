package storage

import (
	"context"
	"log-engine/internal/domain"
	"database/sql"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{
		db:db
	}
}

func (p *PostgresStorage) SaveBatch(ctx context.Context, events []domain.LogEvent) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO domain.LogEvent(ID, Timestamp, Level, Service, Message, RawData) VALUES(?, ?, ?, ?, ?, ?)")

	if err != nil {
		return err	
	}

	for _ , event := range events {
		_, err := stmt.ExecContext(ctx, event.ID, event.Timestamp, event.Level, event.Service, event.Message, event.RawData)

		if err != nil {
			return err	
		}

	} 	
		
	err = tx.Commit()

	if err != nil {
		return err
	}

	return err
	
}

func (p *PostgresStorage) Close() error {
	return nil
}

