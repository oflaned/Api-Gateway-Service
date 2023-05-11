package client

import (
	"Mehmat/config"
	"Mehmat/utils"
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, config config.StorageConfig) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	var err error
	dsn := fmt.Sprintf("postgressql://%s:%s@%s:%s/%s", config.Username, config.Passwd, config.Host, config.Port, config.Database)
	err = utils.DoWithAttmepts(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}

		return nil
	}, config.MaxAttempts, 5*time.Second)

	if err != nil {
		log.Fatal("Error while connect to db")
	}
	return pool, nil
}
