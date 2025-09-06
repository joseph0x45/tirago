create table sessions (
  id text not null primary key,
  session_type text not null, -- can be 'regular' or 'admin'
  user_id text not null,
  valid boolean not null
);
