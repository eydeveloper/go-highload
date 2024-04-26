CREATE TYPE user_genders AS ENUM ('M', 'F');

CREATE TABLE users
(
    id uuid DEFAULT gen_random_uuid(),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    birth_date DATE NOT NULL,
    gender user_genders NOT NULL,
    biography TEXT,
    city VARCHAR(255),
    password_hash VARCHAR(255) NOT NULL,
    CONSTRAINT unique_user_id UNIQUE (id)
);