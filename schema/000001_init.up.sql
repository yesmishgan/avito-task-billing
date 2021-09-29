CREATE TABLE clients
(
    id       serial       not null unique,
    username varchar(255) not null unique,
    balance  int
);

CREATE TABLE transactions
(
    id          serial not null unique,
    description varchar(255),
    sender      int    not null,
    destination int    not null,
    amount      int    not null
);

CREATE TABLE clients_transactions
(
    id             serial                                             not null unique,
    client_id      int references clients (id) on delete cascade      not null,
    transaction_id int references transactions (id) on delete cascade not null
);

CREATE
OR REPLACE FUNCTION write(name_client VARCHAR(255), amount INTEGER) returns int
language plpgsql AS
$$
declare
id_client integer;
BEGIN
    if
EXISTS (SELECT 1 FROM clients WHERE username = name_client) then
UPDATE clients
SET balance = balance + amount
WHERE username = name_client;
else
        INSERT INTO clients (username, balance)
        VALUES (name_client, amount);
END if;
select id
into id_client
from clients
where username = name_client;
return id_client;

END; $$

/*
CREATE OR REPLACE FUNCTION write(name_client VARCHAR(255), amount INTEGER) returns int
language plpgsql AS
$$
declare
id_client integer;
BEGIN
    if EXISTS (SELECT 1 FROM clients WHERE username = name_client) then
        UPDATE clients SET balance = balance + amount
        WHERE username = name_client;
    else
        INSERT INTO clients (username, balance)
        VALUES (name_client, amount);
    END if;
select id into id_client
from clients
where username = name_client;
return id_client;

END; $$

 */