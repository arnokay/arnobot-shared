package storage

import (
	"context"

	"arnobot-shared/db"
	"arnobot-shared/service"
)

type Storager interface {
	Query(ctx context.Context) db.Querier
	Database(ctx context.Context) db.DBTX
}

type Storage struct {
	db    db.DBTX
	query *db.Queries
}

func NewStorage(database db.DBTX) *Storage {
	return &Storage{
		db:    database,
		query: db.New(database),
	}
}

func (s *Storage) Query(ctx context.Context) db.Querier {
	if tx := service.ExtractTx(ctx); tx != nil {
		return s.query.WithTx(tx)
	}
	return s.query
}

func (s *Storage) Database(ctx context.Context) db.DBTX {
	if tx := service.ExtractTx(ctx); tx != nil {
		return tx
	}

	return s.db
}
