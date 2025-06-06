// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package querytest

import (
	"context"
)

const getData = `-- name: GetData :many
select invalid, foo
from my_table
where (cast(?1 as boolean) or not invalid)
`

func (q *Queries) GetData(ctx context.Context, allowInvalid bool) ([]MyTable, error) {
	rows, err := q.db.QueryContext(ctx, getData, allowInvalid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MyTable
	for rows.Next() {
		var i MyTable
		if err := rows.Scan(&i.Invalid, &i.Foo); err != nil {
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
