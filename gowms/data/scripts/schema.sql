-- to run: sqlite3 <datebase.db> < data/scripts/script.sql
drop table if exists vendors;
drop table if exists warehouse_bin;
drop table if exists stock_status;
drop table if exists products;

drop table if exists warehouse;
drop table if exists aisle;
drop table if exists shelf;
drop table if exists warehouse_bin;

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

create table warehouse(
	id integer primary key autoincrement,
	warehouse varchar(60),
	street_1 varchar(100),
	street_2 varchar(100),
	city varchar(60),
	state varchar(2),
	zip varchar(5)
);

create table aisle(
	id integer primary key autoincrement,
	warehouse_id integer,
	aisle text,
	foreign key (warehouse_id) references warehouse(id)
);

create table shelf(
	id integer primary key autoincrement,
	warehouse_id integer,
	asile_id integer,
	foreign key (warehouse_id) references warehouse(id),
	foreign key (asile_id) references aisle(id)
);

create table warehouse_bin(
	id integer primary key autoincrement,
	warehouse_id integer,
	asile_id integer,
	shelf_id integer,
	rfid text,
	foreign key (warehouse_id) references warehouse(id),
	foreign key (aisle_id) references aisle(id),
	foreign key (shelf_id) references shelf(id)
);	