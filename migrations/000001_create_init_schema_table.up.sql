CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email varchar,
    password varchar,
    first_name varchar,
    last_name varchar,
    title varchar,
    created_at timestamp
);

CREATE TABLE IF NOT EXISTS categories_kk
(
    id SERIAL PRIMARY KEY,
    name varchar
);

CREATE TABLE IF NOT EXISTS categories_ru
(
    id SERIAL PRIMARY KEY,
    name varchar
);

CREATE TABLE IF NOT EXISTS categories_en
(
    id SERIAL PRIMARY KEY,
    name varchar
);

CREATE TABLE IF NOT EXISTS mentors (
    id SERIAL PRIMARY KEY,
    first_name varchar,
    last_name varchar,
    info varchar,
    created_at timestamp
);

CREATE TABLE IF NOT EXISTS courses_kk (
    id SERIAL PRIMARY KEY,
    mentor_id int,
    category_id int,
    name varchar,
    title varchar,
    description varchar,
    effect varchar,
    included_data varchar,
    price int DEFAULT 1000,
    views int DEFAULT 0,
    purchases int DEFAULT 0,
    created_at timestamp
);

CREATE TABLE IF NOT EXISTS courses_ru (
    id SERIAL PRIMARY KEY,
    mentor_id int,
    category_id int,
    name varchar,
    title varchar,
    description varchar,
    effect varchar,
    included_data varchar,
    price int DEFAULT 1000,
    views int DEFAULT 0,
    purchases int DEFAULT 0,
    created_at timestamp
);

CREATE TABLE IF NOT EXISTS courses_en (
    id SERIAL PRIMARY KEY,
    mentor_id int,
    category_id int,
    name varchar,
    title varchar,
    description varchar,
    effect varchar,
    included_data varchar,
    price int DEFAULT 1000,
    views int DEFAULT 0,
    purchases int DEFAULT 0,
    created_at timestamp
);

CREATE TABLE IF NOT EXISTS programs (
    id SERIAL PRIMARY KEY,
    course_id int,
    video_id int
);

CREATE TABLE IF NOT EXISTS chapters_kk (
    id int,
    title varchar,
    description varchar
);

CREATE TABLE IF NOT EXISTS chapters_ru (
    id int,
    title varchar,
    description varchar
);

CREATE TABLE IF NOT EXISTS chapters_en (
    id int,
    title varchar,
    description varchar
);

CREATE TABLE IF NOT EXISTS videos (
    id SERIAL PRIMARY KEY,
    course_id int,
    chapter_id int,
    url varchar,
    created_at timestamp
);

CREATE TABLE IF NOT EXISTS video_history (
    id SERIAL PRIMARY KEY,
    user_id int,
    video_id int,
    stop_time int,
    created_at timestamp
);

CREATE TABLE IF NOT EXISTS feedbacks (
    id SERIAL PRIMARY KEY,
    user_id int,
    course_id int,
    comment varchar NOT NULL,
    created_at timestamp
);

CREATE TABLE IF NOT EXISTS payment_statuses (
    id SERIAL PRIMARY KEY,
    name varchar
);

CREATE TABLE IF NOT EXISTS payment_methods (
    id SERIAL PRIMARY KEY,
    name varchar
);

CREATE TABLE IF NOT EXISTS recently_viewed_courses (
    id SERIAL PRIMARY KEY,
    user_id int,
    course_id int,
    created_at timestamp
);

CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    user_id int,
    course_id int,
    status_id int,
    payment_method_id int,
    has_access boolean DEFAULT false,
    created_at timestamp
);

CREATE TABLE IF NOT EXISTS promocodes (
    id SERIAL PRIMARY KEY,
    code varchar,
    sale_percent int,
    issued_date timestamp,
    expired_date timestamp,
    created_at timestamp
);

ALTER TABLE courses ADD FOREIGN KEY (mentor_id) REFERENCES mentors(id);

ALTER TABLE courses ADD FOREIGN KEY (category_id) REFERENCES categories(id);

ALTER TABLE programs ADD FOREIGN KEY (course_id) REFERENCES courses (id);

ALTER TABLE programs ADD FOREIGN KEY (video_id) REFERENCES videos (id);

ALTER TABLE videos ADD FOREIGN KEY (course_id) REFERENCES courses (id);

ALTER TABLE videos ADD FOREIGN KEY (chapter_id) REFERENCES chapters (id);

ALTER TABLE video_history ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE video_history ADD FOREIGN KEY (video_id) REFERENCES videos (id);

ALTER TABLE feedbacks ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE feedbacks ADD FOREIGN KEY (course_id) REFERENCES courses (id);

ALTER TABLE recently_viewed_courses ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE recently_viewed_courses ADD FOREIGN KEY (course_id) REFERENCES courses (id);

ALTER TABLE payments ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE payments ADD FOREIGN KEY (course_id) REFERENCES courses (id);

ALTER TABLE payments ADD FOREIGN KEY (status_id) REFERENCES payment_statuses (id);

ALTER TABLE payments ADD FOREIGN KEY (payment_method_id) REFERENCES payment_methods (id);

ALTER TABLE promocodes ADD FOREIGN KEY (code) REFERENCES promocodes (id);

ALTER TABLE courses ADD FOREIGN KEY (included_data) REFERENCES courses (price);
