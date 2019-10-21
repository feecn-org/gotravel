--go travel 数据库
create sequence t_destination_req_id_seq
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

create table t_destination_req
(
	id integer default nextval('t_destination_req_id_seq'::regclass) not null
		constraint t_destination_req_pk
			primary key,
	doName varchar(1024),
	body varchar(2048),
	method varchar(64),
	param varchar(4096)
);