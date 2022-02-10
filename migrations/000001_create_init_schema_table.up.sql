CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    email VARCHAR,
    password VARCHAR,
    first_name VARCHAR,
    last_name VARCHAR,
    title VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories_kk
(
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR
);

CREATE TABLE IF NOT EXISTS categories_ru
(
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR
);

CREATE TABLE IF NOT EXISTS categories_en
(
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR
);

CREATE TABLE IF NOT EXISTS mentors (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    first_name VARCHAR,
    last_name VARCHAR,
    info VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS course_status
(
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR
);

CREATE TABLE IF NOT EXISTS courses_kk (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    status_id INT references course_status(id),
    mentor_id INT,
    category_id INT,
    name VARCHAR,
    title VARCHAR,
    description VARCHAR,
    effect VARCHAR,
    included_data VARCHAR,
    price INT DEFAULT 1000,
    views INT DEFAULT 0,
    purchases INT DEFAULT 0,
    font VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS courses_ru (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    status_id INT references course_status(id),
    mentor_id INT,
    category_id INT,
    name VARCHAR,
    title VARCHAR,
    description VARCHAR,
    effect VARCHAR,
    included_data VARCHAR,
    price INT DEFAULT 1000,
    views INT DEFAULT 0,
    purchases INT DEFAULT 0,
    font VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS courses_en (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    status_id INT references course_status(id),
    mentor_id INT,
    category_id INT,
    name VARCHAR,
    title VARCHAR,
    description VARCHAR,
    effect VARCHAR,
    included_data VARCHAR,
    price INT DEFAULT 1000,
    views INT DEFAULT 0,
    purchases INT DEFAULT 0,
    font VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS programs (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    course_id INT,
    video_id INT
);

CREATE TABLE IF NOT EXISTS chapters_kk (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    title VARCHAR,
    description VARCHAR
);

CREATE TABLE IF NOT EXISTS chapters_ru (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    title VARCHAR,
    description VARCHAR
);

CREATE TABLE IF NOT EXISTS chapters_en (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    title VARCHAR,
    description VARCHAR
);

CREATE TABLE IF NOT EXISTS videos (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    course_id INT,
    chapter_id INT,
    url VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS video_history (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    user_id INT,
    video_id INT,
    stop_time INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS feedbacks (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    user_id INT,
    course_id INT,
    comment VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS payment_statuses (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR
);

CREATE TABLE IF NOT EXISTS payment_methods (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR
);

CREATE TABLE IF NOT EXISTS recently_viewed_courses (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    user_id INT,
    course_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    user_id INT,
    course_id INT,
    status_id INT,
    payment_method_id INT,
    has_access BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS promocodes (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    code VARCHAR,
    sale_percent INT,
    issued_date TIMESTAMP,
    expired_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);