create table sessions (
  id text not null primary key,
  session_type text not null, --can be 'admin' or 'regular'
  user_id text not null,
  valid boolean not null default true
);
