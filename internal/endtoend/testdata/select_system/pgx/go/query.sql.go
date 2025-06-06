// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getSystemColumns = `-- name: GetSystemColumns :one
SELECT
  tableoid, xmin, cmin, xmax, cmax, ctid
FROM test
`

type GetSystemColumnsRow struct {
	Tableoid pgtype.Uint32
	Xmin     pgtype.Uint32
	Cmin     pgtype.Uint32
	Xmax     pgtype.Uint32
	Cmax     pgtype.Uint32
	Ctid     pgtype.TID
}

func (q *Queries) GetSystemColumns(ctx context.Context) (GetSystemColumnsRow, error) {
	row := q.db.QueryRow(ctx, getSystemColumns)
	var i GetSystemColumnsRow
	err := row.Scan(
		&i.Tableoid,
		&i.Xmin,
		&i.Cmin,
		&i.Xmax,
		&i.Cmax,
		&i.Ctid,
	)
	return i, err
}
