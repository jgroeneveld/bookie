create table if not exists expenses(
  id serial primary key not null,
  username varchar (50) not null,
  category varchar (50) not null,
  amount int not null,
  spent_at timestamp not null,
  created_at timestamp not null default now()
);

