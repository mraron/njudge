create table if not exists users
(
    id             serial not null
        constraint table_name_pkey
            primary key,
    name           text   not null,
    password       text   not null,
    email          text   not null,
    activation_key text,
    role           text   not null
);

create table if not exists problem_rels
(
    problemset text   not null,
    problem    text   not null,
    id         serial not null
        constraint problem_rels_id_pk
            primary key
);

create table if not exists judges
(
    id     serial                not null
        constraint judges_pkey
            primary key,
    state  text                  not null,
    host   text                  not null,
    port   text                  not null,
    ping   integer default 0     not null,
    online boolean default false not null
);

create table if not exists submissions
(
    id         serial                   not null
        constraint submissions_pkey
            primary key,
    status     text                     not null,
    ontest     text,
    user_id    integer                  not null
        constraint submissions_users_id_fk
            references users,
    problemset text                     not null,
    problem    text                     not null,
    language   text                     not null,
    private    boolean                  not null,
    verdict    integer                  not null,
    source     bytea                    not null,
    started    boolean                  not null,
    submitted  timestamp with time zone not null,
    judged     timestamp with time zone
);

