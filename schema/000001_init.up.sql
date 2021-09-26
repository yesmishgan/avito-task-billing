CREATE TABLE clients
(
    id       serial       not null unique,
    username varchar(255) not null unique,
    balance  int
);

CREATE TABLE transactions
(
    id          serial                                        not null unique,
    description varchar(255),
    client_id   int references clients (id) on delete cascade not null,
    destination varchar(255)                                  not null,
    amount      int                                           not null
);