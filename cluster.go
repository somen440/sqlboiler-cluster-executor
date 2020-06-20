package cluster

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Cluster struct {
	Writer *sql.DB
	Reader *sql.DB
	err    error
}

var (
	_ boil.ContextExecutor = &Cluster{}
	_ boil.ContextBeginner = &Cluster{}
)

func New(driver string, dsnList []string) *Cluster {
	cluster := &Cluster{}

	if len(dsnList) != 2 {
		cluster.err = fmt.Errorf("dsnList len expected 2 actual %d", len(dsnList))
		return cluster
	}

	writer, err := createDB(driver, dsnList[0])
	if err != nil {
		cluster.err = err
		return cluster
	}
	reader, err := createDB(driver, dsnList[1])
	if err != nil {
		cluster.err = err
		return cluster
	}

	cluster.Writer = writer
	cluster.Reader = reader

	return cluster
}

func (c *Cluster) Exec(query string, args ...interface{}) (sql.Result, error) {
	if c.HasError() {
		return nil, c.err
	}
	return nil, errors.New("dont use")
}

func (c *Cluster) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if c.HasError() {
		return nil, c.err
	}
	return c.Reader.Query(query, args...)
}

func (c *Cluster) QueryRow(query string, args ...interface{}) *sql.Row {
	if c.HasError() {
		return nil
	}
	return c.Reader.QueryRow(query, args...)
}

func (c *Cluster) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if c.HasError() {
		return nil, c.err
	}
	return nil, errors.New("dont use")
}

func (c *Cluster) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if c.HasError() {
		return nil, c.err
	}
	return c.Reader.QueryContext(ctx, query, args...)
}

func (c *Cluster) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if c.HasError() {
		return nil
	}
	return c.Reader.QueryRowContext(ctx, query, args...)
}

func (c *Cluster) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	if c.HasError() {
		return nil, c.err
	}
	return c.Writer.BeginTx(ctx, opts)
}

func (c *Cluster) HasError() bool {
	return c.err != nil
}

func createDB(driver, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	// db.SetMaxOpenConns
	// db.SetMaxIdleConns
	// db.SetConnMaxLifeTime
	return db, nil
}
