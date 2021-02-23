CREATE DATABASE go_rest_api;

\c go_rest_api;

CREATE TABLE IF NOT EXISTS customers (
  id          BIGINT NOT NULL,
  email       VARCHAR,
  balance     NUMERIC   DEFAULT 0.0,
  active      BOOLEAN   DEFAULT true,
  join_date   DATE
);

ALTER TABLE ONLY customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);

CREATE UNIQUE INDEX unique_email_customers
  ON customers(email);

CREATE SEQUENCE customers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
    OWNED BY customers.id;

INSERT INTO customers (id, email, balance, active, join_date)
  VALUES (nextval('customers_id_seq'), 'mkdika@gmail.com', 99.75, true, NOW());

INSERT INTO customers (id, email, balance, active, join_date)
  VALUES (nextval('customers_id_seq'), 'john@gmail.com', 10.5, false, NOW());