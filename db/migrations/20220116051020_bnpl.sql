-- migrate:up
create table users (
	id serial PRIMARY KEY,
	name VARCHAR (60) UNIQUE NOT NULL,
	email VARCHAR (60) UNIQUE NOT NULL,
    credit_limit INTEGER,
    due_amount INTEGER,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

create table merchants (
	id serial PRIMARY KEY,
	name VARCHAR (60) UNIQUE NOT NULL,
	email VARCHAR (60) UNIQUE NOT NULL,
    discount INTEGER,
    total_payment INTEGER,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);


-- migrate:down
DROP TABLE users;
DROP TABLE merchants;

