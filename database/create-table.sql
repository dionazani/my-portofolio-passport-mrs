create table app_person (
	id UUID,
	fullname varchar(50) not null,
	email varchar(50) not null unique,
	mobile_phone varchar(25) not null,
	created_at timestamp default current_timestamp,
	updated_at timestamp,
	primary key (id)
);

create table app_user (
	id UUID,
	app_user_id UUID not null,
	app_user_role char(3) not null,
	password_login varchar(300) not null,
	must_change_password int not null,
	next_change_password_date date,
	is_lock int default 0,
	created_at timestamp default current_timestamp,
	primary key (id),
	foreign key (app_user_id) references app_person (id)
);

