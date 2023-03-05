package sql

import (
	"database/sql"
	"errors"

	"github.com/gtramontina/go-extlib/must"
	"github.com/gtramontina/go-extlib/sql/internal"
)

// QueryRow forwards the execution to sql.DB.QueryRow and returns the result
// already scanned into the given type. The columns will be scanned in the
// order they are defined in the type. E.g.: if the type is:
//
//	type User struct {
//		ID   int
//		Name string
//	}
//
// The first column will be scanned into ID and the second into Name.
//
// If the query returns no rows, the returned boolean will be false.
func QueryRow[Type any](db *sql.DB, query string, args ...any) (Type, bool) {
	row := db.QueryRow(query, args...)
	must.NoError(row.Err())

	result, fields := internal.NewOf[Type]()

	err := row.Scan(fields...)
	if errors.Is(err, sql.ErrNoRows) {
		return *result, false
	}

	must.NoError(err)

	return *result, true
}

// Query forwards the execution to sql.DB.Query and returns a stream of results.
// The Stream will scan the results into the given type as Next is called. The
// columns will be scanned in the order they are defined in the type. E.g.: if
// the type is:
//
//	type User struct {
//		ID   int
//		Name string
//	}
//
// The first column will be scanned into ID and the second into Name.
// The stream will be closed when the query is done.
func Query[Type any](db *sql.DB, query string, args ...any) *Stream[Type] {
	rows := must.Return(db.Query(query, args...))

	return &Stream[Type]{rows: rows}
}
