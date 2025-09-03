create table admins (
  id serial primary key,
  username text not null unique,
  password text not null
);
