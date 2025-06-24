package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5"

	"github.com/arnokay/arnobot-shared/apperror"
	"github.com/arnokay/arnobot-shared/applog"
)

type ITransactionService interface {
	Begin(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type transactionCtx string

const txKey transactionCtx = "tx-key"

func ExtractTx(ctx context.Context) pgx.Tx {
	tx, ok := ctx.Value(txKey).(pgx.Tx)
	if !ok {
		return nil
	}

	return tx
}

type pgxTransactional interface {
	Begin(ctx context.Context) (pgx.Tx, error)
}

type PgxTransactionService struct {
	db pgxTransactional

	logger *slog.Logger
}

func NewPgxTransactionService(db pgxTransactional) *PgxTransactionService {
	logger := applog.NewServiceLogger("transaction-service")

	return &PgxTransactionService{
		db: db,

		logger: logger,
	}
}

func (s *PgxTransactionService) Begin(ctx context.Context) (context.Context, error) {
	if tx, ok := ctx.Value(txKey).(pgx.Tx); ok && tx != nil {
		nestedTx, err := tx.Begin(ctx)
		if err != nil {
			s.logger.ErrorContext(ctx, "cannot begin nested transaction", "err", err)
			return nil, apperror.ErrInternal
		}
		txCtx := context.WithValue(ctx, txKey, nestedTx)
		s.logger.DebugContext(ctx, "began nested transaction")
		return txCtx, nil
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		s.logger.ErrorContext(ctx, "cannot begin transaction", "err", err)
		return nil, apperror.ErrInternal
	}
	txCtx := context.WithValue(ctx, txKey, tx)

	s.logger.DebugContext(ctx, "begined transaction")

	return txCtx, nil
}

func (s *PgxTransactionService) Commit(ctx context.Context) error {
	tx, ok := ctx.Value(txKey).(pgx.Tx)
	if !ok || tx == nil {
		s.logger.ErrorContext(ctx, "transaction is nil, cannot commit")
		return apperror.ErrInternal
	}

	err := tx.Commit(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrTxClosed) {
			return nil
		}
		if errors.Is(err, pgx.ErrTxCommitRollback) {
			s.logger.WarnContext(ctx, "cannot commit, transaction was rolled back", "err", err)
			return apperror.ErrInternal
		}

		s.logger.ErrorContext(ctx, "cannot commit", "err", err)
		return apperror.ErrInternal
	}

	s.logger.DebugContext(ctx, "commited transaction")

	return nil
}

func (s *PgxTransactionService) Rollback(ctx context.Context) error {
	tx, ok := ctx.Value(txKey).(pgx.Tx)
	if !ok || tx == nil {
		s.logger.ErrorContext(ctx, "transaction is nil, cannot rollback")
		return apperror.ErrInternal
	}

	err := tx.Rollback(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrTxClosed) {
			return nil
		}

		s.logger.ErrorContext(ctx, "cannot rollback", "err", err)
		return apperror.ErrInternal
	}

	s.logger.DebugContext(ctx, "rolled back transaction")

	return nil
}
