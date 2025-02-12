create table task (
    id integer not null primary key autoincrement,
    deleted_at datetime default null,

    name text not null unique,
    description text not null,
    deadline timestamp not null,
    -- 0: small, 1: medium
    size integer not null check(size in (0, 1)),
    -- 0: easy, 1: medium, 2: hard
    challenge integer not null check(challenge in (0, 1, 2))
);

create table project (
    id integer not null primary key autoincrement,
    deleted_at datetime default null,

    name text not null unique,
    description text not null,
    deadline timestamp not null
);

create table project_task (
    id integer not null primary key autoincrement,
    project_id integer not null,
    deleted_at datetime default null,

    name text not null unique,
    description text not null,
    -- 0: small, 1: medium
    size integer not null check(size in (0, 1)),
    -- 0: easy, 1: medium, 2: hard
    challenge integer not null check(challenge in (0, 1, 2)),

    foreign key (project_id) references project(id)
);

create table quota (
    id integer not null primary key autoincrement,
    deleted_at datetime default null,

    description not null,
    -- this fixed time is in minutes since the start of the day
    fixed_time integer not null,
    -- this duration is in minutes
    duration integer not null,
    -- the duration (in days) between each refresh of the quota
    -- if it is < 0, then the quota repeats on custom logic
    recurrence_interval integer not null
);

-- there is only ever one row of the settings table.
-- the id is always 1
create table settings (
    id integer not null primary key,
    timezone text not null
);

