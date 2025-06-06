// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: uuid_ossp.sql

package querytest

import (
	"context"

	"github.com/google/uuid"
)

const generateUUID = `-- name: GenerateUUID :one
SELECT uuid_generate_v4()
`

func (q *Queries) GenerateUUID(ctx context.Context) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, generateUUID)
	var uuid_generate_v4 uuid.UUID
	err := row.Scan(&uuid_generate_v4)
	return uuid_generate_v4, err
}
