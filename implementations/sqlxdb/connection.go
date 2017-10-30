// Package sqlxdb contains an sqlx implementation of the sqldb package
package sqlxdb

import (
	"database/sql"

	sqldb "github.com/Nivl/go-sqldb"
	"github.com/jmoiron/sqlx"
)

var _ sqldb.Connection = (*Connection)(nil)

// New returns a new SQLX connection
func New(dsn string) (*Connection, error) {
	con, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Unsafe returns a version of DB which will silently succeed to scan when
	// columns in the SQL result have no fields in the destination struct.
	con = con.Unsafe()

	return &Connection{
		Queryable: NewQueryable(con),
		con:       con,
		dsn:       dsn,
	}, nil
}

// Connection represents the sqlx implementation of the sqldb.Connection interface
type Connection struct {
	*Queryable
	con *sqlx.DB
	dsn string
}

// Beginx is an Exec that accepts named params (ex where id=:user_id)
func (db *Connection) Beginx() (sqldb.Tx, error) {
	return NewTx(db.con)
}

// Close closes the database connection
func (db *Connection) Close() error {
	return db.con.Close()
}

// SQL returns the sql.DB object
func (db *Connection) SQL() *sql.DB {
	return db.con.DB
}

// DSN returns the DSN used to create the connection
func (db *Connection) DSN() string {
	return db.dsn
}
