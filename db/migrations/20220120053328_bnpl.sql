-- migrate:up
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL UNIQUE,
    email VARCHAR(60) NOT NULL UNIQUE,
    credit_limit DECIMAL NOT NULL,
    due_amount DECIMAL NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE merchants(
    id SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL UNIQUE,
    email VARCHAR(60) NOT NULL UNIQUE,
    discount DECIMAL,
    total_payment DECIMAL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE transactions(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    merchant_id INTEGER NOT NULL,
    amount DECIMAL NOT NULL,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_transactions_user_id_users_id
        FOREIGN KEY (user_id)
        REFERENCES users(id),
    CONSTRAINT fk_transactions_merchant_id_merchants_id
        FOREIGN KEY (merchant_id)
        REFERENCES merchants(id)
);

-- migrate:down
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS merchants;



