package postgres

const foreignKeyOn = `PRAGMA foreign_keys = ON;`

const clientsTable = `create table if not exists clients(
	id bigserial primary key,
	name varchar not null,
	surname text not null,
	login text not null unique,
	password text not null,
	age integer not null,
	gender text not null,
	phone text not null,
	status boolean not null
);`

const clientsAccountsTable = `create table if not exists accounts(
                                       id bigserial primary key,
                                       client_id integer references clients,
                                       account_number integer not null,
                                       balance integer not null check ( balance >= 0 ),
                                       status boolean not null,
                                       card_number text not null unique,
                                       limit_transfer integer not null default 3000,
                                       limit_payment integer not null default 4000,
                                       created_at time default CURRENT_TIMESTAMP,
                                       until_at time default CURRENT_TIMESTAMP
);`

const ATMsTable = `create table if not exists atms(
                                   id bigserial primary key,
                                   address_id integer,
                                   status boolean not null,
                                   created_at time default CURRENT_TIMESTAMP
);`

const servicesTable = `create table if not exists services(
                                       id bigserial primary key,
                                       name text not null
);`

const historiesTable = `create table if not exists histories(
                                         id bigserial primary key,
                                         sender_id integer references clients,
                                         recipient_id integer references clients,
                                         money integer not null,
                                         message text not null,
                                         service_id integer references services,
                                         created_at time default CURRENT_TIMESTAMP
);`

const atmAddresses = `create table if not exists addresses(
                                        id bigserial primary key,
                                        country text not null,
                                        city text not null,
                                        street text not null,
                                        home integer not null,
                                        apartment integer not null
);`
