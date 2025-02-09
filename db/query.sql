-- name: CreateTask :exec
insert into task(id, name, description, deadline, size, challenge) values (?, ?, ?, ?, ?, ?);

-- name: ListTasks :many
select * from task;

-- name: ReadTask :one
select * from task where id = ?;

-- name: UpdateTask :exec
update task set 
    name = ?,
    description = ?,
    deadline = ?,
    size = ?,
    challenge = ?
where id = ?;

-- name: DeleteTask :exec
delete from task where id = ?;

-- name: CreateProject :exec
insert into project(id, name, description, deadline) values (?, ?, ?, ?);

-- name: ListProjects :many
select * from project;

-- name: ReadProject :one
select * from task where id = ?;

-- name: UpdateProject :exec
update project set
    name = ?,
    description = ?,
    deadline = ?
where id = ?;

-- name: DeleteProject :exec
delete from project where id = ?;

-- name: CreateProjectTask :exec
insert into project_task(id, project_id, name, description, size, challenge) values (?, ?, ?, ?, ?, ?);

-- name: ListProjectTasks :many
select * from project_task where project_id = ?;

-- name: ReadProjectTask :one
select * from project_task where id = ?;

-- name: UpdateProjectTask :exec
update project_task set
    name = ?,
    description = ?,
    size = ?,
    challenge = ?
where id = ?;

-- name: DeleteProjectTask :exec
delete from project_task where id = ?;

-- name: CreateQuota :exec
insert into quota(id, fixed_time, duration, recurrence_interval) values (?, ?, ?, ?);

-- name: ListQuotas :many
select * from quota;

-- name: ReadQuota :one
select * from quota where id = ?;

-- name: UpdateQuota :exec
update quota set
    fixed_time = ?,
    duration = ?,
    recurrence_interval = ?
where id = ?;

-- name: DeleteQuota :exec
delete from quota where id = ?;

