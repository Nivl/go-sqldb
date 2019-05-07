// Package sqldb contains interfaces and strutures to deal with SQL database
package sqldb

import "database/sql"

// Connection is an interface representing a database connection
type Connection interface {
	Queryable

	// SQL returns the sql.DB object
	SQL() *sql.DB

	// DSN returns the DNS used to connect to the database
	DSN() string

	// Close closes the database connection
	Close() error

	// Beginx starts a new transaction
	Beginx() (Tx, error)
}

// Queryable represents
type Queryable interface {
	// Get is used to retrieve a single row
	// An error (sql.ErrNoRows) is returned if the result set is empty.
	Get(dest interface{}, query string, args ...interface{}) error

	// NamedGet is a Get that accepts named params (ex where id=:user_id)
	NamedGet(dest interface{}, query string, args interface{}) error

	// Select is used to retrieve multiple rows
	Select(dest interface{}, query string, args ...interface{}) error

	// NamedSelect is a Select() that accepts named params (ex where id=:user_id)
	NamedSelect(dest interface{}, query string, args interface{}) error

	// Exec executes a SQL query and returns the number of rows affected
	Exec(query string, arg ...interface{}) (rowsAffected int64, err error)

	// NamedExec is an Exec that accepts named params (ex where id=:user_id)
	NamedExec(query string, arg interface{}) (rowAffected int64, err error)
}

// Tx is an interface representing a transaction
type Tx interface {
	Queryable

	// Commit commits the transaction
	Commit() error

	// Rollback rollbacks the transaction
	Rollback() error
}
