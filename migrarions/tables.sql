-- Active: 1677320843749@@127.0.0.1@5432@postgres@public

DROP TABLE IF EXISTS user_balance;
CREATE Table user_balance (
    user_id INTEGER PRIMARY KEY,
    name VARCHAR(30),
    adress VARCHAR(30),
    balance INT

)

CREATE TABLE reserve (
    user_id INTEGER,
    name VARCHAR(30),
    service_id INTEGER,
    order_id INTEGER,
    cost INTEGER,
    сonfirmation VARCHAR(30)
)

CREATE TABLE profit(

    user_id INTEGER,
    name VARCHAR(30),
    service_id INTEGER,
    order_id INTEGER,
    sum INTEGER

)