package storage_postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"

	"github.com/jackc/pgx/v4"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

const (
	_defaultConnectionAttempts = 10
	_defaultConnectionTimeout  = time.Second
	_maxConnections            = int32(200)
	_minConnections            = int32(20)
	_maxConnectionLifeTime     = time.Second * 300
	_maxIdleLifeTime           = time.Second * 15
)

type Postgres interface {
	Connect(ctx context.Context) error
	Stats() *pgxpool.Stat
	Query(ctx context.Context, query string, args ...any) (pgx.Rows, error)
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Close(ctx context.Context) error
	TxRunner
}

type Pool struct {
	db  *pgxpool.Pool
	cfg *Config
}

func New(cfg *Config) *Pool {
	return &Pool{cfg: cfg}
}

func (p *Pool) Connect(ctx context.Context) error {
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		p.cfg.Host,
		p.cfg.Port,
		p.cfg.User,
		p.cfg.Password,
		p.cfg.Database,
		p.cfg.SslMode,
	)
	connectionUrl += fmt.Sprintf(
		" pool_max_conns=%d pool_min_conns=%d pool_max_conn_lifetime=%v pool_max_conn_idle_time=%v",
		_maxConnections,
		_minConnections,
		_maxConnectionLifeTime,
		_maxIdleLifeTime,
	)

	connectionAttempts := _defaultConnectionAttempts
	var result *pgxpool.Pool
	var err error
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	for connectionAttempts > 0 {
		result, err = pgxpool.Connect(ctx, connectionUrl)
		if err == nil {
			break
		}

		log.Printf(
			"ATTEMPT %d TO CONNECT TO POSTGRES BY URL %s FAILED: %s\n",
			connectionAttempts,
			connectionUrl,
			err.Error(),
		)

		connectionAttempts--

		time.Sleep(_defaultConnectionTimeout)
	}
	if result == nil {
		log.Printf("POSTGRES CONNECTION(%s) ERROR: %s\n", connectionUrl, err.Error())
		return err
	}
	p.db = result

	migrationUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		p.cfg.User,
		p.cfg.Password,
		p.cfg.Host,
		p.cfg.Port,
		p.cfg.Database,
		p.cfg.SslMode,
	)

	migration, err := migrate.New(p.cfg.MigrationsPath, migrationUrl)
	if err != nil {
		log.Printf("Migrating database failed: %s", err.Error())
		return err
	}

	if err = migration.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			version, dirty, err := migration.Version()
			if err != nil {
				return err
			}
			log.Printf("latest migrations already up, version : %d | dirty : %t", version, dirty)
			return nil
		}
		log.Printf("Migration failed: %s", err.Error())
		return err
	}

	return nil
}

func (p *Pool) Stats() *pgxpool.Stat {
	return p.db.Stat()
}

func (p *Pool) Begin(ctx context.Context) (pgx.Tx, error) {
	return p.db.Begin(ctx)
}

func (p *Pool) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	return p.db.Query(ctx, query, args...)
}

func (p *Pool) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return err
	}
	return pgxscan.DefaultAPI.ScanOne(dest, rows)
}

func (p *Pool) Select(
	ctx context.Context,
	dest interface{},
	query string,
	args ...interface{},
) error {
	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return err
	}
	return pgxscan.DefaultAPI.ScanAll(dest, rows)
}

func (p *Pool) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	return p.db.Exec(ctx, query, args...)
}

func (p *Pool) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return p.db.QueryRow(ctx, query, args...)
}

func (p *Pool) Close(_ context.Context) error {
	p.db.Close()
	return nil
}
