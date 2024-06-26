package factory

import storage_postgres "order-service/pkg/storage/postgres"

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
