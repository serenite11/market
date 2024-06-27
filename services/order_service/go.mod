module order-service

go 1.22.4

replace github.com/serenite11/market/proto => ../../proto

require (
	github.com/georgysavva/scany v1.2.2
	github.com/golang-migrate/migrate/v4 v4.17.1
	github.com/jackc/pgconn v1.14.3
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jackc/pgx/v4 v4.18.2
	github.com/lib/pq v1.10.9
	github.com/mattn/go-colorable v0.1.13
	github.com/serenite11/market/proto v0.0.0
	go.uber.org/config v1.4.0
	go.uber.org/fx v1.22.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.14.0 // indirect
	github.com/jackc/puddle v1.3.0 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/dig v1.17.1 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/lint v0.0.0-20190930215403-16217165b5de // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.10.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/grpc v1.64.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
