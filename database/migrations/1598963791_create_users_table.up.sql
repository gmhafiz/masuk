create table users
(
    id bigserial primary key,
    username text unique,
    email text unique,
    password text,
    first_name text,
    last_name text,
    mobile text,
    phone text,
    active boolean default false ,
    created_at timestamptz set default now();,
    updated_at timestamptz set default now();,
    deleted_at timestamptz
)