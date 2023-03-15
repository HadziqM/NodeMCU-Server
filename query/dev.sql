create table if not exists flow_sens(
  id serial,
  sens_id int,
  flow real,
  created_at timestamp without time zone default now(),
  PRIMARY KEY(id)
);
