package sqlxdb

import (
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	sqldb "github.com/Nivl/go-sqldb"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

// sqlxQueryable is an interface used to group sqlx.Tx and sqlx.DB
type sqlxQueryable interface {
	Get(dest interface{}, query string, args ...interface{}) error
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	Select(dest interface{}, query string, args ...interface{}) error
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Rebind(query string) string
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
	var err error
	query, args, err = db.handleInClauses(query, args)
	if err != nil {
		return err
	}
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
	var err error
	query, args, err = db.handleInClauses(query, args)
	if err != nil {
		return err
	}
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
	query, args, err = db.handleInClauses(query, args)
	if err != nil {
		return 0, err
	}

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

type argInfo struct {
	reflectValue reflect.Value

	// arg corresponds to the arg itself (often a slice)
	arg interface{}

	// len contains the number of item in the Slice (arg) (0 if not a slice)
	len int

	// pos contains the position of the arg in the arg list.
	// for `IN ($3)`, pos will be 2 (3-1 since args starts at 0)
	pos int

	// originalBindvarNum contains the original position of the bindvar
	// for `IN ($3)`, originalBindvarNum will be 3
	originalBindvarNum int
}

type byPos []argInfo

func (a byPos) Len() int           { return len(a) }
func (a byPos) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byPos) Less(i, j int) bool { return a[i].pos < a[j].pos }

// handleInClauses handles queries containing IN clauses.
// ex: "IN (?)" or "IN ($1)"
func (db *Queryable) handleInClauses(query string, args []interface{}) (finalQuery string, finalArgs []interface{}, err error) {
	// Find all $x placeholders
	re := regexp.MustCompile(`\$(\d+)`)
	matches := re.FindAllStringSubmatch(query, -1)
	highestNumber := 0
	for _, match := range matches {
		im, _ := strconv.Atoi(match[1])
		if im > highestNumber {
			highestNumber = im
		}
	}
	if highestNumber == 0 {
		// Let's check if we have any `IN (?)`, in which case sqlx can handle
		// everything by itself
		re = regexp.MustCompile(`IN ?\( ?\? ?\)`)
		hasBindvar := len(re.FindStringSubmatch(query)) > 0
		if hasBindvar {
			query, args, err := sqlx.In(query, args...)
			return query, args, err
		}
		return query, args, nil
	}

	// Find all IN clauses
	re = regexp.MustCompile(`IN ?\( ?\$(\d+) ?\)`)
	globalMatches := re.FindAllStringSubmatch(query, -1)
	if len(globalMatches) == 0 {
		return query, args, nil
	}

	totalArgs := len(args)
	infoList := []argInfo{}
	// for every IN clause we have we group their info and put them in infoList
	for _, match := range globalMatches {
		argPos, _ := strconv.Atoi(match[1])
		argPos-- // $1 is args[0]

		if argPos >= totalArgs {
			return "", nil, fmt.Errorf("$%s is missing from the args", match[1])
		}

		info := argInfo{
			originalBindvarNum: argPos + 1,
			pos:                argPos,
			arg:                args[argPos],
			reflectValue:       reflect.ValueOf(args[argPos]),
		}

		// if it's not a slice then we continue
		t := reflectx.Deref(info.reflectValue.Type())
		if t.Kind() != reflect.Slice {
			continue
		}

		// We don't accept empty slice
		info.len = info.reflectValue.Len()
		if info.len == 0 {
			return "", nil, fmt.Errorf("failed to parse IN clause: arg at position %d is an empty slice", argPos)
		}

		infoList = append(infoList, info)
	}

	// We order all the info because we might have something like:
	// WHERE x IN ($3) AND y IN ($1) AND z=$2
	// and we want to deal with $1 first, then $3
	sort.Sort(byPos(infoList))

	finalQuery = query
	previousArgPos := -1
	// bindShift contains the current number of shifting made on the bindvars
	// EX. if we have `x IN ($1) and y=$2` and we have a slice of 3
	// bindShift will be increased by 2 so we can have
	// x IN ($1, $2, $3) and y=$4
	bindShift := 0

	for _, info := range infoList {
		// We insert all non-IN args we had between this clause and the previous
		// one
		// EX: WHERE x IN ($1) AND y=$2 AND z IN ($3)
		// We're at $3, so we need to insert $2
		if info.pos-previousArgPos > 1 {
			for i := previousArgPos + 1; i < info.pos; i++ {
				finalArgs = append(finalArgs, args[i])
			}
		}
		previousArgPos = info.pos

		// We insert the slice content into the finalArgs array
		for si := 0; si < info.len; si++ {
			finalArgs = append(finalArgs, info.reflectValue.Index(si).Interface())
		}

		// If the current slice only has one elem then no need to update the query
		if info.len == 1 {
			continue
		}

		// rename all the bindvars depending an how many elements we are inserting
		// Ex. if we are inserting 2 elements for $3, All bindvars above $3 will
		// have their number increased by 2 ($4 => $6, $5 => $7, etc.)
		newHigh := highestNumber + bindShift
		newCurrent := info.originalBindvarNum + bindShift
		addedPadding := info.len - 1
		for i := newHigh; i > newCurrent; i-- {
			oldBindvar := fmt.Sprintf("$%d", i)
			newBindvar := fmt.Sprintf("$%d", i+addedPadding)
			finalQuery = strings.ReplaceAll(finalQuery, oldBindvar, newBindvar)
		}

		// insert all the new bindvar inside the IN clause
		// Ex: if we are inserting 2 elements for $3:
		// `IN ($3)` needs to be `IN ($3, $4)`
		bindvarToInsert := []string{}
		for i := 0; i < info.len; i++ {
			bindvar := fmt.Sprintf("$%d", newCurrent+i)
			bindvarToInsert = append(bindvarToInsert, bindvar)
		}
		bindvarInOriginalQuery := fmt.Sprintf("$%d", newCurrent)
		finalQuery = strings.ReplaceAll(finalQuery, bindvarInOriginalQuery, strings.Join(bindvarToInsert, ", "))

		bindShift += addedPadding
	}

	// Add the last params (if there are any)
	for i := previousArgPos + 1; i < totalArgs; i++ {
		finalArgs = append(finalArgs, args[i])
	}

	return finalQuery, finalArgs, nil
}
