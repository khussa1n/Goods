CREATE TABLE projects (
    id                  serial primary key,
    name			    varchar(255) not null,
    created_at          timestamp not null
);

CREATE TABLE goods (
    id                  serial primary key,
    project_id          integer references projects(id) not null,
    name			    varchar(255) not null,
    description         text not null,
    priority            integer not null,
    removed             boolean not null,
    created_at          timestamp not null
);

INSERT INTO projects (
    name, -- 1
    created_at -- 2
) VALUES (
    'Первая запись',
    CURRENT_TIMESTAMP
);

