create table if not exists forgotten_password_keys (
    id serial not null
        constraint forgotten_password_keys_pkey
        primary key,
    user_id int not null
        constraint forgotten_password_keys_user_fkey
        references users,
    key varchar(64) not null,
    valid timestamp with time zone not null
)