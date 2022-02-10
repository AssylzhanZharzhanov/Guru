CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    email VARCHAR,
    password VARCHAR,
    first_name VARCHAR,
    last_name VARCHAR,
    title VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories
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

CREATE TABLE IF NOT EXISTS courses (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    status_id INT REFERENCES course_status(id),
    mentor_id INT REFERENCES mentors(id),
    category_id INT REFERENCES categories(id),
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

CREATE TABLE IF NOT EXISTS videos (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    title VARCHAR,
    description VARCHAR,
    url VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS programs (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    course_id INT REFERENCES courses,
    video_id INT REFERENCES videos(id)
);

CREATE TABLE IF NOT EXISTS videos_history (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    user_id INT REFERENCES users(id),
    video_id INT REFERENCES videos(id),
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

INSERT INTO course_status(id, name) VALUES (1, 'active');
INSERT INTO course_status(id, name) VALUES (2, 'in progress');
INSERT INTO course_status(id, name) VALUES (3, 'closed');

INSERT INTO categories(id, name) VALUES (1, 'Art');
INSERT INTO categories(id, name) VALUES (2, 'Food');
INSERT INTO categories(id, name) VALUES (3, 'Sport');
INSERT INTO categories(id, name) VALUES (4, 'Space');
INSERT INTO categories(id, name) VALUES (5, 'Policy');

INSERT INTO mentors(first_name, last_name, info) VALUES ('Neil', 'Armstrong', 'First Moon Visitor');
INSERT INTO mentors(first_name, last_name, info) VALUES ('K-ZH', 'Tokayev', 'The President');

INSERT INTO courses(status_id, mentor_id, category_id, name, title, description, effect, included_data, font)
VALUES(1, 1, 4, 'Moon', 'To the Moon and back', 'America was on the moon',
       'USSR sucks', 'Make Moon Great Again', 'Times Moon Roman');
INSERT INTO courses(status_id, mentor_id, category_id, name, title, description, effect, included_data, font)
VALUES(1, 2, 5, 'New Kazakhstan', 'Shal Ket', 'New Kazakhstan is coming',
       'N1 OUT!', 'Fire!', 'Times KZ Roman');