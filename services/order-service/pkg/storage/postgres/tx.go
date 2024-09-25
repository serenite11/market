package storage_postgres

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Tx struct {
	DB pgx.Tx
}

type TxReq func(tx Tx) error

type TxRunner interface {
	Begin(ctx context.Context) (pgx.Tx, error)
}

func ExecTx(ctx context.Context, runner TxRunner, _ *string, req TxReq) error {
	pgxTx, err := runner.Begin(ctx)
	if err != nil {
		return err
	}

	tx := Tx{
		DB: pgxTx,
	}

	defer tx.Rollback(context.TODO())

	err = req(tx)
	if err != nil {
		return err
	}
	err = tx.Commit(context.TODO())
	if err != nil {
		return err
	}

	return nil
}

func (p Tx) Stats() *pgxpool.Stat {
	return nil
}

func (p Tx) Connect(_ context.Context) error { return nil }

func (p Tx) Begin(ctx context.Context) (pgx.Tx, error) {
	return p.DB.Begin(ctx)
}

func (p Tx) Rollback(ctx context.Context) {
	_ = p.DB.Rollback(ctx)
}

func (p Tx) Commit(ctx context.Context) error {
	return p.DB.Commit(ctx)
}

func (p Tx) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	return p.DB.Query(ctx, query, args...)
}

func (p Tx) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	rows, err := p.DB.Query(ctx, query, args...)
	if err != nil {
		return err
	}
	return pgxscan.DefaultAPI.ScanOne(dest, rows)
}

func (p Tx) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	rows, err := p.DB.Query(ctx, query, args...)
	if err != nil {
		return err
	}
	return pgxscan.DefaultAPI.ScanAll(dest, rows)
}

func (p Tx) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	return p.DB.Exec(ctx, query, args...)
}

func (p Tx) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return p.DB.QueryRow(ctx, query, args...)
}

func (p Tx) Close(ctx context.Context) error {
	if p.DB.Conn() == nil {
		return nil
	}
	return p.DB.Conn().Close(ctx)
}
