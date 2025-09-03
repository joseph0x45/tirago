create table if not exists admins (
  id serial primary key,
  username text not null,
  password text not null
);

create table if not exists membership_requests (
  id text not null primary key,
  email text not null unique,
  phone text not null unique,
  full_name text not null,
  account_type text not null,
  status text not null default "pending",
  created_at datetime not null default now(),
  refusal_reason text not null default ""
);

create table if not exists membership_request_docs (
  id text not null primary key,
  request_id text not null references membership_requests(id),
  document_url text not null
);

create table if not exists users (
  id text not null primary key,
  email text not null unique,
  phone text not null unique,
  town text not null,
  password text not null,
  full_name text not null,
  profile_picture text not null default 'https://picsum.photos/200/300',
  account_type text not null,
  created_at datetime not null default now()
);
