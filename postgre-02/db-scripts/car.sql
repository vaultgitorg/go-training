create table if not exists main.car
(
	brand integer not null
		constraint car_brand_fk
			references main.brand,
	model integer not null
		constraint car_model_fk
			references main.model,
	id serial not null
		constraint car_pk
			primary key,
	ban_type text not null,
	miles integer,
	price numeric default 0 not null,
	currency text default 'AZN'::text not null
		constraint car_currency_chk
			check (currency in ('AZN', 'USD', 'EUR')),
	on_active_loan integer default 0 not null,
	can_be_bartered integer default 0 not null,
	description text
);

comment on column main.car.miles is 'yurus';

comment on column main.car.on_active_loan is '0 - false, 1 - true';

comment on column main.car.can_be_bartered is '0 - false, 1 - true';

