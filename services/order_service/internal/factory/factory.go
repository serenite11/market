package factory

import (
	"context"
	storage_postgres "order-service/pkg/storage/postgres"
)

type Factory struct {
	postgres storage_postgres.Postgres
}

func New(
	postgres storage_postgres.Postgres,
) *Factory {
	return &Factory{
		postgres: postgres,
	}
}

type CallbackFunc = func(f *Factory) error

func (f *Factory) ExecuteTx(ctx context.Context, cb CallbackFunc) error {
	tx, err := f.postgres.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	store := storage_postgres.Tx{
		DB: tx,
	}
	factoryExecute := Factory{
		postgres: store,
	}

	if err = cb(&factoryExecute); err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}
