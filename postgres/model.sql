create database realcargo;


drop table if exists items cascade;
drop table if exists cities cascade;
drop table if exists senders cascade;
drop table if exists receivers cascade;
drop table if exists orders cascade;


create table items (
	item_id serial not null primary key,
	item_name varchar(64) not null,
	quantity int not null
);

create table senders (
	sender_id serial not null primary key,
	sender_first_name varchar(32) not null,
	sender_last_name varchar(32) not null,
	phone_number varchar(32) not null,
	email varchar(64) not null,
	address varchar(64) not null,
	city_name varchar(32) not null,
	post_code varchar(16) not null
);

create table receivers (
	receiver_id serial not null primary key,
	receiver_first_name varchar(32) not null,
	receiver_last_name varchar(32) not null,
	passport_number varchar(16) not null,
	phone_number varchar(32) not null,
	address varchar(64) not null,
	city_name varchar(32) not null
);

create table orders (
	order_id serial not null primary key,
	comment text,
	status bool default false,
	sign_signature varchar(32) not null,
	created_at timestamp default current_timestamp,
	item_id int not null references items(item_id),
	sender_id int not null references senders(sender_id),
	receiver_id int not null references receivers(receiver_id)
);

create table admins (
	username varchar(32) not null primary key,
	password varchar(60) not null
);

create extension pgcrypto;

insert into admins (username, password) values ('javlon', crypt('cargo', gen_salt('bf')));
