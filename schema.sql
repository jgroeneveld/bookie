create table if not exists expenses(
  id serial PRIMARY KEY NOT NULL,
  username varchar (50),
  category varchar (50),
  amount int,
  created_at timestamp not null
);
