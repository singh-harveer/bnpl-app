CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(255) primary key);
CREATE TABLE sqlite_sequence(name,seq);
CREATE TABLE users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR (60) UNIQUE NOT NULL,
	email VARCHAR (60) UNIQUE NOT NULL,
    credit_limit INTEGER,
    due_amount INTEGER,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
CREATE TABLE merchants (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR (60) UNIQUE NOT NULL,
	email VARCHAR (60) UNIQUE NOT NULL,
    discount INTEGER,
    total_payment INTEGER,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20220116051020');
