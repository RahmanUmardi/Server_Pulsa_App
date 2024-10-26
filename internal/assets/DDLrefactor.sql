CREATE DATABASE server_pulsa_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE roles AS ENUM ('admin', 'employee');

CREATE TABLE mst_supplier(
    supplier_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    supplier_name VARCHAR(255) NOT NULL,
    balance DECIMAL(10, 2) NOT NULL
);

CREATE TABLE mst_product(
    product_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    provider_name VARCHAR(255) NOT NULL,
    nominal DECIMAL(10, 2) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    supplier_id UUID REFERENCES mst_supplier(supplier_id)
);

CREATE TABLE mst_user(
    user_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role roles NOT NULL
);

CREATE TABLE mst_merchant(
    merchant_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id UUID REFERENCES mst_user(user_id),
    merchant_name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    product_id UUID REFERENCES mst_product(product_id),
    balance DECIMAL(10, 2) NOT NULL,
);

CREATE TABLE transactions(
    transactions_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    merchant_id UUID REFERENCES mst_merchant(merchant_id),
    user_id UUID REFERENCES mst_user(user_id),
    customer_name VARCHAR(255) NOT NULL,
    destination_number VARCHAR(15) NOT NULL,
    transaction_date DATE
);

CREATE TABLE transaction_detail(
    transaction_detail_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    transactions_id UUID REFERENCES transactions(transactions_id),
    product_id UUID REFERENCES mst_product(product_id),
    total_price DECIMAL(10, 2) NOT NULL
);


