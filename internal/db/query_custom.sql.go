package db

import (
	"context"
	"time"
)

const clearDeletedBefore = `-- name: ClearDeletedBefore :exec
begin transaction;
    delete from task where deleted_at < $1;
    delete from project_task where deleted_at < $1;
    delete from project where deleted_at < $1;
    delete from quota where deleted_at < $1;
commit;`

func (q *Queries) ClearDeletedBefore(ctx context.Context, before time.Time) error {
	_, err := q.db.ExecContext(ctx, updateTask, before)
	return err
}
