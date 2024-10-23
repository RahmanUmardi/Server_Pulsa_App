CREATE DATABASE server_pulsa_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE roles AS ENUM ('admin', 'employee');

CREATE TABLE mst_supliyer(
    id_supliyer uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name_supliyer VARCHAR(255) NOT NULL,
    balance DOUBLE PRECISION NOT NULL
);

CREATE TABLE mst_product(
    id_product uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name_provider VARCHAR(255) NOT NULL,
    nominal DOUBLE PRECISION NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    id_supliyer uuid REFERENCES mst_supliyer(id_supliyer)
);

CREATE TABLE mst_user(
    id_user uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role roles NOT NULL
);

CREATE TABLE mst_merchant(
    id_merchant uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    id_user uuid REFERENCES mst_user(id_user),
    name_merchant VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    id_product uuid REFERENCES mst_produk(id_product),
    balance DOUBLE PRECISION NOT NULL
);

CREATE TABLE tx_trx_detail(
    id_tx_trx_detail uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    destination_number VARCHAR(15) NOT NULL,
    transaction_date DATE,
    total_price DOUBLE PRECISION NOT NULL
);

CREATE TABLE tx_trx(
    id_tx_trx uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    id_merchant uuid REFERENCES mst_merchant(id_merchant),
    id_tx_trx_detail uuid REFERENCES tx_trx_detail(id_tx_trx_detail),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);