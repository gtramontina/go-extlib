package sql

import (
	"database/sql"

	"github.com/gtramontina/go-extlib/must"
	"github.com/gtramontina/go-extlib/sql/internal"
)

type Stream[Type any] struct {
	zero Type
	rows *sql.Rows
}

func (s *Stream[Type]) Next() (Type, bool) {
	if !s.rows.Next() {
		must.NoError(s.rows.Close())
		must.NoError(s.rows.Err())

		return s.zero, false
	}

	result, fields := internal.NewOf[Type]()
	must.NoError(s.rows.Scan(fields...))

	return *result, true
}
