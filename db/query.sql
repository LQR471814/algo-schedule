-- name: CreateTask :one
insert into task(name, description, deadline, size, challenge) values (?, ?, ?, ?, ?)
returning id;

-- name: ListTasks :many
select * from task where deleted_at is null;

-- name: ListDeletedTasks :many
select * from task where deleted_at is not null;

-- name: ReadTask :one
select * from task where id = ? and deleted_at is null;

-- name: UpdateTask :exec
update task set 
    name = ?,
    description = ?,
    deadline = ?,
    size = ?,
    challenge = ?
where id = ?;

-- name: DeleteTask :exec
update task set deleted_at = datetime('now') where id = ?;

-- name: UnDeleteTask :exec
update task set deleted_at = null where id = ?;


-- name: CreateProject :one
insert into project(name, description, deadline) values (?, ?, ?)
returning id;

-- name: ListProjects :many
select * from project where deleted_at is null;

-- name: ListDeletedProjects :many
select * from project where deleted_at is not null;

-- name: ReadProject :one
select * from task where id = ? and deleted_at is null;

-- name: UpdateProject :exec
update project set
    name = ?,
    description = ?,
    deadline = ?
where id = ?;

-- name: DeleteProject :exec
update project set deleted_at = datetime('now') where id = ?;

-- name: UnDeleteProject :exec
update project set deleted_at = null where id = ?;


-- name: CreateProjectTask :one
insert into project_task(project_id, name, description, size, challenge) values (?, ?, ?, ?, ?)
returning id;

-- name: ListProjectTasks :many
select * from project_task where project_id = ? and deleted_at is null;

-- name: ListDeletedProjectTasks :many
select * from project_task where project_id = ? and deleted_at is not null;

-- name: ReadProjectTask :one
select * from project_task where id = ? and deleted_at is null;

-- name: UpdateProjectTask :exec
update project_task set
    name = ?,
    description = ?,
    size = ?,
    challenge = ?
where id = ?;

-- name: DeleteProjectTask :exec
update project_task set deleted_at = datetime('now') where id = ?;

-- name: UnDeleteProjectTask :exec
update project_task set deleted_at = null where id = ?;



-- name: CreateQuota :one
insert into quota(fixed_time, duration, recurrence_interval) values (?, ?, ?)
returning id;

-- name: ListQuotas :many
select * from quota where deleted_at is null;

-- name: ListDeletedQuotas :many
select * from quota where deleted_at is not null;

-- name: ReadQuota :one
select * from quota where id = ? and deleted_at is null;

-- name: UpdateQuota :exec
update quota set
    fixed_time = ?,
    duration = ?,
    recurrence_interval = ?
where id = ?;

-- name: DeleteQuota :exec
update quota set deleted_at = datetime('now') where id = ?;

-- name: UnDeleteQuota :exec
update quota set deleted_at = null where id = ?;



-- name: UpdateSettings :exec
insert into settings(id, timezone) values (1, ?)
on conflict do update set
    timezone = excluded.timezone;

