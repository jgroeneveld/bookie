create table if not exists expenses(
  id serial primary key not null,
  username varchar (50),
  category varchar (50),
  amount int,
  created_at timestamp not null default now()
);
