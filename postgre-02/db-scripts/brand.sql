create table main.brand
(
	id serial not null
		constraint brand_pk
			primary key,
	brand_name text not null
);

create unique index brand_name_uindex
	on main.brand (brand_name);


