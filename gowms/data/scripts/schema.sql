-- to run: sqlite3 <datebase.db> < data/scripts/script.sql
drop table if exists vendors;
drop table if exists warehouse_bin;
drop table if exists stock_status;
drop table if exists products;

create table vendors (
	id integer primary key autoincrement,
	vendor varchar(100),
	address_street text,
	address_street_2 text,
	address_city text,
	address_state varchar(2),
	address_zip varchar(5)
);

create table warehouse_bin (
	id integer primary key autoincrement,
	bin varchar(100),
	aisle varchar(10),
	shelf varchar(10)
);

create table stock_status (
	id integer primary key autoincrement,
	disposition text
);

create table products (
	id integer primary key autoincrement,
	vendor_id integer,
	status_id integer,
	product varchar(50) not null,
	product_code varchar(60) not null,
	p_description text,
	units_ctn integer,
	ctn_pallet integer,
	units_pallet integer,
	cost_pallet real,
	selling_pallet real,
	cost_ctn real,
	selling_ctn real,
	cost_unit real,
	selling_unit real,
	wh_bin_id integer,
	foreign key (vendor_id) references vendors(id),
	foreign key (status_id) references stock_status(id)
	foreign key (wh_bin_id) references warehouse_bin(id)
);

