package storage

import (
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/arnokay/arnobot-shared/apperror"
	"github.com/arnokay/arnobot-shared/applog"
	"github.com/arnokay/arnobot-shared/db"
	"github.com/arnokay/arnobot-shared/service"
)

type Storager interface {
	Query(ctx context.Context) db.Querier
	Database(ctx context.Context) db.DBTX
	HandleErr(ctx context.Context, err error) error
}

type Storage struct {
	db    db.DBTX
	query *db.Queries

	logger applog.Logger
}

func NewStorage(database db.DBTX) *Storage {
	logger := applog.NewServiceLogger("storage")

	return &Storage{
		db:     database,
		query:  db.New(database),
		logger: logger,
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

func (s *Storage) HandleErr(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		s.logger.DebugContext(ctx, "resource not found")
		return apperror.ErrNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // unique_violation
			s.logger.DebugContext(ctx, "resource already exists", "err", pgErr)
			return apperror.ErrAlreadyExists
		case "23503": // foreign_key_violation
			s.logger.DebugContext(ctx, "resource foreign key violation", "err", pgErr)
			return apperror.ErrInvalidInput
		case "23514": // check_violation
			s.logger.DebugContext(ctx, "resource check violation", "err", pgErr)
			return apperror.ErrInvalidInput
		default:
			s.logger.ErrorContext(ctx, "unknown pgx error", "err", pgErr)
			return err
		}
	}

	if errors.Is(err, context.Canceled) ||
		errors.Is(err, context.DeadlineExceeded) ||
		strings.Contains(err.Error(), "connection refused") {
		s.logger.ErrorContext(ctx, "connection error", "err", err)
		return err
	}

	s.logger.ErrorContext(ctx, "unknown error", "err", err)
	return err
}
