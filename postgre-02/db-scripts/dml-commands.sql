delete from car where currency is null;
commit;

insert into car(brand, model, ban_type, miles, description, currency, price, can_be_bartered)
values (1, 7, 'traktor', 78000, 'teze', 'azn', 2222 , 1);
commit;

select m.id, m.model_name, b.brand_name from model m
join brand b on b.id = m.brand;


select * from model where brand = 1;
insert  into model(brand, model_name)  values(1, 'X5');insert  into model(brand, model_name)  values(1, 'X3');
commit;

select * from brand;
select * from model where brand = 2;


update car set can_be_bartered = 1 where id = 2;
commit;


select * from car;

select c.id, b.brand_name, m.model_name, c.ban_type, c.miles, c.price, c.currency, c.on_active_loan, c.can_be_bartered, c.description from car c
join brand b on c.brand = b.id
join model m on m.id = c.model
where c.id = 5;




ALTER TABLE car ADD CONSTRAINT car_currency_chk CHECK (	currency in ('AZN', 'USD', 'EUR') );



select * from car order by 3 desc

