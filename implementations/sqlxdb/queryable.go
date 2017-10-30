package sqlxdb

import (
	"database/sql"

	sqldb "github.com/Nivl/go-sqldb"
	"github.com/jmoiron/sqlx"
)

// sqlxQueryable is an interface used to group sqlx.Tx and sqlx.DB
type sqlxQueryable interface {
	Get(dest interface{}, query string, args ...interface{}) error
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	Select(dest interface{}, query string, args ...interface{}) error
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var _ sqldb.Queryable = (*Queryable)(nil)

// NewQueryable creates a new Queryable
func NewQueryable(con sqlxQueryable) *Queryable {
	return &Queryable{con: con}
}

// Queryable represents the sqlx implementation of the Queryable interface
type Queryable struct {
	con sqlxQueryable
}

// Get is used to retrieve a single row
// An error (sql.ErrNoRows) is returned if the result set is empty.
func (db *Queryable) Get(dest interface{}, query string, args ...interface{}) error {
	return db.con.Get(dest, query, args...)
}

// NamedGet is a Get() that accepts named params (ex where id=:user_id)
func (db *Queryable) NamedGet(dest interface{}, query string, args interface{}) error {
	namedStmt, err := db.con.PrepareNamed(query)
	if err != nil {
		return err
	}
	return namedStmt.Get(dest, args)
}

// Select is used to retrieve multiple rows
func (db *Queryable) Select(dest interface{}, query string, args ...interface{}) error {
	return db.con.Select(dest, query, args...)
}

// NamedSelect is a Select() that accepts named params (ex where id=:user_id)
func (db *Queryable) NamedSelect(dest interface{}, query string, args interface{}) error {
	namedStmt, err := db.con.PrepareNamed(query)
	if err != nil {
		return err
	}
	return namedStmt.Select(dest, args)
}

// Exec executes a SQL query and returns the number of rows affected
func (db *Queryable) Exec(query string, args ...interface{}) (rowsAffected int64, err error) {
	res, err := db.con.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// NamedExec is an Exec that accepts named params (ex where id=:user_id)
func (db *Queryable) NamedExec(query string, arg interface{}) (rowAffected int64, err error) {
	res, err := db.con.NamedExec(query, arg)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
