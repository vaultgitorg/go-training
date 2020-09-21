create table if not exists main.model
(
	id serial not null
		constraint model_pk
			primary key,
	brand integer
		constraint model_brand_fk
			references main.brand,
	model_name text not null
);


