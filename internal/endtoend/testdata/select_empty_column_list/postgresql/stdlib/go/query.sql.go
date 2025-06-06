// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"
)

const getBars = `-- name: GetBars :many
SELECT FROM bar LIMIT 5
`

type GetBarsRow struct {
}

func (q *Queries) GetBars(ctx context.Context) ([]GetBarsRow, error) {
	rows, err := q.db.QueryContext(ctx, getBars)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetBarsRow
	for rows.Next() {
		var i GetBarsRow
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
