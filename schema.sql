create table if not exists expenses(
  id serial primary key not null,
  username varchar (50) not null,
  category varchar (50) not null,
  amount int not null,
  spent_at timestamp not null,
  created_at timestamp not null default now()
);

# seed data
insert into expenses (username, category, amount, created_at) VALUES ('Jaap', 'Edeka', 12.20, now());
