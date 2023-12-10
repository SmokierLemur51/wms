-- to run: sqlite3 <datebase.db> < data/scripts/script.sql

drop table if exists products;
drop table if exists vendors;

create table vendors (
	id integer primary key autoincrement,
	vendor varchar(100),
	address_street text,
	address_street_2 text,
	address_state varchar(2),
	address_zip varchar(5)
);

create table products (
	id integer primary key autoincrement,
	vendor_id integer,
	product varchar(50) not null,
	product_code varchar(60) not null,
	description text,
	cost real,
	selling real,
	units_ctn integer,
	ctn_pallet integer,
	units_pallet integer,
	cost_pallet integer,
	foreign key (vendor_id) references vendors(id)
);