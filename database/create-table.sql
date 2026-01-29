-- 1. Sign Up Table (Entry point)
CREATE TABLE sign_up (
    id UUID PRIMARY KEY,
    fullname VARCHAR(100) NOT NULL,
    sign_up_from CHAR(3) NOT NULL, -- e.g., 'WEB', 'MOB'
    email VARCHAR(255),
    mobile_phone VARCHAR(25),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- Ensures that at least one contact method is provided
    CONSTRAINT chk_contact_presence CHECK (email IS NOT NULL OR mobile_phone IS NOT NULL)
);
-- sign_up, app_person, app_person, app_user
-- 2. Person Table (The core identity)
CREATE TABLE app_person (
    id UUID PRIMARY KEY,
    fullname VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE,
    mobile_phone VARCHAR(25) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

-- create index idx_app_person_email on app_person (email);

-- 3. Mapping Table (Tracks which sign-up record created which person)
CREATE TABLE sign_up_app_person (
    sign_up_id UUID PRIMARY KEY, -- Changed to PK if 1:1 relationship
    app_person_id UUID NOT NULL,
    FOREIGN KEY (sign_up_id) REFERENCES sign_up (id),
    FOREIGN KEY (app_person_id) REFERENCES app_person (id)
);

-- 4. User Table (Login credentials and security)
CREATE TABLE app_user (
    id UUID PRIMARY KEY,
    app_person_id UUID NOT NULL UNIQUE, -- Link to app_person
    app_user_role CHAR(3) NOT NULL,
    app_password VARCHAR(300) NOT NULL, -- Renamed for clarity
    must_change_password int default 0,
    next_change_password_date DATE DEFAULT CURRENT_DATE + INTERVAL '30 days',
    is_locked int DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (app_person_id) REFERENCES app_person (id)
);

create table app_user_activate (
	id UUID primary key,
	app_user_id UUID not null,
	activate_by varchar(6) not null,
	activate_date date,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	foreign key (app_user_id) references app_user (id)
);
