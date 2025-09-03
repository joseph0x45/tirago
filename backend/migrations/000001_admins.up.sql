create table admins (
  id text not null primary key,
  username text not null unique,
  password text not null
);
