-- no 1 create database ata_online_shop
create database ata_online_shop;

use ata_online_shop;

-- no 2a (create table user)
create table users(
	id int(11) primary key,
    name varchar(255),
    status smallint,
    created_at timestamp,
    updated_at timestamp
);

-- no 2b (create table product, product_type, operator, product_description, payment_method)

-- create table products
create table products(
	id int(11) primary key,
    product_type_id int(11),
    operator_id int(11),
    code varchar(50),
    name varchar(100),
    status smallint,
   created_at timestamp,
   updated_at timestamp
);

-- create table product_types
create table product_types(
	id int(11) primary key,
    name varchar(255),
	created_at timestamp,
   updated_at timestamp
);

-- create table operators
create table operators(
	id int(11) primary key,
    name varchar(30),
    created_at timestamp,
    updated_at timestamp
);

-- create table product_description
create table product_descriptions(
	id int(11) primary key,
    description text,
    created_at timestamp,
    updated_at timestamp
);

-- create table payment_methods
create table payment_methods(
	id int(11) primary key,
    name varchar(255),
    status smallint,
    created_at timestamp,
    updated_at timestamp
);

-- no 2c (create table tansaction, transaction_detail)

-- create table transactions
create table transactions(
	id int(11) primary key,
    user_id int(11),
    payment_method_id int(11),
    status varchar(10),
    total_qty int(11),
    total_price numeric(25,2),
    created_at timestamp,
    updated_at timestamp
);

-- create table transaction_details
create table transaction_details(
-- 	id int(11) primary key,
	transaction_id int(11),
    product_id int(11),
    status varchar(10),
    qty int(11),
    price numeric(25,2),
    created_at timestamp,
    updated_at timestamp
);

-- no 3 (create table kurir)
create table kurir(
	id int(11) primary key,
    name varchar(255),
    created_at timestamp,
    updated_at timestamp
);

-- no 4 (tambahkan ongkos_dasar column di table kurir)
alter table kurir
add column ongkos_dasar numeric(25,2);

-- no 5 rename table kurir menjadi shipping
rename table kurir to shipping;

-- no 6 drop table shipping
drop table shipping;


-- no 7 (tambahkan relasi)

-- no 7a one to one (product to product description)

-- tambah column product_description_id di table products
alter table products add column product_description_id int(11);

-- jadikan kolom product_description_id sebagai foreign key di table product dari tabel product_description
alter table products add foreign key (product_description_id) references product_descriptions(id);

-- no 7b one to many 

-- product_types to product
-- tambahkan foreign key (product_type_id) di table products
alter table products add foreign key (product_type_id) references product_types(id);

-- operator to product
alter table products add foreign key (operator_id) references operators(id);

-- payment_methods to transactions
alter table transactions add foreign key (payment_method_id) references payment_methods(id);

-- users to transaction
alter table transactions add foreign key (user_id) references users(id);


-- no 7c many to many 
-- transaction to product (menghasilkan tabel transaction detail)
-- tambahkan foreign key (transaction_id, product_id) di tabel transaction_detail
alter table transaction_details add foreign key (transaction_id) references transactions(id);
alter table transaction_details add foreign key (product_id) references products(id);
