CREATE TABLE IF NOT EXISTS Account (
    id SERIAL PRIMARY KEY,
    names VARCHAR(255) NOT NULL,
    usernames VARCHAR(255) NOT NULL,
    emails VARCHAR(255) NOT NULL UNIQUE,
    countries VARCHAR(255) NOT NULL,
    shipping_addresses VARCHAR(255) NOT NULL
);
