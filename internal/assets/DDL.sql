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
    price DECIMAL(10, 2) NOT NULL
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
    id_product uuid REFERENCES mst_product(id_product),
    balance DOUBLE PRECISION
);

CREATE TABLE transactions(
    transaction_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    id_merchant UUID REFERENCES mst_merchant(id_merchant),
    id_user UUID REFERENCES mst_user(id_user),
    customer_name VARCHAR(255) NOT NULL,
    destination_number VARCHAR(15) NOT NULL,
    transaction_date DATE
);

CREATE TABLE transaction_detail(
    transaction_detail_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    transaction_id UUID REFERENCES transactions(transaction_id),
    id_product UUID REFERENCES mst_product(id_product),
    price DECIMAL(10, 2) NOT NULL
);

CREATE TABLE tx_topup (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    id_merchant UUID REFERENCES mst_merchant(id_merchant),
    id_supliyer UUID REFERENCES mst_supliyer(id_supliyer),
    item_name VARCHAR(255) NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    payment_method VARCHAR(255),
    status VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW()
);
